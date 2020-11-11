package main

import (
	//contracts "github.com/block42-blockchain-company/lachesis_exporter/contracts"
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

type ValidationStake struct {
	Status           *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	StakeAmount      *big.Int
	PaidUntilEpoch   *big.Int
	DelegatedMe      *big.Int
	DagAddress       common.Address
	SfcAddress       common.Address
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

func getStakers() map[int]ValidationStake {
	contract := sfc()
	stakers := make(map[int]ValidationStake)
	for i := 1; i < int(getStakersLastID()); i++ {
		staker, err := contract.Stakers(nil, big.NewInt(int64(i)))
		if err != nil {
			log.Fatal(err)
		}
		stakers[i] = staker
	}
	return stakers
}

func getStakersLastID() uint64 {
	stakerLastID, err := sfc().StakersLastID(nil)
	if err != nil {
		log.Fatal(err)
	}
	return stakerLastID.Uint64()
}

func getTotalSupply() uint64 {
	return getEpochSnapshot(big.NewInt(getCurrentEpoch())).TotalSupply.Uint64()
}

func getTotalStaked() uint64 {
	return getEpochSnapshot(big.NewInt(getCurrentEpoch())).StakeTotalAmount.Uint64()
}

func getTotalDelegated() uint64 {
	return getEpochSnapshot(big.NewInt(getCurrentEpoch())).DelegationsTotalAmount.Uint64()
}

func getTotalSelfStaked() uint64 {
	var totalSelfStaked uint64
	for _, staker := range getStakers() {
		totalSelfStaked += staker.StakeAmount.Uint64()
	}
	return totalSelfStaked
}
