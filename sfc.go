package main

import (
	_ "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"lachesis_exporter/contracts"
	"log"
	"math/big"
)

type Epoch struct {
	EndTime                *big.Int
	Duration               *big.Int
	EpochFee               *big.Int
	TotalBaseRewardWeight  *big.Int
	TotalTxRewardWeight    *big.Int
	BaseRewardPerSecond    *big.Int
	StakeTotalAmount       *big.Int
	DelegationsTotalAmount *big.Int
	TotalSupply            *big.Int
}

const (
	url              = "https://rpc.fantom.network"
	stakerAddressHex = "0xfc00face00000000000000000000000000000000"
)

func sfc() *contracts.Stakers {

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(stakerAddressHex)
	contract, err := contracts.NewStakers(address, client)
	if err != nil {
		log.Fatal(err)
	}

	num, err := contract.StakersNum(nil) // sanity check
	if err != nil {
		log.Fatal(err)
	} else if num.Cmp(big.NewInt(0)) != 0 {
		log.Fatal(num)
	}

	return contract
}

func getEpochSnapshot(epochID *big.Int) struct {
	EndTime                *big.Int
	Duration               *big.Int
	EpochFee               *big.Int
	TotalBaseRewardWeight  *big.Int
	TotalTxRewardWeight    *big.Int
	BaseRewardPerSecond    *big.Int
	StakeTotalAmount       *big.Int
	DelegationsTotalAmount *big.Int
	TotalSupply            *big.Int
} {
	epoch, err := sfc().EpochSnapshots(nil, epochID)
	if err != nil {
		log.Fatal(err)
	}
	return epoch
}
