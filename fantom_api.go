package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

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
