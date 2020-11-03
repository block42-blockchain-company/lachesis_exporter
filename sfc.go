package main

import (
	_ "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	contracts "lachesis_exporter/contracts"
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
	} else if num.Cmp(big.NewInt(0)) <= 0 {
		log.Fatal(num)
	}

	return contract
}

func getEpochSnapshot(epochID *big.Int) Epoch {
	epoch, err := sfc().EpochSnapshots(nil, epochID)
	if err != nil {
		log.Fatal(err)
	}
	return epoch
}

func getTotalSupply() int64 {
	return getEpochSnapshot(big.NewInt(getCurrentEpoch())).TotalSupply.Int64()
}

func getTotalStaked() int64 {
	return getEpochSnapshot(big.NewInt(getCurrentEpoch())).StakeTotalAmount.Int64()
}

func getTotalDelegated() int64 {
	return getEpochSnapshot(big.NewInt(getCurrentEpoch())).DelegationsTotalAmount.Int64()
}
