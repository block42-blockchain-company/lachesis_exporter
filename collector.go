package main

import (
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

var totalSupply = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "total_supply", Help: "FTM Total Supply"})

var totalStaked = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "total_staked", Help: "FTM Total Staked"})

var totalDelegated = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "total_delegated", Help: "FTM Total Delegated"})

/*
Node Related Metrics:
[x] # of connected peers
[] Up-/Downtime (sfc_getDowntime, stakerID)
[] Transactions-per-Second
[] Pending Transactions
[] Hardware and System Specs (# of CPU cores, OS, RAM)
*/

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

	go func() {
		for {
			totalSupply.Set(float64(getTotalSupply()))
			totalStaked.Set(float64(getTotalStaked()))
			totalDelegated.Set(float64(getTotalDelegated()))
			time.Sleep(2 * time.Second)
		}
	}()
}
