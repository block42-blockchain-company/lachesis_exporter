package main

import (
	"fmt"
	_ "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"lachesis_exporter/contracts"
	"log"
)

const (
	url              = "https://rpc.fantom.network"
	stakerAddressHex = "0xfc00face00000000000000000000000000000000"
)

func sfc() {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(stakerAddressHex)
	contract, err := contracts.NewStakers(address, client)

	if err != nil {
		log.Fatal(err)
	}

	stakersSum, err := contract.StakersNum(nil)
	fmt.Println(stakersSum)
}
