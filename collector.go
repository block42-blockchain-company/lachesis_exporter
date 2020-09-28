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

// URL of Lachesis API
const URL = "http://localhost:18545"

var stakerID string = "0xE"

// Declaring implemented metrics here

var currentEpoch = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "current_epoch", Help: "Current epoch number"})

var blockHeight = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "block_height", Help: "Total number of blocks"})

var peerCount = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "peer_count", Help: "Number of peers connected"})

var downTime = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "down_time", Help: "Seconds of node being down"})

var missedBlocks = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "missed_blocks", Help: "Amount of blocks missed"})

/*
Node Related Metrics:
[x] # of connected peers
[] Up-/Downtime (sfc_getDowntime, stakerID)
[] Transactions-per-Second
[] Pending Transactions
[] Hardware and System Specs (# of CPU cores, OS, RAM)
*/

func getBlockHeight() int64 {
	header := "application/json"
	body, _ := json.Marshal(&StringParamRequestBody{
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
		blockHeightVal, _ := strconv.ParseInt(data.Result, 0, 64)
		return blockHeightVal
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
		currentEpochVal, _ := strconv.ParseInt(data.Result, 0, 64)
		return currentEpochVal
	}
}

func getPeerCount() int64 {
	header := "application/json"
	body, _ := json.Marshal(&StringParamRequestBody{
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

func getDownTime() (int64, int64) {
	header := "application/json"
	var param []string
	param = append(param, stakerID)
	body, _ := json.Marshal(&StringParamRequestBody{
		JSONRPC: "2.0",
		Method:  "sfc_getDowntime",
		ID:      1,
		Params:  param,
	})

	response, err := http.Post(URL, header, bytes.NewBuffer(body))
	PrintResponse(*response)
	if err != nil {
		fmt.Println(err.Error())
		return 0, 0
	} else {
		var data ResponseBody
		err := json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			panic(err)
		}

		var Result DownTimeResponse

		json.Unmarshal([]byte(data.Result), &Result)
		downtimeVal, _ := strconv.ParseInt(Result.Downtime, 0, 64)
		missedBlocksVal, _ := strconv.ParseInt(Result.MissedBlocks, 0, 64)

		fmt.Printf("%+v", Result)

		return downtimeVal, missedBlocksVal
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

	go func() {
		for {
			downTimeVal, missedBlocksVal := getDownTime()
			downTime.Set(float64(downTimeVal))
			missedBlocks.Set(float64(missedBlocksVal))
			time.Sleep(2 * time.Second)
		}
	}()

}
