package prometheus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	baseURL string
}

func NewClient(baseURL string) (*Client, error) {
	return &Client{baseURL: baseURL}, nil
}

func (client *Client) Query(method string, request *QueryRequest) (*QueryResponse, error) {
	var urlString = fmt.Sprintf("%s/api/v1/query", client.baseURL)

	date, err := time.Parse("2006-01-02 15:04:05", request.Time)
	if err != nil {
		return nil, err
	}
	date = date.Add(-8 * time.Hour)
	request.Time = date.Format(time.RFC3339)
	bytes, err := makePrometheusRequest(urlString, strings.ToUpper(method), nil, request)
	if err != nil {
		return nil, err
	}

	var response QueryResponse
	return &response, json.Unmarshal(bytes, &response)
}

func makePrometheusRequest(url, method string, header map[string]string, data interface{}) ([]byte, error) {
	var (
		request *http.Request
		err     error
	)

	values, err := structToMap(data)
	if err != nil {
		return nil, err
	}

	switch strings.ToUpper(method) {
	case "GET":
		request, err = http.NewRequest(strings.ToUpper(method), url, nil)
		if err != nil {
			return nil, err
		}
		request.URL.RawQuery = values.Encode()
	case "POST":
		request, err = http.NewRequest(strings.ToUpper(method), url, strings.NewReader(values.Encode()))
		if err != nil {
			return nil, err
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		request.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))
	default:
		return nil, fmt.Errorf("prometheus request method: %s is not support", method)
	}

	for key, value := range header {
		request.Header.Add(key, value)
	}

	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func structToMap(inter interface{}) (url.Values, error) {
	values := url.Values{}

	val := reflect.ValueOf(inter).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag.Get("json")

		switch typeField.Type.Kind() {
		case reflect.String:
			value := fmt.Sprintf("%s", valueField.Interface())
			values.Add(tag, value)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value := fmt.Sprintf("%d", valueField.Interface())
			values.Add(tag, value)
		case reflect.Float32, reflect.Float64:
			value := fmt.Sprintf("%d", valueField.Interface())
			values.Add(tag, value)
		case reflect.Bool:
			value := fmt.Sprintf("%t", valueField.Interface())
			values.Add(tag, value)
		case reflect.Slice:
			for index := 0; index < valueField.Len(); index++ {
				item := valueField.Index(index)
				switch item.Kind() {
				case reflect.String:
					value := fmt.Sprintf("%s", item.Interface())
					values.Add(tag, value)
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					value := fmt.Sprintf("%d", valueField.Interface())
					values.Add(tag, value)
				case reflect.Float32, reflect.Float64:
					value := fmt.Sprintf("%d", valueField.Interface())
					values.Add(tag, value)
				case reflect.Bool:
					value := fmt.Sprintf("%t", valueField.Interface())
					values.Add(tag, value)
				default:
					return nil, fmt.Errorf("slice type %s not support", typeField.Name)
				}
			}
		default:
			return nil, fmt.Errorf("type %s not support", typeField.Name)
		}
	}
	return values, nil
}
