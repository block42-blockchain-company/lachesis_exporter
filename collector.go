package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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

var txPerSecond = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "tx_per_second", Help: "Transactions per second"})

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
	if err != nil {
		fmt.Println(err.Error())
		return 0, 0
	} else {
		var data DownTimeResponse
		err := json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			panic(err)
		}

		downtimeVal, _ := strconv.ParseInt(data.Result.Downtime, 0, 64)
		missedBlocksVal, _ := strconv.ParseInt(data.Result.MissedBlocks, 0, 64)

		return downtimeVal, missedBlocksVal
	}
}

func getEventIDs(epochNumber string) []string {
	header := "application/json"
	body, _ := json.Marshal(&StringParamRequestBody{
		JSONRPC: "2.0",
		Method:  "ftm_getHeads",
		ID:      1,
		Params:  strings.Fields(epochNumber),
	})

	response, err := http.Post(URL, header, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err.Error())
		return strings.Fields("") // lol xD
	} else {
		var data ResponseBodyHead
		err := json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			panic(err)
		}
		return data.Result
	}
}

func getEvent(id string) map[string]interface{} {
	header := "application/json"
	body, _ := json.Marshal(&InterfaceParamRequestBody{
		JSONRPC: "2.0",
		Method:  "ftm_getEvent",
		ID:      1,
		Params:  []interface{}{id, true},
	})
	response, err := http.Post(URL, header, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	} else {
		var data ResponseBodyHeadInterface
		err := json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			panic(err)
		}
		return data.Result
	}
}

func getTxPerSecond() int64 {
	// try to open a transactions.json, if not existing create a new one
	transactions, err := openJson("transactions.json")
	if err != nil {
		transactions = &Transactions{0, time.Now().Unix()}
	}

	IDs := getEventIDs("latest")
	for _, id := range IDs {
		event := getEvent(id)
		txArr := event["transactions"].([]interface{})
		if len(txArr) > 0 {
			transactions.Count += len(txArr)
		}
	}
	saveJson("transactions.json", transactions)

	timeDiff := transactions.Start - time.Now().Unix()
	if timeDiff == 0 {
		timeDiff = 1
	}
	return int64(transactions.Count) / timeDiff
}

func openJson(fileName string) (*Transactions, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var transactions Transactions
	json.Unmarshal(byteValue, &transactions)
	return &transactions, nil
}

func saveJson(fileName string, jsonData interface{}) {
	jsonString, _ := json.MarshalIndent(jsonData, "", " ")
	_ = ioutil.WriteFile(fileName, jsonString, 0644)
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

	go func() {
		for {
			txPerSecond.Set(float64(getTxPerSecond()))
			time.Sleep(2 * time.Second)
		}
	}()
}
