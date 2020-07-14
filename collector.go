package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

// TODO: Heavy Refactor!
var currentEpoch = prometheus.NewSummary(prometheus.SummaryOpts{
	Name: "current_epoch", Help: "Current epoch number"})

const url = "http://localhost:18545"

// GetEpochRequestBody For evaluating the current epoch
func GetEpochRequestBody() ([]byte, error) {
	return json.Marshal(&struct {
		JSONRPC string `json:"jsonrpc"`
		Method  string `json:"method"`
		ID      int64  `json:"id"`
	}{
		JSONRPC: "2.0",
		Method:  "ftm_currentEpoch",
		ID:      1,
	})
}

func getBlockHeight() int16 {
	header := "application/json"
	body, _ := GetEpochRequestBody()
	response, err := http.Post(url, header, bytes.NewBuffer(body))

	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	return 1
}

func main() {
	getBlockHeight()
}

func init() {
	prometheus.MustRegister(currentEpoch)
}
