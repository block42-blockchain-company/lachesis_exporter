package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
)

func PrintResponse(resp http.Response) {
	log.Info("Reading Body")
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Info(bodyString)
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
