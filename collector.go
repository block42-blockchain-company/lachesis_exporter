package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const URL = "http://localhost:18545"

// Declaring implemented metrics here

var currentEpoch = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "current_epoch", Help: "Current epoch number"})

var blockHeight = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "block_height", Help: "Total number of blocks"})

var peerCount = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "peer_count", Help: "Number of peers connected"})

/*
Node Related Metrics:
[x] # of connected peers
[] Up-/Downtime (sfc_getDowntime, stakerID)
[] Transactions-per-Second
[] Pending Transactions
[] Hardware and System Specs (# of CPU cores, OS, RAM)
*/

type ResponseBody struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int64  `json:"id"`
	Result  string `json:"result"`
}

type RequestBody struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	ID      int64  `json:"id"`
}

type ParamRequestBody struct {
	JSONRPC string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	ID      int64    `json:"id"`
	Params  []string `json:"params"`
}

func getBlockHeight() int64 {
	header := "application/json"
	body, _ := json.Marshal(&ParamRequestBody{
		JSONRPC: "2.0",
		Method:  "eth_blockNumber",
		ID:      83,
		Params:  nil,
	})

	response, err := http.Post(URL, header, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err.Error())
		return 0
	} else {
		var data ResponseBody
		err := json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			panic(err)
		}
		blocknr, _ := strconv.ParseInt(data.Result, 0, 64)
		return blocknr
	}
}

func getCurrentEpoch() int64 {
	header := "application/json"
	body, _ := json.Marshal(&RequestBody{
		JSONRPC: "2.0",
		Method:  "ftm_currentEpoch",
		ID:      1,
	})

	response, err := http.Post(URL, header, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err.Error())
		return 0
	} else {
		var data ResponseBody
		err := json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			panic(err)
		}
		epoch, _ := strconv.ParseInt(data.Result, 0, 64)
		return epoch
	}
}

func getPeerCount() int64 {
	header := "application/json"
	body, _ := json.Marshal(&ParamRequestBody{
		JSONRPC: "2.0",
		Method:  "net_peerCount",
		ID:      1,
		Params:  nil,
	})

	response, err := http.Post(URL, header, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err.Error())
		return 0
	} else {
		var data ResponseBody
		err := json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			panic(err)
		}
		peerCountVal, _ := strconv.ParseInt(data.Result, 0, 64)
		return peerCountVal
	}
}

// RecordMetrics | Update all metrics
func RecordMetrics() {
	go func() {
		for {
			currentEpoch.Set(float64(getCurrentEpoch()))
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			blockHeight.Set(float64(getBlockHeight()))
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			peerCount.Set(float64(getPeerCount()))
			time.Sleep(2 * time.Second)
		}
	}()

}
