package main

import (
	"io/ioutil"
	"net/http"

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
