package prometheus

import (
	"encoding/json"
	"fmt"
)

type BaseResponse struct {
	Status    string   `json:"status"`
	ErrorType string   `json:"errorType"`
	Error     string   `json:"error"`
	Warnings  []string `json:"warnings"`
}

type QueryRequest struct {
	Query   string `json:"query"`
	Time    string `json:"time"`
	Timeout string `json:"timeout"`
}

type QueryResponse struct {
	BaseResponse
	Data QueryData `json:"data"`
}

type QueryData struct {
	ResultType string           `json:"resultType"`
	Result     *json.RawMessage `json:"result"`
}

func (data *QueryData) GetMareixResult() ([]MatrixResult, error) {
	if data.ResultType != "matrix" {
		return nil, fmt.Errorf("query response data type: %s is not matrix", data.ResultType)
	}

	var result []MatrixResult
	return result, json.Unmarshal(*data.Result, &result)
}

func (data *QueryData) GetVectorResult() ([]VectorResult, error) {
	if data.ResultType != "vector" {
		return nil, fmt.Errorf("query response data type: %s is not vector", data.ResultType)
	}

	var result []VectorResult
	return result, json.Unmarshal(*data.Result, &result)
}

func (data *QueryData) GetScalarResult() (*ScalarsResult, error) {
	if data.ResultType != "scalar" {
		return nil, fmt.Errorf("query response data type: %s is not scalar", data.ResultType)
	}

	var result ScalarsResult
	return &result, json.Unmarshal(*data.Result, &result)
}

func (data *QueryData) GetStringResult() (*StringsResult, error) {
	if data.ResultType != "string" {
		return nil, fmt.Errorf("query response data type: %s is not string", data.ResultType)
	}

	var result StringsResult
	return &result, json.Unmarshal(*data.Result, &result)
}

type MatrixResult struct {
	Metric map[string]string `json:"metric"`
	Values [][]interface{}   `json:"value"`
}

type VectorResult struct {
	Metric map[string]string `json:"metric"`
	Values []interface{}     `json:"value"`
}

type ScalarsResult []interface{}

type StringsResult []interface{}
