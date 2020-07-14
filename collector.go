package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

// TODO: Heavy Refactor!
var currentEpoch = prometheus.NewSummary(prometheus.SummaryOpts{
	Name: "current_epoch", Help: "Current epoch number"})

const url = "http://localhost:18545"

type EpochResponseBody struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int64  `json:"id"`
	Result  string `json:"result"`
}

type EpochRequestBody struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	ID      int64  `json:"id"`
}

func getBlockHeight() int64 {
	header := "application/json"
	body, _ := json.Marshal(&EpochRequestBody{
		JSONRPC: "2.0",
		Method:  "ftm_currentEpoch",
		ID:      1,
	})

	response, err := http.Post(url, header, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	} else {
		var data EpochResponseBody
		err := json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			panic(err)
		}
		fmt.Println(data.Result)
		epoch, _ := strconv.ParseInt(data.Result, 0, 64)
		return epoch
	}
}

func main() {
	result := fmt.Sprint(getBlockHeight())
	fmt.Println(result)
}

func init() {
	prometheus.MustRegister(currentEpoch)
}
