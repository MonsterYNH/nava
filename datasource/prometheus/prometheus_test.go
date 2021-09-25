package prometheus

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	client, _ := NewClient("http://localhost:9090")

	response, err := client.Query("get", &QueryRequest{
		Query: `node_cpu_seconds_total{instance="172.11.0.5:9291"}`,
		Time:  "2021-09-12 20:00:00",
		// Timeout: "5",
	})

	if err != nil {
		t.Fatal(err)
	}

	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))

	datas, err := response.Data.GetVectorResult()
	if err != nil {
		t.Fatal(err)
	}
	bytes, _ = json.Marshal(datas)
	fmt.Println(string(bytes))

	// response.Data.GetMareixResult()
}
