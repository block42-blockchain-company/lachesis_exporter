// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StakersABI is the input ABI used to generate the binding from.
const StakersABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isDelegation\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BurntRewardStash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"untilEpoch\",\"type\":\"uint256\"}],\"name\":\"ClaimedDelegationReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"untilEpoch\",\"type\":\"uint256\"}],\"name\":\"ClaimedValidatorReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CreatedDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dagSfcAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CreatedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"delegation\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CreatedWithdrawRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"DeactivatedDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"DeactivatedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"IncreasedDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"IncreasedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"LockingDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"LockingStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"delegation\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"PartialWithdrawnByRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"PreparedToWithdrawDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"PreparedToWithdrawStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"UnstashedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"UpdatedBaseRewardPerSec\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldStakerID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newStakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UpdatedDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"short\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"long\",\"type\":\"uint256\"}],\"name\":\"UpdatedGasPowerAllocationRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegatedMe\",\"type\":\"uint256\"}],\"name\":\"UpdatedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"UpdatedStakerMetadata\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldSfcAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSfcAddress\",\"type\":\"address\"}],\"name\":\"UpdatedStakerSfcAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"WithdrawnDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"WithdrawnStake\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isDelegation\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"_rewardsBurnableOnDeactivation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sfcAddress\",\"type\":\"address\"}],\"name\":\"_sfcAddressToStakerID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"_syncDelegator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"_syncStaker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"_upgradeStakerStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bondedRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bondedTargetPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bondedTargetRewardUnlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bondedTargetStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"}],\"name\":\"calcDelegationRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"}],\"name\":\"calcValidatorRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"}],\"name\":\"claimDelegationRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"}],\"name\":\"claimValidatorRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"}],\"name\":\"createDelegation\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"createStake\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dagAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sfcAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"createStakeWithAddresses\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSealedEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delegationEarlyWithdrawalPenalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationLockPeriodEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationLockPeriodTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paidUntilEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationsNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationsTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochSnapshots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBaseRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTxRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardPerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeTotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegationsTotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"e\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"epochValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedMe\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txRewardWeight\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"firstLockedUpEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getStakerID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseDelegation\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseStake\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lockDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"lockUpDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lockDuration\",\"type\":\"uint256\"}],\"name\":\"lockUpStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lockedDelegations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lockedStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxDelegatedRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxStakerMetadataSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDelegationDecrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDelegationIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStakeDecrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStakeIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"}],\"name\":\"partialWithdrawByRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareToWithdrawDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"prepareToWithdrawDelegationPartial\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareToWithdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"prepareToWithdrawStakePartial\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardsAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardsStash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slashedDelegationsTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slashedStakeTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeLockPeriodEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeLockPeriodTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakerMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paidUntilEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedMe\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dagAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sfcAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakersLastID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakersNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochNum\",\"type\":\"uint256\"}],\"name\":\"startLockedUp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBurntLockupRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unbondingStartDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unbondingUnlockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unlockedRewardRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unstashRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"updateBaseRewardPerSec\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"short\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"long\",\"type\":\"uint256\"}],\"name\":\"updateGasPowerAllocationRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"updateStakerMetadata\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSfcAddress\",\"type\":\"address\"}],\"name\":\"updateStakerSfcAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawalRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegation\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StakersBin is the compiled bytecode used for deploying new contracts.
var StakersBin = "0x60806040819052600080546001600160a01b031916339081178255918291907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506156ec806100536000396000f3fe6080604052600436106105225760003560e01c80637b8c6b02116102af578063bed9d86111610179578063d9e257ef116100d6578063ec6a7f1c1161008a578063f3ae5b1a1161006f578063f3ae5b1a14611125578063f8b18d8a1461114f578063fd5e6dd11461117957610522565b8063ec6a7f1c14610856578063f2fde38b146110f257610522565b8063df4f49d4116100bb578063df4f49d41461106e578063e7ff9e7814611098578063eac3baf2146110c857610522565b8063d9e257ef14611014578063dd099bb61461101c57610522565b8063c4b5dd7e1161012d578063cb1c4e6711610112578063cb1c4e6714610527578063cc8c212014610f59578063ce5aa00014610fff57610522565b8063c4b5dd7e14610527578063c9400d4f14610f2f57610522565b8063c312eb071161015e578063c312eb0714610eca578063c3d74f1a14610ee7578063c41b640514610f1a57610522565b8063bed9d86114610e4a578063bffe348614610e5f57610522565b806398ec2de511610227578063ab2273c0116101db578063b42cb58d116101c0578063b42cb58d14610d82578063b6732f3d14610db5578063b9029d5014610df457610522565b8063ab2273c014610d3b578063b1a3ebfa14610d5057610522565b8063a4b89fab1161020c578063a4b89fab14610ce1578063a70da4d214610d11578063a778651514610d2657610522565b806398ec2de514610c2d578063a289ad6e14610ccc57610522565b8063876f7e2a1161027e5780638f32d59b116102635780638f32d59b14610b0557806390475ae414610b1a57806396060e7114610bd957610522565b8063876f7e2a14610abf5780638da5cb5b14610ad457610522565b80637b8c6b0214610a575780637cacb1d614610a6c57806381d9dc7a14610a815780638447c4df14610a9657610522565b80633a0af4d4116103f05780635e2308d2116103685780636e1a767a1161031c578063715018a611610301578063715018a614610a035780637667180814610a18578063793c45ce14610a2d57610522565b80636e1a767a146109b55780636f498663146109ca57610522565b806363321e271161034d57806363321e271461093457806366799a54146109675780636a1cf400146109a057610522565b80635e2308d2146106de57806360c7e37f1461052757610522565b80634bd202dc116103bf57806353a72586116103a457806353a72586146108e657806354d77ed2146106135780635dc03f1f146108fb57610522565b80634bd202dc1461086b5780634e5a23281461088057610522565b80633a0af4d4146108245780633a274ff6146108395780633d0317fe146108415780633fee10a81461085657610522565b80631d58179c1161049e57806328dca8ff1161045257806330fa99291161043757806330fa99291461074757806333a149121461075c578063375b3c0a1461080f57610522565b806328dca8ff146106f3578063295cccba1461071d57610522565b80632265f284116104835780632265f2841461069957806326682c71146106ae5780632709275e146106de57610522565b80631d58179c146106135780631e8a69561461062857610522565b8063119e351a116104f557806319ddb54f116104da57806319ddb54f146105275780631b593d8a146105d45780631c333318146105fe57610522565b8063119e351a1461058d57806316bfdd81146105bf57610522565b80630298599214610527578063041d97a31461054e57806308728f6e146105635780630a29180c14610578575b600080fd5b34801561053357600080fd5b5061053c6111ff565b60408051918252519081900360200190f35b34801561055a57600080fd5b5061053c61120c565b34801561056f57600080fd5b5061053c611287565b34801561058457600080fd5b5061053c61128d565b34801561059957600080fd5b506105bd600480360360408110156105b057600080fd5b5080359060200135611293565b005b3480156105cb57600080fd5b506105bd61132b565b3480156105e057600080fd5b506105bd600480360360208110156105f757600080fd5b5035611623565b34801561060a57600080fd5b506105bd6116b2565b34801561061f57600080fd5b5061053c61180c565b34801561063457600080fd5b506106526004803603602081101561064b57600080fd5b5035611811565b60408051998a5260208a0198909852888801969096526060880194909452608087019290925260a086015260c085015260e084015261010083015251908190036101200190f35b3480156106a557600080fd5b5061053c61185f565b3480156106ba57600080fd5b506105bd600480360360408110156106d157600080fd5b5080359060200135611866565b3480156106ea57600080fd5b5061053c611b71565b3480156106ff57600080fd5b506105bd6004803603602081101561071657600080fd5b5035611b81565b34801561072957600080fd5b506105bd6004803603602081101561074057600080fd5b5035611c39565b34801561075357600080fd5b5061053c611d1a565b34801561076857600080fd5b506105bd6004803603602081101561077f57600080fd5b81019060208101813564010000000081111561079a57600080fd5b8201836020820111156107ac57600080fd5b803590602001918460018302840111640100000000831117156107ce57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611d20945050505050565b34801561081b57600080fd5b5061053c611de2565b34801561083057600080fd5b5061053c611df1565b6105bd611df8565b34801561084d57600080fd5b5061053c612004565b34801561086257600080fd5b5061053c61200a565b34801561087757600080fd5b5061053c612011565b34801561088c57600080fd5b506108b9600480360360408110156108a357600080fd5b506001600160a01b038135169060200135612017565b60408051958652602086019490945284840192909252606084015215156080830152519081900360a00190f35b3480156108f257600080fd5b5061053c612054565b34801561090757600080fd5b506105bd6004803603604081101561091e57600080fd5b506001600160a01b03813516906020013561205c565b34801561094057600080fd5b5061053c6004803603602081101561095757600080fd5b50356001600160a01b03166120c0565b34801561097357600080fd5b5061053c6004803603604081101561098a57600080fd5b506001600160a01b0381351690602001356120df565b3480156109ac57600080fd5b5061053c6120fc565b3480156109c157600080fd5b5061053c612165565b3480156109d657600080fd5b5061053c600480360360408110156109ed57600080fd5b506001600160a01b03813516906020013561216b565b348015610a0f57600080fd5b506105bd612188565b348015610a2457600080fd5b5061053c612238565b348015610a3957600080fd5b506105bd60048036036020811015610a5057600080fd5b5035612241565b348015610a6357600080fd5b5061053c612350565b348015610a7857600080fd5b5061053c612358565b348015610a8d57600080fd5b5061053c61235e565b348015610aa257600080fd5b50610aab612364565b604080519115158252519081900360200190f35b348015610acb57600080fd5b506105bd612398565b348015610ae057600080fd5b50610ae9612505565b604080516001600160a01b039092168252519081900360200190f35b348015610b1157600080fd5b50610aab612514565b6105bd60048036036060811015610b3057600080fd5b6001600160a01b038235811692602081013590911691810190606081016040820135640100000000811115610b6457600080fd5b820183602082011115610b7657600080fd5b80359060200191846001830284011164010000000083111715610b9857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612525945050505050565b348015610be557600080fd5b50610c0f60048036036060811015610bfc57600080fd5b50803590602081013590604001356125a2565b60408051938452602084019290925282820152519081900360600190f35b348015610c3957600080fd5b50610c5760048036036020811015610c5057600080fd5b50356125e2565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610c91578181015183820152602001610c79565b50505050905090810190601f168015610cbe5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610cd857600080fd5b5061053c61267d565b348015610ced57600080fd5b506105bd60048036036040811015610d0457600080fd5b5080359060200135612683565b348015610d1d57600080fd5b5061053c61297f565b348015610d3257600080fd5b5061053c612985565b348015610d4757600080fd5b5061053c612992565b348015610d5c57600080fd5b50610aab60048036036040811015610d7357600080fd5b50803515159060200135612998565b348015610d8e57600080fd5b5061053c60048036036020811015610da557600080fd5b50356001600160a01b03166129ef565b348015610dc157600080fd5b50610c0f60048036036060811015610dd857600080fd5b506001600160a01b038135169060208101359060400135612a44565b348015610e0057600080fd5b50610e2460048036036040811015610e1757600080fd5b5080359060200135612a5f565b604080519485526020850193909352838301919091526060830152519081900360800190f35b348015610e5657600080fd5b506105bd612a92565b348015610e6b57600080fd5b50610e9260048036036020811015610e8257600080fd5b50356001600160a01b0316612d7d565b604080519788526020880196909652868601949094526060860192909252608085015260a084015260c0830152519081900360e00190f35b6105bd60048036036020811015610ee057600080fd5b5035612dba565b348015610ef357600080fd5b506105bd60048036036020811015610f0a57600080fd5b50356001600160a01b0316612dc7565b348015610f2657600080fd5b506105bd613042565b348015610f3b57600080fd5b506105bd60048036036020811015610f5257600080fd5b5035613130565b6105bd60048036036020811015610f6f57600080fd5b810190602081018135640100000000811115610f8a57600080fd5b820183602082011115610f9c57600080fd5b80359060200191846001830284011164010000000083111715610fbe57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550613247945050505050565b34801561100b57600080fd5b5061053c613253565b6105bd613261565b34801561102857600080fd5b506110556004803603604081101561103f57600080fd5b506001600160a01b03813516906020013561335d565b6040805192835260208301919091528051918290030190f35b34801561107a57600080fd5b506110556004803603602081101561109157600080fd5b5035613381565b3480156110a457600080fd5b506105bd600480360360408110156110bb57600080fd5b508035906020013561339a565b3480156110d457600080fd5b506105bd600480360360208110156110eb57600080fd5b5035613671565b3480156110fe57600080fd5b506105bd6004803603602081101561111557600080fd5b50356001600160a01b03166136d0565b34801561113157600080fd5b506105bd6004803603602081101561114857600080fd5b5035613732565b34801561115b57600080fd5b506105bd6004803603602081101561117257600080fd5b503561391e565b34801561118557600080fd5b506111a36004803603602081101561119c57600080fd5b5035613d75565b604080519a8b5260208b0199909952898901979097526060890195909552608088019390935260a087019190915260c086015260e08501526001600160a01b039081166101008501521661012083015251908190036101400190f35b670de0b6b3a76400005b90565b601e546000908152601f602052604081206009015480611230576000915050611209565b601e546000908152601f60205260408120600881015460079091015461125b9163ffffffff613dd216565b90506112808261127483620f424063ffffffff613e2c16565b9063ffffffff613e8516565b9250505090565b60235481565b60285481565b61129b612514565b6112ec576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b604080518381526020810183905281517f95ae5488127de4bc98492f4487556e7af9f37eb4b6d5e94f6d849e03ff76cc7c929181900390910190a15050565b336113346154f0565b506001600160a01b038116600090815260296020908152604091829020825160e0810184528154815260018201549281019290925260028101549282019290925260038201546060820181905260048301546080830152600583015460a083015260069092015460c0820181905290916113f5576040805162461bcd60e51b815260206004820152601d60248201527f64656c65676174696f6e207761736e2774206465616374697661746564000000604482015290519081900360640190fd5b6000818152602080526040902060050154156114d65761141361200a565b82606001510142101561146d576040805162461bcd60e51b815260206004820152601660248201527f6e6f7420656e6f7567682074696d652070617373656400000000000000000000604482015290519081900360640190fd5b61147561180c565b826040015101611483612238565b10156114d6576040805162461bcd60e51b815260206004820152601860248201527f6e6f7420656e6f7567682065706f636873207061737365640000000000000000604482015290519081900360640190fd5b6000818152602080805260408083205460808601516001600160a01b03881680865260298552838620868155600181810188905560028201889055600382018890556004820188905560058201889055600690910187905560318652848720888852865284872087815581018790559086526032855283862087875290945291842084905560258054600019019055602654921615159161157d908263ffffffff613ec716565b602655816115c1576040516001600160a01b0387169082156108fc029083906000818181858888f193505050501580156115bb573d6000803e3d6000fd5b506115c5565b8092505b6027546115d8908463ffffffff613dd216565b60275560408051848152905185916001600160a01b038916917f87e86b3710b72c10173ca52c6a9f9cf2df27e77ed177741a8b4feb12bb7a606f9181900360200190a3505050505050565b61162b612514565b61167c576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6040805182815290517f8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae23969181900360200190a150565b33600081815260296020526040902060068101546116d08382613f09565b6116da8382613f85565b6116f06001828585600401548660040154613ff6565b6116f8612238565b6002830155426003830155600482015460008281526020805260409020600501541561175157600082815260208052604090206007015461173f908263ffffffff613ec716565b60008381526020805260409020600701555b600061175d8584614106565b156117a85761176d858484614165565b905081811061177d575060001981015b6001600160a01b03851660009081526032602090815260408083208684529091529020805482900390555b600484018054829003905560268054829003905560405183906001600160a01b038716907f912c4125a208704a342cbdc4726795d26556b0170b7fc95bc706d5cb1f50646990600090a36117fc858461205c565b61180583613671565b5050505050565b600390565b601f6020528060005260406000206000915090508060010154908060020154908060030154908060040154908060050154908060060154908060070154908060080154908060090154905089565b62e4e1c090565b336000611872826129ef565b905061187d816141b5565b61188681614221565b61188f81614287565b156118e1576040805162461bcd60e51b815260206004820152600f60248201527f7374616b65206973206c6f636b65640000000000000000000000000000000000604482015290519081900360640190fd5b6118e96111ff565b83101561193d576040805162461bcd60e51b815260206004820152601060248201527f746f6f20736d616c6c20616d6f756e7400000000000000000000000000000000604482015290519081900360640190fd5b600081815260208052604090206005015480611957611de2565b850111156119ac576040805162461bcd60e51b815260206004820152601c60248201527f6d757374206c65617665206174206c65617374206d696e5374616b6500000000604482015290519081900360640190fd5b6000828152602080526040902060070154848203906119ca826142bb565b1015611a1d576040805162461bcd60e51b815260206004820152601460248201527f746f6f206d7563682064656c65676174696f6e73000000000000000000000000604482015290519081900360640190fd5b6001600160a01b0384166000908152602d6020908152604080832089845290915290206003015415611a96576040805162461bcd60e51b815260206004820152601360248201527f7772494420616c72656164792065786973747300000000000000000000000000604482015290519081900360640190fd5b611aa4600084868886613ff6565b600083815260208080526040808320600501805489900390556001600160a01b0387168352602d82528083208984529091529020838155600301859055611ae9612238565b6001600160a01b0385166000818152602d602090815260408083208b8452825280832060018101959095554260029095019490945583518a8152908101919091528083018890529151859282917fde2d2a87af2fa2de55bde86f04143144eb632fa6be266dc224341a371fb8916d9181900360600190a4611b6983613671565b505050505050565b600060646301c9c3805b04905090565b60008181526020805260409020600901546001600160a01b031615611bed576040805162461bcd60e51b815260206004820152600f60248201527f616c726561647920757064617465640000000000000000000000000000000000604482015290519081900360640190fd5b611bf6816142dc565b6000908152602080526040902060088101546009909101805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909216919091179055565b336000611c45826129ef565b9050611c50816142dc565b611c5861552d565b600080611c678460008861433e565b925092509250600083604001518460200151856000015101019050611ca36020600087815260200190815260200160002060060154848461447d565b600085815260208052604090206006018290556060840151603380549091019055611cce868261457d565b6040805182815260208101859052808201849052905186917f2ea54c2b22a07549d19fb5eb8e4e48ebe1c653117215e94d5468c5612750d35c919081900360600190a250505050505050565b60265481565b6000611d2b336129ef565b9050611d36816142dc565b611d3e612992565b82511115611d93576040805162461bcd60e51b815260206004820152601060248201527f746f6f20626967206d6574616461746100000000000000000000000000000000604482015290519081900360640190fd5b6000818152602b602090815260409091208351611db292850190615555565b5060405181907fb7a99a0df6a9e15c2689e6a55811ef76cdb514c67d4a0e37fcb125ada0e3cd8390600090a25050565b6a02a055184a310c1260000090565b62ed4e0090565b33600081815260296020526040902060060154611e158282613f09565b611e1f8282613f85565b611e276111ff565b341015611e7b576040805162461bcd60e51b815260206004820152601360248201527f696e73756666696369656e7420616d6f756e7400000000000000000000000000604482015290519081900360640190fd5b6000818152602080526040902060070154611e9c903463ffffffff613dd216565b6000828152602080526040902060050154611eb6906142bb565b1015611f09576040805162461bcd60e51b815260206004820152601a60248201527f7374616b65722773206c696d6974206973206578636565646564000000000000604482015290519081900360640190fd5b611f128161462c565b6001600160a01b038216600090815260296020526040812060040154611f3e903463ffffffff613dd216565b6001600160a01b0384166000908152602960209081526040808320600401849055858352908052902060070154909150611f7e903463ffffffff613dd216565b6000838152602080526040902060070155602654611fa2903463ffffffff613dd216565b60265560408051828152346020820152815184926001600160a01b038716927f4ca781bfe171e588a2661d5a7f2f5f59df879c53489063552fbad2145b707fc1929081900390910190a3611ff6838361205c565b611fff82613671565b505050565b60245481565b62093a8090565b60255481565b602d602090815260009283526040808420909152908252902080546001820154600283015460038401546004909401549293919290919060ff1685565b635e0580f890565b6120668282614695565b6001600160a01b03821660008181526029602090815260409182902060040154825190815291518493849390927f19b46b9014e4dc8ca74f505b8921797c6a8a489860217d15b3c7d741637dfcff92918290030190a45050565b6001600160a01b0381166000908152602160205260409020545b919050565b603260209081526000928352604080842090915290825290205481565b60008061211761210a612054565b429063ffffffff613ec716565b90506000612139612126612350565b611274620f42408563ffffffff613e2c16565b9050612143613253565b811061215457600092505050611209565b8061215d613253565b039250505090565b602f5481565b602c60209081526000928352604080842090915290825290205481565b612190612514565b6121e1576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36000805473ffffffffffffffffffffffffffffffffffffffff19169055565b601e5460010190565b336000818152602960205260409020600681015461225f8382613f09565b61226761552d565b60008061227686600089614702565b9250925092506000836040015184602001518560000151010190506122a08660050154848461447d565b600586018290556040848101516020808701516001600160a01b038b166000908152603283528481208a825290925292902080546002909304909101909101905560608401516033805490910190556122f9878261457d565b6040805182815260208101859052808201849052905186916001600160a01b038a16917f2676e1697cf4731b93ddb4ef54e0e5a98c06cccbbbb2202848a3c6286595e6ce9181900360600190a35050505050505050565b63039ada0090565b601e5481565b60225481565b600061236e611df1565b612376612054565b014210158061239357506123886120fc565b61239061120c565b10155b905090565b336000818152602c60209081526040808320838052909152902054819080612407576040805162461bcd60e51b815260206004820152600a60248201527f6e6f207265776172647300000000000000000000000000000000000000000000604482015290519081900360640190fd5b61240f612364565b612460576040805162461bcd60e51b815260206004820152601c60248201527f6265666f7265206d696e696d756d20756e6c6f636b20706572696f6400000000604482015290519081900360640190fd5b6001600160a01b038084166000908152602c60209081526040808320838052909152808220829055519184169183156108fc0291849190818181858888f193505050501580156124b4573d6000803e3d6000fd5b50816001600160a01b0316836001600160a01b03167f80b36a0e929d7e7925087e54acfeecf4c6043e451b9d71ac5e908b66f9e5d126836040518082815260200191505060405180910390a3505050565b6000546001600160a01b031690565b6000546001600160a01b0316331490565b6001600160a01b0383161580159061254557506001600160a01b03821615155b612596576040805162461bcd60e51b815260206004820152600f60248201527f696e76616c696420616464726573730000000000000000000000000000000000604482015290519081900360640190fd5b611fff838334846148ae565b60008060006125af61552d565b6000806125bd89898961433e565b6040830151602084015193519093019092019750955093505050505b93509350939050565b602b6020908152600091825260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156126755780601f1061264a57610100808354040283529160200191612675565b820191906000526020600020905b81548152906001019060200180831161265857829003601f168201915b505050505081565b60335481565b602f541580159061269b5750601e54600101602f5411155b6126ec576040805162461bcd60e51b815260206004820152601960248201527f6665617475726520776173206e6f742061637469766174656400000000000000604482015290519081900360640190fd5b336126f78183614695565b6127008261462c565b6001600160a01b038116600090815260296020526040902060060154821461276f576040805162461bcd60e51b815260206004820152600e60248201527f77726f6e67207374616b65724944000000000000000000000000000000000000604482015290519081900360640190fd5b62127500831015801561278657506301e133808311155b6127d7576040805162461bcd60e51b815260206004820152601260248201527f696e636f7272656374206475726174696f6e0000000000000000000000000000604482015290519081900360640190fd5b60006127e9428563ffffffff613dd216565b60008481526030602052604090206001015490915081111561283c5760405162461bcd60e51b81526004018080602001828103825260228152602001806156546022913960400191505060405180910390fd5b6001600160a01b038216600090815260316020908152604080832086845290915290206001015481116128b6576040805162461bcd60e51b815260206004820152601160248201527f616c7265616479206c6f636b6564207570000000000000000000000000000000604482015290519081900360640190fd5b6128c08284614106565b6128e9576001600160a01b03821660009081526032602090815260408083208684529091528120555b60405180604001604052806128fc612238565b815260209081018390526001600160a01b03841660008181526031835260408082208883528452902083518155929091015160019092019190915583907f823f252f996e1f519fd0215db7eb4d5a688d78587bf03bfb03d77bfca939806d612962612238565b60408051918252602082018690528051918290030190a350505050565b60275481565b6000606462e4e1c0611b7b565b61010090565b60008215806129e657506000828152602080526040902060050154158015906129cc57506000828152602080526040902054155b80156129e657506000828152602080526040902060040154155b90505b92915050565b6001600160a01b03811660009081526021602052604081205480612a175760009150506120da565b60008181526020805260409020600901546001600160a01b038481169116146129e95760009150506120da565b6000806000612a5161552d565b6000806125bd898989614702565b6000918252601f602090815260408084209284529190529020805460018201546002830154600390930154919390929190565b336000612a9e826129ef565b6000818152602080526040902060040154909150612b03576040805162461bcd60e51b815260206004820152601960248201527f7374616b6572207761736e277420646561637469766174656400000000000000604482015290519081900360640190fd5b612b0b61200a565b600082815260208052604090206004015401421015612b71576040805162461bcd60e51b815260206004820152601660248201527f6e6f7420656e6f7567682074696d652070617373656400000000000000000000604482015290519081900360640190fd5b612b7961180c565b600082815260208052604090206003015401612b93612238565b1015612be6576040805162461bcd60e51b815260206004820152601860248201527f6e6f7420656e6f7567682065706f636873207061737365640000000000000000604482015290519081900360640190fd5b600081815260208080526040808320600881018054600583018054845488865560018087018a9055600287018a9055600387018a9055600487018a905592899055600686018990556007860189905573ffffffffffffffffffffffffffffffffffffffff1980851690955560099095018054909416909355602b9095529285206001600160a01b039093169490939092908216151590612c8690846155d3565b6001600160a01b038088166000908152602160205260408082208290559187168152908120558115612cc357600086815260208052604090208290555b60238054600019019055602454612ce0908563ffffffff613ec716565b60245580612d24576040516001600160a01b0388169085156108fc029086906000818181858888f19350505050158015612d1e573d6000803e3d6000fd5b50612d28565b8392505b602854612d3b908463ffffffff613dd216565b60285560408051848152905187917f8c6548258f8f12a9d4b593fa89a223417ed901d4ee9712ba09beb4d56f5262b6919081900360200190a250505050505050565b6029602052600090815260409020805460018201546002830154600384015460048501546005860154600690960154949593949293919290919087565b612dc43382614bf2565b50565b6001600160a01b038116600090815260296020526040902060040154339015612e37576040805162461bcd60e51b815260206004820152601260248201527f616c72656164792064656c65676174696e670000000000000000000000000000604482015290519081900360640190fd5b816001600160a01b0316816001600160a01b03161415612e9e576040805162461bcd60e51b815260206004820152601060248201527f7468652073616d65206164647265737300000000000000000000000000000000604482015290519081900360640190fd5b6000612ea9826129ef565b9050612eb4816142dc565b6001600160a01b0383166000908152602160205260409020541580612ef057506001600160a01b03831660009081526021602052604090205481145b612f41576040805162461bcd60e51b815260206004820152601460248201527f6164647265737320616c72656164792075736564000000000000000000000000604482015290519081900360640190fd5b60008181526020808052604080832060098101805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03898116918217909255878216808752602186528487208790559086528386208790556008909201541684528184208590558352602c825280832083805290915290205415612ffc576001600160a01b038281166000908152602c60208181526040808420848052808352818520958916855292825280842084805282528320845490555290555b826001600160a01b0316826001600160a01b0316827f7cc102ee500cbca85691c9642080562e8f012b04d27f5b7f389453672b20694660405160405180910390a4505050565b33600061304e826129ef565b9050613059816141b5565b61306281614221565b61306b81614287565b156130bd576040805162461bcd60e51b815260206004820152600f60248201527f7374616b65206973206c6f636b65640000000000000000000000000000000000604482015290519081900360640190fd5b60008181526020805260408120600501546130dd91908390859080613ff6565b6130e5612238565b6000828152602080526040808220600381019390935542600490930192909255905182917ff7c308d0d978cce3aec157d1b34e355db4636b4e71ce91b4f5ec9e7a4f5cdc6091a25050565b613138612514565b613189576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b601e5481116131df576040805162461bcd60e51b815260206004820152601760248201527f63616e277420737461727420696e207468652070617374000000000000000000604482015290519081900360640190fd5b602f5415806131f15750601e54602f54115b613242576040805162461bcd60e51b815260206004820152601360248201527f6665617475726520776173207374617274656400000000000000000000000000604482015290519081900360640190fd5b602f55565b612dc4333334846148ae565b600060646304c4b400611b7b565b600061326c336129ef565b90506132766111ff565b3410156132ca576040805162461bcd60e51b815260206004820152601360248201527f696e73756666696369656e7420616d6f756e7400000000000000000000000000604482015290519081900360640190fd5b6132d38161462c565b60008181526020805260408120600501546132f4903463ffffffff613dd216565b6000838152602080526040902060050181905560245490915061331d903463ffffffff613dd216565b60245560408051828152346020820152815184927fa1d93e9a2a16bf4c2d0cdc6f47fe0fa054c741c96b3dac1297c79eaca31714e9928290030190a25050565b60316020908152600092835260408084209091529082529020805460019091015482565b6030602052600090815260409020805460019091015482565b33600081815260296020526040902060068101546133b88382613f09565b6133c28382613f85565b6133ca6111ff565b84101561341e576040805162461bcd60e51b815260206004820152601060248201527f746f6f20736d616c6c20616d6f756e7400000000000000000000000000000000604482015290519081900360640190fd5b60048201548061342c6111ff565b8601111561346b5760405162461bcd60e51b81526004018080602001828103825260218152602001806156766021913960400191505060405180910390fd5b6001600160a01b0384166000908152602d60209081526040808320898452909152902060030154156134e4576040805162461bcd60e51b815260206004820152601360248201527f7772494420616c72656164792065786973747300000000000000000000000000604482015290519081900360640190fd5b6134f2600183868885613ff6565b60006134fe8584614106565b156135495761350e858488614165565b905085811061351e575060001985015b6001600160a01b03851660009081526032602090815260408083208684529091529020805482900390555b600484018054879003905560008381526020805260409020600501541561359d57600083815260208052604090206007015461358b908763ffffffff613ec716565b60008481526020805260409020600701555b6001600160a01b0385166000908152602d602090815260408083208a845290915290208381558187036003909101556135d4612238565b6001600160a01b0386166000818152602d602090815260408083208c8452825291829020600180820195909555426002820155600401805460ff19168517905581518b81529081019390935282810189905251859282917fde2d2a87af2fa2de55bde86f04143144eb632fa6be266dc224341a371fb8916d9181900360600190a461365f858461205c565b61366883613671565b50505050505050565b61367a816142dc565b600081815260208080526040918290206005810154600790910154835191825291810191909152815183927f509404fa75ce234a1273cf9f7918bcf54e0ef19f2772e4f71b6526606a723b7c928290030190a250565b6136d8612514565b613729576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b612dc481614ed1565b602f541580159061374a5750601e54600101602f5411155b61379b576040805162461bcd60e51b815260206004820152601960248201527f6665617475726520776173206e6f742061637469766174656400000000000000604482015290519081900360640190fd5b60006137a6336129ef565b90506137b18161462c565b6212750082101580156137c857506301e133808211155b613819576040805162461bcd60e51b815260206004820152601260248201527f696e636f7272656374206475726174696f6e0000000000000000000000000000604482015290519081900360640190fd5b613829428363ffffffff613dd216565b6000828152603060205260409020600101541061388d576040805162461bcd60e51b815260206004820152601160248201527f616c7265616479206c6f636b6564207570000000000000000000000000000000604482015290519081900360640190fd5b600061389f428463ffffffff613dd216565b905060405180604001604052806138b4612238565b815260209081018390526000848152603082526040902082518155910151600190910155817f71f8e76b11dde805fa567e857c4beba340500f4ca9da003520a25014f542162b613902612238565b60408051918252602082018590528051918290030190a2505050565b336000818152602d60209081526040808320858452909152902060020154819061398f576040805162461bcd60e51b815260206004820152601560248201527f7265717565737420646f65736e27742065786973740000000000000000000000604482015290519081900360640190fd5b6001600160a01b0382166000908152602d6020908152604080832086845290915290206004810154905460ff909116908180156139db5750600081815260208052604090206005015415155b15613af4576139e861200a565b6001600160a01b0385166000908152602d6020908152604080832089845290915290206002015401421015613a64576040805162461bcd60e51b815260206004820152601660248201527f6e6f7420656e6f7567682074696d652070617373656400000000000000000000604482015290519081900360640190fd5b613a6c61180c565b6001600160a01b0385166000908152602d6020908152604080832089845290915290206001015401613a9c612238565b1015613aef576040805162461bcd60e51b815260206004820152601860248201527f6e6f7420656e6f7567682065706f636873207061737365640000000000000000604482015290519081900360640190fd5b613c08565b81613c0857613b0161200a565b6001600160a01b0385166000908152602d6020908152604080832089845290915290206002015401421015613b7d576040805162461bcd60e51b815260206004820152601660248201527f6e6f7420656e6f7567682074696d652070617373656400000000000000000000604482015290519081900360640190fd5b613b8561180c565b6001600160a01b0385166000908152602d6020908152604080832089845290915290206001015401613bb5612238565b1015613c08576040805162461bcd60e51b815260206004820152601860248201527f6e6f7420656e6f7567682065706f636873207061737365640000000000000000604482015290519081900360640190fd5b600081815260208080526040808320546001600160a01b0388168452602d83528184208985529092528220600381018054848355600180840186905560028401869055918590556004909201805460ff191690559091161515908415613c8357602654613c7b908263ffffffff613ec716565b602655613c9a565b602454613c96908263ffffffff613ec716565b6024555b81613cdb576040516001600160a01b0387169082156108fc029083906000818181858888f19350505050158015613cd5573d6000803e3d6000fd5b50613cdf565b8092505b8415613d0057602754613cf8908463ffffffff613dd216565b602755613d17565b602854613d13908463ffffffff613dd216565b6028555b604080518981528615156020820152808201859052905185916001600160a01b03808a1692908b16917fd5304dabc5bd47105b6921889d1b528c4b2223250248a916afd129b1c0512ddd919081900360600190a45050505050505050565b60208052600090815260409020805460018201546002830154600384015460048501546005860154600687015460078801546008890154600990990154979896979596949593949293919290916001600160a01b0391821691168a565b6000828201838110156129e6576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b600082613e3b575060006129e9565b82820282848281613e4857fe5b04146129e65760405162461bcd60e51b81526004018080602001828103825260218152602001806156976021913960400191505060405180910390fd5b60006129e683836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250614f7e565b60006129e683836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250615020565b613f138282614695565b6001600160a01b03821660009081526029602052604090206003015415613f81576040805162461bcd60e51b815260206004820152601960248201527f64656c65676174696f6e20697320646561637469766174656400000000000000604482015290519081900360640190fd5b5050565b601e546001600160a01b03831660009081526029602052604090206005015414613f81576040805162461bcd60e51b815260206004820152601760248201527f6e6f7420616c6c207265776172647320636c61696d6564000000000000000000604482015290519081900360640190fd5b6140008585612998565b15611805576000614017828463ffffffff613ec716565b6001600160a01b0385166000908152602c6020908152604080832083805290915281205491925061405284611274848663ffffffff613e2c16565b905080614082576001600160a01b0386166000908152602c602090815260408083208380529091528120556140a9565b6001600160a01b0386166000908152602c6020908152604080832083805290915290208190555b8181146140fc576040805189151581528284036020820152815189926001600160a01b038a16927f0ea92567e76d40ddc52d2c1d74a521a59329a38b50411451de6ad2e565466d0f929081900390910190a35b5050505050505050565b6001600160a01b0382166000908152603160209081526040808320848452909152812060010154158015906129e65750506001600160a01b03919091166000908152603160209081526040808320938352929052206001015442111590565b6001600160a01b038316600090815260296020908152604080832060040154603283528184208685529092528220546141aa908290611274908663ffffffff613e2c16565b9150505b9392505050565b6141be816142dc565b600081815260208052604090206004015415612dc4576040805162461bcd60e51b815260206004820152601560248201527f7374616b65722069732064656163746976617465640000000000000000000000604482015290519081900360640190fd5b601e54600082815260208052604090206006015414612dc4576040805162461bcd60e51b815260206004820152601760248201527f6e6f7420616c6c207265776172647320636c61696d6564000000000000000000604482015290519081900360640190fd5b600081815260306020526040812060010154158015906129e957505060009081526030602052604090206001015442111590565b60006129e9620f42406112746142cf61185f565b859063ffffffff613e2c16565b6000818152602080526040902060050154612dc4576040805162461bcd60e51b815260206004820152601460248201527f7374616b657220646f65736e2774206578697374000000000000000000000000604482015290519081900360640190fd5b61434661552d565b6000848152602080526040812060060154819061436790869060010161507a565b600087815260208052604090206006015490955085116143af5750506040805160808101825260008082526020820181905291810182905260608101829052915083906125d9565b6143b761552d565b60405180608001604052806000815260200160008152602001600081526020016000815250905060008690505b601e5481111580156143f7575085870181105b156144565761440461552d565b6144168983614411612985565b61508f565b80518451018452602080820151908501805190910190526040808201519085018051909101905260609081015190840180519091019052506001016143e4565b60008782116144675750600061446e565b5060001981015b91989697509095945050505050565b8183106144d1576040805162461bcd60e51b815260206004820152601560248201527f65706f636820697320616c726561647920706169640000000000000000000000604482015290519081900360640190fd5b601e54821115614528576040805162461bcd60e51b815260206004820152600c60248201527f6675747572652065706f63680000000000000000000000000000000000000000604482015290519081900360640190fd5b81811015611fff576040805162461bcd60e51b815260206004820152601160248201527f6e6f2065706f63687320636c61696d6564000000000000000000000000000000604482015290519081900360640190fd5b8061458757613f81565b61458f612364565b156145d0576040516001600160a01b0383169082156108fc029083906000818181858888f193505050501580156145ca573d6000803e3d6000fd5b50613f81565b6001600160a01b0382166000908152602c60209081526040808320838052909152902054614604908263ffffffff613dd216565b6001600160a01b0383166000908152602c602090815260408083208380529091529020555050565b614635816141b5565b600081815260208052604090205415612dc4576040805162461bcd60e51b815260206004820152601760248201527f7374616b65722073686f756c6420626520616374697665000000000000000000604482015290519081900360640190fd5b6001600160a01b038216600090815260296020526040902060040154613f81576040805162461bcd60e51b815260206004820152601860248201527f64656c65676174696f6e20646f65736e27742065786973740000000000000000604482015290519081900360640190fd5b61470a61552d565b6000806147156154f0565b506001600160a01b038616600090815260296020908152604091829020825160e08101845281548152600180830154938201939093526002820154938101939093526003810154606084015260048101546080840152600581015460a0840181905260069091015460c08401819052916147919189910161507a565b965081606001516000146147a157fe5b868260a00151106147dd5750506040805160808101825260008082526020820181905291810182905260608101829052935085925090506125d9565b6147e561552d565b60405180608001604052806000815260200160008152602001600081526020016000815250905060008890505b601e548111158015614825575087890181105b156148855761483261552d565b6148458b8584614840612985565b6151c8565b8051845101845260208082015190850180519091019052604080820151908501805190910190526060908101519084018051909101905250600101614812565b60008982116148965750600061489d565b5060001981015b919a98995090979650505050505050565b6001600160a01b0384166000908152602160205260409020541580156148ea57506001600160a01b038316600090815260216020526040902054155b61493b576040805162461bcd60e51b815260206004820152601560248201527f7374616b657220616c7265616479206578697374730000000000000000000000604482015290519081900360640190fd5b6001600160a01b038416600090815260296020526040902060040154156149a9576040805162461bcd60e51b815260206004820152601260248201527f616c72656164792064656c65676174696e670000000000000000000000000000604482015290519081900360640190fd5b6001600160a01b03831660009081526029602052604090206004015415614a17576040805162461bcd60e51b815260206004820152601260248201527f616c72656164792064656c65676174696e670000000000000000000000000000604482015290519081900360640190fd5b614a1f611de2565b821015614a73576040805162461bcd60e51b815260206004820152601360248201527f696e73756666696369656e7420616d6f756e7400000000000000000000000000604482015290519081900360640190fd5b60228054600101908190556001600160a01b0380861660009081526021602090815260408083208590559287168252828220849055838252805220600501839055614abc612238565b600082815260208052604090206001808201929092554260028201556008810180546001600160a01b03808a1673ffffffffffffffffffffffffffffffffffffffff199283161790925560098301805492891692909116919091179055601e54600690910155602380549091019055602454614b3e908463ffffffff613dd216565b6024556040805184815290516001600160a01b0387169183917f0697dfe5062b9db8108e4b31254f47a912ae6bbb78837667b2e923a6f5160d399181900360200190a3815115614b9157614b9182611d20565b836001600160a01b0316856001600160a01b03161461180557836001600160a01b0316856001600160a01b0316827f7cc102ee500cbca85691c9642080562e8f012b04d27f5b7f389453672b20694660405160405180910390a45050505050565b614bfb8161462c565b614c036111ff565b341015614c57576040805162461bcd60e51b815260206004820152601360248201527f696e73756666696369656e7420616d6f756e7400000000000000000000000000604482015290519081900360640190fd5b6001600160a01b03821660009081526029602052604090206004015415614cc5576040805162461bcd60e51b815260206004820152601960248201527f64656c65676174696f6e20616c72656164792065786973747300000000000000604482015290519081900360640190fd5b6001600160a01b03821660009081526021602052604090205415614d30576040805162461bcd60e51b815260206004820152600f60248201527f616c7265616479207374616b696e670000000000000000000000000000000000604482015290519081900360640190fd5b6000818152602080526040902060070154614d51903463ffffffff613dd216565b6000828152602080526040902060050154614d6b906142bb565b1015614dbe576040805162461bcd60e51b815260206004820152601a60248201527f7374616b65722773206c696d6974206973206578636565646564000000000000604482015290519081900360640190fd5b614dc66154f0565b614dce612238565b8152426020808301918252346080840181815260c08501868152601e5460a087019081526001600160a01b038916600090815260298652604080822089518155975160018901558089015160028901556060890151600389015593516004880155905160058701559051600690950194909455858452918052912060070154614e5c9163ffffffff613dd216565b6000838152602080526040902060070155602580546001019055602654614e89903463ffffffff613dd216565b60265560408051348152905183916001600160a01b038616917ffd8c857fb9acd6f4ad59b8621a2a77825168b7b4b76de9586d08e00d4ed462be9181900360200190a3505050565b6001600160a01b038116614f165760405162461bcd60e51b815260040180806020018281038252602681526020018061562e6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b6000818361500a5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015614fcf578181015183820152602001614fb7565b50505050905090810190601f168015614ffc5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083858161501657fe5b0495945050505050565b600081848411156150725760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315614fcf578181015183820152602001614fb7565b505050900390565b6000826150885750806129e9565b5090919050565b61509761552d565b6000806150a48686615339565b6000868152601f602090815260408083208a84529091528120805460019091015492935091906150d48383613dd2565b90508061510c5760405180608001604052806000815260200160008152602001600081526020016000815250955050505050506141ae565b6000615135615128620f4240611274868c63ffffffff613e2c16565b859063ffffffff613dd216565b905061514b82611274878463ffffffff613e2c16565b95505050505050600080602f541180156151675750602f548510155b6000878152603060205260408120549192509086108015906151b0575060001986016000908152601f602090815260408083206001908101548b85526030909352922090910154115b90506151bd838383615429565b979650505050505050565b6151d061552d565b6000806151dd8686615339565b6000868152601f602090815260408083208a845290915281208054600190910154929350919061520d8383613dd2565b905080615245576040518060800160405280600081526020016000815260200160008152602001600081525095505050505050615331565b6001600160a01b038a166000908152602960205260408120600401549061527c620f42406112746142cf828d63ffffffff613ec716565b905061529283611274888463ffffffff613e2c16565b9650505050505050600080602f541180156152af5750602f548510155b6001600160a01b03881660009081526031602090815260408083208a845290915281205491925090861080159061531e575060001986016000908152601f602090815260408083206001908101546001600160a01b038d168552603184528285208c8652909352922090910154115b905061532b838383615429565b93505050505b949350505050565b6000818152601f6020908152604080832060048101548685529281905290832060028101546005909201546003909101548483156153b3576000878152601f6020526040812060068101546002909101546153999163ffffffff613e2c16565b90506153af86611274838863ffffffff613e2c16565b9150505b6000821561540c576000888152601f60205260409020600301546153e3908590611274908663ffffffff613e2c16565b9050615409620f42406112746153f7611b71565b8490620f42400363ffffffff613e2c16565b90505b61541c828263ffffffff613dd216565b9998505050505050505050565b61543161552d565b60405180608001604052806000815260200160008152602001600081526020016000815250905082156154c257811561549a5760008152615487620f424061127461547a611b71565b879063ffffffff613e2c16565b60208201819052840360408201526154bd565b6154ac620f424061127461547a611b71565b815260006020820181905260408201525b6154d4565b83815260006020820181905260408201525b6040810151602082015182518603030360608201529392505050565b6040518060e00160405280600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b6040518060800160405280600081526020016000815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061559657805160ff19168380011785556155c3565b828001600101855582156155c3579182015b828111156155c35782518255916020019190600101906155a8565b506155cf929150615613565b5090565b50805460018160011615610100020316600290046000825580601f106155f95750612dc4565b601f016020900490600052602060002090810190612dc491905b61120991905b808211156155cf576000815560010161561956fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573737374616b65722773206c6f636b696e672077696c6c2066696e6973682066697273746d757374206c65617665206174206c65617374206d696e44656c65676174696f6e536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a265627a7a7231582013b98e647eb33b45d0f93ce48ff61dec2330e3ea7536473513087419baac802f64736f6c634300050c0032"

// DeployStakers deploys a new Ethereum contract, binding an instance of Stakers to it.
func DeployStakers(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Stakers, error) {
	parsed, err := abi.JSON(strings.NewReader(StakersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StakersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stakers{StakersCaller: StakersCaller{contract: contract}, StakersTransactor: StakersTransactor{contract: contract}, StakersFilterer: StakersFilterer{contract: contract}}, nil
}

// Stakers is an auto generated Go binding around an Ethereum contract.
type Stakers struct {
	StakersCaller     // Read-only binding to the contract
	StakersTransactor // Write-only binding to the contract
	StakersFilterer   // Log filterer for contract events
}

// StakersCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakersSession struct {
	Contract     *Stakers          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakersCallerSession struct {
	Contract *StakersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakersTransactorSession struct {
	Contract     *StakersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakersRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakersRaw struct {
	Contract *Stakers // Generic contract binding to access the raw methods on
}

// StakersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakersCallerRaw struct {
	Contract *StakersCaller // Generic read-only contract binding to access the raw methods on
}

// StakersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakersTransactorRaw struct {
	Contract *StakersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakers creates a new instance of Stakers, bound to a specific deployed contract.
func NewStakers(address common.Address, backend bind.ContractBackend) (*Stakers, error) {
	contract, err := bindStakers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stakers{StakersCaller: StakersCaller{contract: contract}, StakersTransactor: StakersTransactor{contract: contract}, StakersFilterer: StakersFilterer{contract: contract}}, nil
}

// NewStakersCaller creates a new read-only instance of Stakers, bound to a specific deployed contract.
func NewStakersCaller(address common.Address, caller bind.ContractCaller) (*StakersCaller, error) {
	contract, err := bindStakers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakersCaller{contract: contract}, nil
}

// NewStakersTransactor creates a new write-only instance of Stakers, bound to a specific deployed contract.
func NewStakersTransactor(address common.Address, transactor bind.ContractTransactor) (*StakersTransactor, error) {
	contract, err := bindStakers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakersTransactor{contract: contract}, nil
}

// NewStakersFilterer creates a new log filterer instance of Stakers, bound to a specific deployed contract.
func NewStakersFilterer(address common.Address, filterer bind.ContractFilterer) (*StakersFilterer, error) {
	contract, err := bindStakers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakersFilterer{contract: contract}, nil
}

// bindStakers binds a generic wrapper to an already deployed contract.
func bindStakers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakers *StakersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stakers.Contract.StakersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakers *StakersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakers.Contract.StakersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakers *StakersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakers.Contract.StakersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakers *StakersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stakers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakers *StakersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakers *StakersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakers.Contract.contract.Transact(opts, method, params...)
}

// RewardsBurnableOnDeactivation is a free data retrieval call binding the contract method 0xb1a3ebfa.
//
// Solidity: function _rewardsBurnableOnDeactivation(bool isDelegation, uint256 stakerID) view returns(bool)
func (_Stakers *StakersCaller) RewardsBurnableOnDeactivation(opts *bind.CallOpts, isDelegation bool, stakerID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "_rewardsBurnableOnDeactivation", isDelegation, stakerID)
	return *ret0, err
}

// RewardsBurnableOnDeactivation is a free data retrieval call binding the contract method 0xb1a3ebfa.
//
// Solidity: function _rewardsBurnableOnDeactivation(bool isDelegation, uint256 stakerID) view returns(bool)
func (_Stakers *StakersSession) RewardsBurnableOnDeactivation(isDelegation bool, stakerID *big.Int) (bool, error) {
	return _Stakers.Contract.RewardsBurnableOnDeactivation(&_Stakers.CallOpts, isDelegation, stakerID)
}

// RewardsBurnableOnDeactivation is a free data retrieval call binding the contract method 0xb1a3ebfa.
//
// Solidity: function _rewardsBurnableOnDeactivation(bool isDelegation, uint256 stakerID) view returns(bool)
func (_Stakers *StakersCallerSession) RewardsBurnableOnDeactivation(isDelegation bool, stakerID *big.Int) (bool, error) {
	return _Stakers.Contract.RewardsBurnableOnDeactivation(&_Stakers.CallOpts, isDelegation, stakerID)
}

// SfcAddressToStakerID is a free data retrieval call binding the contract method 0xb42cb58d.
//
// Solidity: function _sfcAddressToStakerID(address sfcAddress) view returns(uint256)
func (_Stakers *StakersCaller) SfcAddressToStakerID(opts *bind.CallOpts, sfcAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "_sfcAddressToStakerID", sfcAddress)
	return *ret0, err
}

// SfcAddressToStakerID is a free data retrieval call binding the contract method 0xb42cb58d.
//
// Solidity: function _sfcAddressToStakerID(address sfcAddress) view returns(uint256)
func (_Stakers *StakersSession) SfcAddressToStakerID(sfcAddress common.Address) (*big.Int, error) {
	return _Stakers.Contract.SfcAddressToStakerID(&_Stakers.CallOpts, sfcAddress)
}

// SfcAddressToStakerID is a free data retrieval call binding the contract method 0xb42cb58d.
//
// Solidity: function _sfcAddressToStakerID(address sfcAddress) view returns(uint256)
func (_Stakers *StakersCallerSession) SfcAddressToStakerID(sfcAddress common.Address) (*big.Int, error) {
	return _Stakers.Contract.SfcAddressToStakerID(&_Stakers.CallOpts, sfcAddress)
}

// BondedRatio is a free data retrieval call binding the contract method 0x041d97a3.
//
// Solidity: function bondedRatio() view returns(uint256)
func (_Stakers *StakersCaller) BondedRatio(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "bondedRatio")
	return *ret0, err
}

// BondedRatio is a free data retrieval call binding the contract method 0x041d97a3.
//
// Solidity: function bondedRatio() view returns(uint256)
func (_Stakers *StakersSession) BondedRatio() (*big.Int, error) {
	return _Stakers.Contract.BondedRatio(&_Stakers.CallOpts)
}

// BondedRatio is a free data retrieval call binding the contract method 0x041d97a3.
//
// Solidity: function bondedRatio() view returns(uint256)
func (_Stakers *StakersCallerSession) BondedRatio() (*big.Int, error) {
	return _Stakers.Contract.BondedRatio(&_Stakers.CallOpts)
}

// BondedTargetPeriod is a free data retrieval call binding the contract method 0x7b8c6b02.
//
// Solidity: function bondedTargetPeriod() pure returns(uint256)
func (_Stakers *StakersCaller) BondedTargetPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "bondedTargetPeriod")
	return *ret0, err
}

// BondedTargetPeriod is a free data retrieval call binding the contract method 0x7b8c6b02.
//
// Solidity: function bondedTargetPeriod() pure returns(uint256)
func (_Stakers *StakersSession) BondedTargetPeriod() (*big.Int, error) {
	return _Stakers.Contract.BondedTargetPeriod(&_Stakers.CallOpts)
}

// BondedTargetPeriod is a free data retrieval call binding the contract method 0x7b8c6b02.
//
// Solidity: function bondedTargetPeriod() pure returns(uint256)
func (_Stakers *StakersCallerSession) BondedTargetPeriod() (*big.Int, error) {
	return _Stakers.Contract.BondedTargetPeriod(&_Stakers.CallOpts)
}

// BondedTargetRewardUnlock is a free data retrieval call binding the contract method 0x6a1cf400.
//
// Solidity: function bondedTargetRewardUnlock() view returns(uint256)
func (_Stakers *StakersCaller) BondedTargetRewardUnlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "bondedTargetRewardUnlock")
	return *ret0, err
}

// BondedTargetRewardUnlock is a free data retrieval call binding the contract method 0x6a1cf400.
//
// Solidity: function bondedTargetRewardUnlock() view returns(uint256)
func (_Stakers *StakersSession) BondedTargetRewardUnlock() (*big.Int, error) {
	return _Stakers.Contract.BondedTargetRewardUnlock(&_Stakers.CallOpts)
}

// BondedTargetRewardUnlock is a free data retrieval call binding the contract method 0x6a1cf400.
//
// Solidity: function bondedTargetRewardUnlock() view returns(uint256)
func (_Stakers *StakersCallerSession) BondedTargetRewardUnlock() (*big.Int, error) {
	return _Stakers.Contract.BondedTargetRewardUnlock(&_Stakers.CallOpts)
}

// BondedTargetStart is a free data retrieval call binding the contract method 0xce5aa000.
//
// Solidity: function bondedTargetStart() pure returns(uint256)
func (_Stakers *StakersCaller) BondedTargetStart(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "bondedTargetStart")
	return *ret0, err
}

// BondedTargetStart is a free data retrieval call binding the contract method 0xce5aa000.
//
// Solidity: function bondedTargetStart() pure returns(uint256)
func (_Stakers *StakersSession) BondedTargetStart() (*big.Int, error) {
	return _Stakers.Contract.BondedTargetStart(&_Stakers.CallOpts)
}

// BondedTargetStart is a free data retrieval call binding the contract method 0xce5aa000.
//
// Solidity: function bondedTargetStart() pure returns(uint256)
func (_Stakers *StakersCallerSession) BondedTargetStart() (*big.Int, error) {
	return _Stakers.Contract.BondedTargetStart(&_Stakers.CallOpts)
}

// CalcDelegationRewards is a free data retrieval call binding the contract method 0xb6732f3d.
//
// Solidity: function calcDelegationRewards(address delegator, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_Stakers *StakersCaller) CalcDelegationRewards(opts *bind.CallOpts, delegator common.Address, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Stakers.contract.Call(opts, out, "calcDelegationRewards", delegator, _fromEpoch, maxEpochs)
	return *ret0, *ret1, *ret2, err
}

// CalcDelegationRewards is a free data retrieval call binding the contract method 0xb6732f3d.
//
// Solidity: function calcDelegationRewards(address delegator, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_Stakers *StakersSession) CalcDelegationRewards(delegator common.Address, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Stakers.Contract.CalcDelegationRewards(&_Stakers.CallOpts, delegator, _fromEpoch, maxEpochs)
}

// CalcDelegationRewards is a free data retrieval call binding the contract method 0xb6732f3d.
//
// Solidity: function calcDelegationRewards(address delegator, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_Stakers *StakersCallerSession) CalcDelegationRewards(delegator common.Address, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Stakers.Contract.CalcDelegationRewards(&_Stakers.CallOpts, delegator, _fromEpoch, maxEpochs)
}

// CalcValidatorRewards is a free data retrieval call binding the contract method 0x96060e71.
//
// Solidity: function calcValidatorRewards(uint256 stakerID, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_Stakers *StakersCaller) CalcValidatorRewards(opts *bind.CallOpts, stakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Stakers.contract.Call(opts, out, "calcValidatorRewards", stakerID, _fromEpoch, maxEpochs)
	return *ret0, *ret1, *ret2, err
}

// CalcValidatorRewards is a free data retrieval call binding the contract method 0x96060e71.
//
// Solidity: function calcValidatorRewards(uint256 stakerID, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_Stakers *StakersSession) CalcValidatorRewards(stakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Stakers.Contract.CalcValidatorRewards(&_Stakers.CallOpts, stakerID, _fromEpoch, maxEpochs)
}

// CalcValidatorRewards is a free data retrieval call binding the contract method 0x96060e71.
//
// Solidity: function calcValidatorRewards(uint256 stakerID, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_Stakers *StakersCallerSession) CalcValidatorRewards(stakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Stakers.Contract.CalcValidatorRewards(&_Stakers.CallOpts, stakerID, _fromEpoch, maxEpochs)
}

// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
//
// Solidity: function contractCommission() pure returns(uint256)
func (_Stakers *StakersCaller) ContractCommission(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "contractCommission")
	return *ret0, err
}

// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
//
// Solidity: function contractCommission() pure returns(uint256)
func (_Stakers *StakersSession) ContractCommission() (*big.Int, error) {
	return _Stakers.Contract.ContractCommission(&_Stakers.CallOpts)
}

// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
//
// Solidity: function contractCommission() pure returns(uint256)
func (_Stakers *StakersCallerSession) ContractCommission() (*big.Int, error) {
	return _Stakers.Contract.ContractCommission(&_Stakers.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Stakers *StakersCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "currentEpoch")
	return *ret0, err
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Stakers *StakersSession) CurrentEpoch() (*big.Int, error) {
	return _Stakers.Contract.CurrentEpoch(&_Stakers.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Stakers *StakersCallerSession) CurrentEpoch() (*big.Int, error) {
	return _Stakers.Contract.CurrentEpoch(&_Stakers.CallOpts)
}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_Stakers *StakersCaller) CurrentSealedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "currentSealedEpoch")
	return *ret0, err
}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_Stakers *StakersSession) CurrentSealedEpoch() (*big.Int, error) {
	return _Stakers.Contract.CurrentSealedEpoch(&_Stakers.CallOpts)
}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_Stakers *StakersCallerSession) CurrentSealedEpoch() (*big.Int, error) {
	return _Stakers.Contract.CurrentSealedEpoch(&_Stakers.CallOpts)
}

// DelegationEarlyWithdrawalPenalty is a free data retrieval call binding the contract method 0x66799a54.
//
// Solidity: function delegationEarlyWithdrawalPenalty(address , uint256 ) view returns(uint256)
func (_Stakers *StakersCaller) DelegationEarlyWithdrawalPenalty(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "delegationEarlyWithdrawalPenalty", arg0, arg1)
	return *ret0, err
}

// DelegationEarlyWithdrawalPenalty is a free data retrieval call binding the contract method 0x66799a54.
//
// Solidity: function delegationEarlyWithdrawalPenalty(address , uint256 ) view returns(uint256)
func (_Stakers *StakersSession) DelegationEarlyWithdrawalPenalty(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Stakers.Contract.DelegationEarlyWithdrawalPenalty(&_Stakers.CallOpts, arg0, arg1)
}

// DelegationEarlyWithdrawalPenalty is a free data retrieval call binding the contract method 0x66799a54.
//
// Solidity: function delegationEarlyWithdrawalPenalty(address , uint256 ) view returns(uint256)
func (_Stakers *StakersCallerSession) DelegationEarlyWithdrawalPenalty(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Stakers.Contract.DelegationEarlyWithdrawalPenalty(&_Stakers.CallOpts, arg0, arg1)
}

// DelegationLockPeriodEpochs is a free data retrieval call binding the contract method 0x1d58179c.
//
// Solidity: function delegationLockPeriodEpochs() pure returns(uint256)
func (_Stakers *StakersCaller) DelegationLockPeriodEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "delegationLockPeriodEpochs")
	return *ret0, err
}

// DelegationLockPeriodEpochs is a free data retrieval call binding the contract method 0x1d58179c.
//
// Solidity: function delegationLockPeriodEpochs() pure returns(uint256)
func (_Stakers *StakersSession) DelegationLockPeriodEpochs() (*big.Int, error) {
	return _Stakers.Contract.DelegationLockPeriodEpochs(&_Stakers.CallOpts)
}

// DelegationLockPeriodEpochs is a free data retrieval call binding the contract method 0x1d58179c.
//
// Solidity: function delegationLockPeriodEpochs() pure returns(uint256)
func (_Stakers *StakersCallerSession) DelegationLockPeriodEpochs() (*big.Int, error) {
	return _Stakers.Contract.DelegationLockPeriodEpochs(&_Stakers.CallOpts)
}

// DelegationLockPeriodTime is a free data retrieval call binding the contract method 0xec6a7f1c.
//
// Solidity: function delegationLockPeriodTime() pure returns(uint256)
func (_Stakers *StakersCaller) DelegationLockPeriodTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "delegationLockPeriodTime")
	return *ret0, err
}

// DelegationLockPeriodTime is a free data retrieval call binding the contract method 0xec6a7f1c.
//
// Solidity: function delegationLockPeriodTime() pure returns(uint256)
func (_Stakers *StakersSession) DelegationLockPeriodTime() (*big.Int, error) {
	return _Stakers.Contract.DelegationLockPeriodTime(&_Stakers.CallOpts)
}

// DelegationLockPeriodTime is a free data retrieval call binding the contract method 0xec6a7f1c.
//
// Solidity: function delegationLockPeriodTime() pure returns(uint256)
func (_Stakers *StakersCallerSession) DelegationLockPeriodTime() (*big.Int, error) {
	return _Stakers.Contract.DelegationLockPeriodTime(&_Stakers.CallOpts)
}

// Delegations is a free data retrieval call binding the contract method 0xbffe3486.
//
// Solidity: function delegations(address ) view returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_Stakers *StakersCaller) Delegations(opts *bind.CallOpts, arg0 common.Address) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	ret := new(struct {
		CreatedEpoch     *big.Int
		CreatedTime      *big.Int
		DeactivatedEpoch *big.Int
		DeactivatedTime  *big.Int
		Amount           *big.Int
		PaidUntilEpoch   *big.Int
		ToStakerID       *big.Int
	})
	out := ret
	err := _Stakers.contract.Call(opts, out, "delegations", arg0)
	return *ret, err
}

// Delegations is a free data retrieval call binding the contract method 0xbffe3486.
//
// Solidity: function delegations(address ) view returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_Stakers *StakersSession) Delegations(arg0 common.Address) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	return _Stakers.Contract.Delegations(&_Stakers.CallOpts, arg0)
}

// Delegations is a free data retrieval call binding the contract method 0xbffe3486.
//
// Solidity: function delegations(address ) view returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_Stakers *StakersCallerSession) Delegations(arg0 common.Address) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	return _Stakers.Contract.Delegations(&_Stakers.CallOpts, arg0)
}

// DelegationsNum is a free data retrieval call binding the contract method 0x4bd202dc.
//
// Solidity: function delegationsNum() view returns(uint256)
func (_Stakers *StakersCaller) DelegationsNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "delegationsNum")
	return *ret0, err
}

// DelegationsNum is a free data retrieval call binding the contract method 0x4bd202dc.
//
// Solidity: function delegationsNum() view returns(uint256)
func (_Stakers *StakersSession) DelegationsNum() (*big.Int, error) {
	return _Stakers.Contract.DelegationsNum(&_Stakers.CallOpts)
}

// DelegationsNum is a free data retrieval call binding the contract method 0x4bd202dc.
//
// Solidity: function delegationsNum() view returns(uint256)
func (_Stakers *StakersCallerSession) DelegationsNum() (*big.Int, error) {
	return _Stakers.Contract.DelegationsNum(&_Stakers.CallOpts)
}

// DelegationsTotalAmount is a free data retrieval call binding the contract method 0x30fa9929.
//
// Solidity: function delegationsTotalAmount() view returns(uint256)
func (_Stakers *StakersCaller) DelegationsTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "delegationsTotalAmount")
	return *ret0, err
}

// DelegationsTotalAmount is a free data retrieval call binding the contract method 0x30fa9929.
//
// Solidity: function delegationsTotalAmount() view returns(uint256)
func (_Stakers *StakersSession) DelegationsTotalAmount() (*big.Int, error) {
	return _Stakers.Contract.DelegationsTotalAmount(&_Stakers.CallOpts)
}

// DelegationsTotalAmount is a free data retrieval call binding the contract method 0x30fa9929.
//
// Solidity: function delegationsTotalAmount() view returns(uint256)
func (_Stakers *StakersCallerSession) DelegationsTotalAmount() (*big.Int, error) {
	return _Stakers.Contract.DelegationsTotalAmount(&_Stakers.CallOpts)
}

// EpochSnapshots is a free data retrieval call binding the contract method 0x1e8a6956.
//
// Solidity: function epochSnapshots(uint256 ) view returns(uint256 endTime, uint256 duration, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 stakeTotalAmount, uint256 delegationsTotalAmount, uint256 totalSupply)
func (_Stakers *StakersCaller) EpochSnapshots(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EndTime                *big.Int
	Duration               *big.Int
	EpochFee               *big.Int
	TotalBaseRewardWeight  *big.Int
	TotalTxRewardWeight    *big.Int
	BaseRewardPerSecond    *big.Int
	StakeTotalAmount       *big.Int
	DelegationsTotalAmount *big.Int
	TotalSupply            *big.Int
}, error) {
	ret := new(struct {
		EndTime                *big.Int
		Duration               *big.Int
		EpochFee               *big.Int
		TotalBaseRewardWeight  *big.Int
		TotalTxRewardWeight    *big.Int
		BaseRewardPerSecond    *big.Int
		StakeTotalAmount       *big.Int
		DelegationsTotalAmount *big.Int
		TotalSupply            *big.Int
	})
	out := ret
	err := _Stakers.contract.Call(opts, out, "epochSnapshots", arg0)
	return *ret, err
}

// EpochSnapshots is a free data retrieval call binding the contract method 0x1e8a6956.
//
// Solidity: function epochSnapshots(uint256 ) view returns(uint256 endTime, uint256 duration, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 stakeTotalAmount, uint256 delegationsTotalAmount, uint256 totalSupply)
func (_Stakers *StakersSession) EpochSnapshots(arg0 *big.Int) (struct {
	EndTime                *big.Int
	Duration               *big.Int
	EpochFee               *big.Int
	TotalBaseRewardWeight  *big.Int
	TotalTxRewardWeight    *big.Int
	BaseRewardPerSecond    *big.Int
	StakeTotalAmount       *big.Int
	DelegationsTotalAmount *big.Int
	TotalSupply            *big.Int
}, error) {
	return _Stakers.Contract.EpochSnapshots(&_Stakers.CallOpts, arg0)
}

// EpochSnapshots is a free data retrieval call binding the contract method 0x1e8a6956.
//
// Solidity: function epochSnapshots(uint256 ) view returns(uint256 endTime, uint256 duration, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 stakeTotalAmount, uint256 delegationsTotalAmount, uint256 totalSupply)
func (_Stakers *StakersCallerSession) EpochSnapshots(arg0 *big.Int) (struct {
	EndTime                *big.Int
	Duration               *big.Int
	EpochFee               *big.Int
	TotalBaseRewardWeight  *big.Int
	TotalTxRewardWeight    *big.Int
	BaseRewardPerSecond    *big.Int
	StakeTotalAmount       *big.Int
	DelegationsTotalAmount *big.Int
	TotalSupply            *big.Int
}, error) {
	return _Stakers.Contract.EpochSnapshots(&_Stakers.CallOpts, arg0)
}

// EpochValidator is a free data retrieval call binding the contract method 0xb9029d50.
//
// Solidity: function epochValidator(uint256 e, uint256 v) view returns(uint256 stakeAmount, uint256 delegatedMe, uint256 baseRewardWeight, uint256 txRewardWeight)
func (_Stakers *StakersCaller) EpochValidator(opts *bind.CallOpts, e *big.Int, v *big.Int) (struct {
	StakeAmount      *big.Int
	DelegatedMe      *big.Int
	BaseRewardWeight *big.Int
	TxRewardWeight   *big.Int
}, error) {
	ret := new(struct {
		StakeAmount      *big.Int
		DelegatedMe      *big.Int
		BaseRewardWeight *big.Int
		TxRewardWeight   *big.Int
	})
	out := ret
	err := _Stakers.contract.Call(opts, out, "epochValidator", e, v)
	return *ret, err
}

// EpochValidator is a free data retrieval call binding the contract method 0xb9029d50.
//
// Solidity: function epochValidator(uint256 e, uint256 v) view returns(uint256 stakeAmount, uint256 delegatedMe, uint256 baseRewardWeight, uint256 txRewardWeight)
func (_Stakers *StakersSession) EpochValidator(e *big.Int, v *big.Int) (struct {
	StakeAmount      *big.Int
	DelegatedMe      *big.Int
	BaseRewardWeight *big.Int
	TxRewardWeight   *big.Int
}, error) {
	return _Stakers.Contract.EpochValidator(&_Stakers.CallOpts, e, v)
}

// EpochValidator is a free data retrieval call binding the contract method 0xb9029d50.
//
// Solidity: function epochValidator(uint256 e, uint256 v) view returns(uint256 stakeAmount, uint256 delegatedMe, uint256 baseRewardWeight, uint256 txRewardWeight)
func (_Stakers *StakersCallerSession) EpochValidator(e *big.Int, v *big.Int) (struct {
	StakeAmount      *big.Int
	DelegatedMe      *big.Int
	BaseRewardWeight *big.Int
	TxRewardWeight   *big.Int
}, error) {
	return _Stakers.Contract.EpochValidator(&_Stakers.CallOpts, e, v)
}

// FirstLockedUpEpoch is a free data retrieval call binding the contract method 0x6e1a767a.
//
// Solidity: function firstLockedUpEpoch() view returns(uint256)
func (_Stakers *StakersCaller) FirstLockedUpEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "firstLockedUpEpoch")
	return *ret0, err
}

// FirstLockedUpEpoch is a free data retrieval call binding the contract method 0x6e1a767a.
//
// Solidity: function firstLockedUpEpoch() view returns(uint256)
func (_Stakers *StakersSession) FirstLockedUpEpoch() (*big.Int, error) {
	return _Stakers.Contract.FirstLockedUpEpoch(&_Stakers.CallOpts)
}

// FirstLockedUpEpoch is a free data retrieval call binding the contract method 0x6e1a767a.
//
// Solidity: function firstLockedUpEpoch() view returns(uint256)
func (_Stakers *StakersCallerSession) FirstLockedUpEpoch() (*big.Int, error) {
	return _Stakers.Contract.FirstLockedUpEpoch(&_Stakers.CallOpts)
}

// GetStakerID is a free data retrieval call binding the contract method 0x63321e27.
//
// Solidity: function getStakerID(address addr) view returns(uint256)
func (_Stakers *StakersCaller) GetStakerID(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "getStakerID", addr)
	return *ret0, err
}

// GetStakerID is a free data retrieval call binding the contract method 0x63321e27.
//
// Solidity: function getStakerID(address addr) view returns(uint256)
func (_Stakers *StakersSession) GetStakerID(addr common.Address) (*big.Int, error) {
	return _Stakers.Contract.GetStakerID(&_Stakers.CallOpts, addr)
}

// GetStakerID is a free data retrieval call binding the contract method 0x63321e27.
//
// Solidity: function getStakerID(address addr) view returns(uint256)
func (_Stakers *StakersCallerSession) GetStakerID(addr common.Address) (*big.Int, error) {
	return _Stakers.Contract.GetStakerID(&_Stakers.CallOpts, addr)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Stakers *StakersCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Stakers *StakersSession) IsOwner() (bool, error) {
	return _Stakers.Contract.IsOwner(&_Stakers.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Stakers *StakersCallerSession) IsOwner() (bool, error) {
	return _Stakers.Contract.IsOwner(&_Stakers.CallOpts)
}

// LockedDelegations is a free data retrieval call binding the contract method 0xdd099bb6.
//
// Solidity: function lockedDelegations(address , uint256 ) view returns(uint256 fromEpoch, uint256 endTime)
func (_Stakers *StakersCaller) LockedDelegations(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
}, error) {
	ret := new(struct {
		FromEpoch *big.Int
		EndTime   *big.Int
	})
	out := ret
	err := _Stakers.contract.Call(opts, out, "lockedDelegations", arg0, arg1)
	return *ret, err
}

// LockedDelegations is a free data retrieval call binding the contract method 0xdd099bb6.
//
// Solidity: function lockedDelegations(address , uint256 ) view returns(uint256 fromEpoch, uint256 endTime)
func (_Stakers *StakersSession) LockedDelegations(arg0 common.Address, arg1 *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
}, error) {
	return _Stakers.Contract.LockedDelegations(&_Stakers.CallOpts, arg0, arg1)
}

// LockedDelegations is a free data retrieval call binding the contract method 0xdd099bb6.
//
// Solidity: function lockedDelegations(address , uint256 ) view returns(uint256 fromEpoch, uint256 endTime)
func (_Stakers *StakersCallerSession) LockedDelegations(arg0 common.Address, arg1 *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
}, error) {
	return _Stakers.Contract.LockedDelegations(&_Stakers.CallOpts, arg0, arg1)
}

// LockedStakes is a free data retrieval call binding the contract method 0xdf4f49d4.
//
// Solidity: function lockedStakes(uint256 ) view returns(uint256 fromEpoch, uint256 endTime)
func (_Stakers *StakersCaller) LockedStakes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
}, error) {
	ret := new(struct {
		FromEpoch *big.Int
		EndTime   *big.Int
	})
	out := ret
	err := _Stakers.contract.Call(opts, out, "lockedStakes", arg0)
	return *ret, err
}

// LockedStakes is a free data retrieval call binding the contract method 0xdf4f49d4.
//
// Solidity: function lockedStakes(uint256 ) view returns(uint256 fromEpoch, uint256 endTime)
func (_Stakers *StakersSession) LockedStakes(arg0 *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
}, error) {
	return _Stakers.Contract.LockedStakes(&_Stakers.CallOpts, arg0)
}

// LockedStakes is a free data retrieval call binding the contract method 0xdf4f49d4.
//
// Solidity: function lockedStakes(uint256 ) view returns(uint256 fromEpoch, uint256 endTime)
func (_Stakers *StakersCallerSession) LockedStakes(arg0 *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
}, error) {
	return _Stakers.Contract.LockedStakes(&_Stakers.CallOpts, arg0)
}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() pure returns(uint256)
func (_Stakers *StakersCaller) MaxDelegatedRatio(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "maxDelegatedRatio")
	return *ret0, err
}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() pure returns(uint256)
func (_Stakers *StakersSession) MaxDelegatedRatio() (*big.Int, error) {
	return _Stakers.Contract.MaxDelegatedRatio(&_Stakers.CallOpts)
}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() pure returns(uint256)
func (_Stakers *StakersCallerSession) MaxDelegatedRatio() (*big.Int, error) {
	return _Stakers.Contract.MaxDelegatedRatio(&_Stakers.CallOpts)
}

// MaxStakerMetadataSize is a free data retrieval call binding the contract method 0xab2273c0.
//
// Solidity: function maxStakerMetadataSize() pure returns(uint256)
func (_Stakers *StakersCaller) MaxStakerMetadataSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "maxStakerMetadataSize")
	return *ret0, err
}

// MaxStakerMetadataSize is a free data retrieval call binding the contract method 0xab2273c0.
//
// Solidity: function maxStakerMetadataSize() pure returns(uint256)
func (_Stakers *StakersSession) MaxStakerMetadataSize() (*big.Int, error) {
	return _Stakers.Contract.MaxStakerMetadataSize(&_Stakers.CallOpts)
}

// MaxStakerMetadataSize is a free data retrieval call binding the contract method 0xab2273c0.
//
// Solidity: function maxStakerMetadataSize() pure returns(uint256)
func (_Stakers *StakersCallerSession) MaxStakerMetadataSize() (*big.Int, error) {
	return _Stakers.Contract.MaxStakerMetadataSize(&_Stakers.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() pure returns(uint256)
func (_Stakers *StakersCaller) MinDelegation(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "minDelegation")
	return *ret0, err
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() pure returns(uint256)
func (_Stakers *StakersSession) MinDelegation() (*big.Int, error) {
	return _Stakers.Contract.MinDelegation(&_Stakers.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() pure returns(uint256)
func (_Stakers *StakersCallerSession) MinDelegation() (*big.Int, error) {
	return _Stakers.Contract.MinDelegation(&_Stakers.CallOpts)
}

// MinDelegationDecrease is a free data retrieval call binding the contract method 0xcb1c4e67.
//
// Solidity: function minDelegationDecrease() pure returns(uint256)
func (_Stakers *StakersCaller) MinDelegationDecrease(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "minDelegationDecrease")
	return *ret0, err
}

// MinDelegationDecrease is a free data retrieval call binding the contract method 0xcb1c4e67.
//
// Solidity: function minDelegationDecrease() pure returns(uint256)
func (_Stakers *StakersSession) MinDelegationDecrease() (*big.Int, error) {
	return _Stakers.Contract.MinDelegationDecrease(&_Stakers.CallOpts)
}

// MinDelegationDecrease is a free data retrieval call binding the contract method 0xcb1c4e67.
//
// Solidity: function minDelegationDecrease() pure returns(uint256)
func (_Stakers *StakersCallerSession) MinDelegationDecrease() (*big.Int, error) {
	return _Stakers.Contract.MinDelegationDecrease(&_Stakers.CallOpts)
}

// MinDelegationIncrease is a free data retrieval call binding the contract method 0x60c7e37f.
//
// Solidity: function minDelegationIncrease() pure returns(uint256)
func (_Stakers *StakersCaller) MinDelegationIncrease(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "minDelegationIncrease")
	return *ret0, err
}

// MinDelegationIncrease is a free data retrieval call binding the contract method 0x60c7e37f.
//
// Solidity: function minDelegationIncrease() pure returns(uint256)
func (_Stakers *StakersSession) MinDelegationIncrease() (*big.Int, error) {
	return _Stakers.Contract.MinDelegationIncrease(&_Stakers.CallOpts)
}

// MinDelegationIncrease is a free data retrieval call binding the contract method 0x60c7e37f.
//
// Solidity: function minDelegationIncrease() pure returns(uint256)
func (_Stakers *StakersCallerSession) MinDelegationIncrease() (*big.Int, error) {
	return _Stakers.Contract.MinDelegationIncrease(&_Stakers.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_Stakers *StakersCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "minStake")
	return *ret0, err
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_Stakers *StakersSession) MinStake() (*big.Int, error) {
	return _Stakers.Contract.MinStake(&_Stakers.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_Stakers *StakersCallerSession) MinStake() (*big.Int, error) {
	return _Stakers.Contract.MinStake(&_Stakers.CallOpts)
}

// MinStakeDecrease is a free data retrieval call binding the contract method 0x19ddb54f.
//
// Solidity: function minStakeDecrease() pure returns(uint256)
func (_Stakers *StakersCaller) MinStakeDecrease(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "minStakeDecrease")
	return *ret0, err
}

// MinStakeDecrease is a free data retrieval call binding the contract method 0x19ddb54f.
//
// Solidity: function minStakeDecrease() pure returns(uint256)
func (_Stakers *StakersSession) MinStakeDecrease() (*big.Int, error) {
	return _Stakers.Contract.MinStakeDecrease(&_Stakers.CallOpts)
}

// MinStakeDecrease is a free data retrieval call binding the contract method 0x19ddb54f.
//
// Solidity: function minStakeDecrease() pure returns(uint256)
func (_Stakers *StakersCallerSession) MinStakeDecrease() (*big.Int, error) {
	return _Stakers.Contract.MinStakeDecrease(&_Stakers.CallOpts)
}

// MinStakeIncrease is a free data retrieval call binding the contract method 0xc4b5dd7e.
//
// Solidity: function minStakeIncrease() pure returns(uint256)
func (_Stakers *StakersCaller) MinStakeIncrease(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "minStakeIncrease")
	return *ret0, err
}

// MinStakeIncrease is a free data retrieval call binding the contract method 0xc4b5dd7e.
//
// Solidity: function minStakeIncrease() pure returns(uint256)
func (_Stakers *StakersSession) MinStakeIncrease() (*big.Int, error) {
	return _Stakers.Contract.MinStakeIncrease(&_Stakers.CallOpts)
}

// MinStakeIncrease is a free data retrieval call binding the contract method 0xc4b5dd7e.
//
// Solidity: function minStakeIncrease() pure returns(uint256)
func (_Stakers *StakersCallerSession) MinStakeIncrease() (*big.Int, error) {
	return _Stakers.Contract.MinStakeIncrease(&_Stakers.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stakers *StakersCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stakers *StakersSession) Owner() (common.Address, error) {
	return _Stakers.Contract.Owner(&_Stakers.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stakers *StakersCallerSession) Owner() (common.Address, error) {
	return _Stakers.Contract.Owner(&_Stakers.CallOpts)
}

// RewardsAllowed is a free data retrieval call binding the contract method 0x8447c4df.
//
// Solidity: function rewardsAllowed() view returns(bool)
func (_Stakers *StakersCaller) RewardsAllowed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "rewardsAllowed")
	return *ret0, err
}

// RewardsAllowed is a free data retrieval call binding the contract method 0x8447c4df.
//
// Solidity: function rewardsAllowed() view returns(bool)
func (_Stakers *StakersSession) RewardsAllowed() (bool, error) {
	return _Stakers.Contract.RewardsAllowed(&_Stakers.CallOpts)
}

// RewardsAllowed is a free data retrieval call binding the contract method 0x8447c4df.
//
// Solidity: function rewardsAllowed() view returns(bool)
func (_Stakers *StakersCallerSession) RewardsAllowed() (bool, error) {
	return _Stakers.Contract.RewardsAllowed(&_Stakers.CallOpts)
}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address , uint256 ) view returns(uint256 amount)
func (_Stakers *StakersCaller) RewardsStash(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "rewardsStash", arg0, arg1)
	return *ret0, err
}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address , uint256 ) view returns(uint256 amount)
func (_Stakers *StakersSession) RewardsStash(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Stakers.Contract.RewardsStash(&_Stakers.CallOpts, arg0, arg1)
}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address , uint256 ) view returns(uint256 amount)
func (_Stakers *StakersCallerSession) RewardsStash(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Stakers.Contract.RewardsStash(&_Stakers.CallOpts, arg0, arg1)
}

// SlashedDelegationsTotalAmount is a free data retrieval call binding the contract method 0xa70da4d2.
//
// Solidity: function slashedDelegationsTotalAmount() view returns(uint256)
func (_Stakers *StakersCaller) SlashedDelegationsTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "slashedDelegationsTotalAmount")
	return *ret0, err
}

// SlashedDelegationsTotalAmount is a free data retrieval call binding the contract method 0xa70da4d2.
//
// Solidity: function slashedDelegationsTotalAmount() view returns(uint256)
func (_Stakers *StakersSession) SlashedDelegationsTotalAmount() (*big.Int, error) {
	return _Stakers.Contract.SlashedDelegationsTotalAmount(&_Stakers.CallOpts)
}

// SlashedDelegationsTotalAmount is a free data retrieval call binding the contract method 0xa70da4d2.
//
// Solidity: function slashedDelegationsTotalAmount() view returns(uint256)
func (_Stakers *StakersCallerSession) SlashedDelegationsTotalAmount() (*big.Int, error) {
	return _Stakers.Contract.SlashedDelegationsTotalAmount(&_Stakers.CallOpts)
}

// SlashedStakeTotalAmount is a free data retrieval call binding the contract method 0x0a29180c.
//
// Solidity: function slashedStakeTotalAmount() view returns(uint256)
func (_Stakers *StakersCaller) SlashedStakeTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "slashedStakeTotalAmount")
	return *ret0, err
}

// SlashedStakeTotalAmount is a free data retrieval call binding the contract method 0x0a29180c.
//
// Solidity: function slashedStakeTotalAmount() view returns(uint256)
func (_Stakers *StakersSession) SlashedStakeTotalAmount() (*big.Int, error) {
	return _Stakers.Contract.SlashedStakeTotalAmount(&_Stakers.CallOpts)
}

// SlashedStakeTotalAmount is a free data retrieval call binding the contract method 0x0a29180c.
//
// Solidity: function slashedStakeTotalAmount() view returns(uint256)
func (_Stakers *StakersCallerSession) SlashedStakeTotalAmount() (*big.Int, error) {
	return _Stakers.Contract.SlashedStakeTotalAmount(&_Stakers.CallOpts)
}

// StakeLockPeriodEpochs is a free data retrieval call binding the contract method 0x54d77ed2.
//
// Solidity: function stakeLockPeriodEpochs() pure returns(uint256)
func (_Stakers *StakersCaller) StakeLockPeriodEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "stakeLockPeriodEpochs")
	return *ret0, err
}

// StakeLockPeriodEpochs is a free data retrieval call binding the contract method 0x54d77ed2.
//
// Solidity: function stakeLockPeriodEpochs() pure returns(uint256)
func (_Stakers *StakersSession) StakeLockPeriodEpochs() (*big.Int, error) {
	return _Stakers.Contract.StakeLockPeriodEpochs(&_Stakers.CallOpts)
}

// StakeLockPeriodEpochs is a free data retrieval call binding the contract method 0x54d77ed2.
//
// Solidity: function stakeLockPeriodEpochs() pure returns(uint256)
func (_Stakers *StakersCallerSession) StakeLockPeriodEpochs() (*big.Int, error) {
	return _Stakers.Contract.StakeLockPeriodEpochs(&_Stakers.CallOpts)
}

// StakeLockPeriodTime is a free data retrieval call binding the contract method 0x3fee10a8.
//
// Solidity: function stakeLockPeriodTime() pure returns(uint256)
func (_Stakers *StakersCaller) StakeLockPeriodTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "stakeLockPeriodTime")
	return *ret0, err
}

// StakeLockPeriodTime is a free data retrieval call binding the contract method 0x3fee10a8.
//
// Solidity: function stakeLockPeriodTime() pure returns(uint256)
func (_Stakers *StakersSession) StakeLockPeriodTime() (*big.Int, error) {
	return _Stakers.Contract.StakeLockPeriodTime(&_Stakers.CallOpts)
}

// StakeLockPeriodTime is a free data retrieval call binding the contract method 0x3fee10a8.
//
// Solidity: function stakeLockPeriodTime() pure returns(uint256)
func (_Stakers *StakersCallerSession) StakeLockPeriodTime() (*big.Int, error) {
	return _Stakers.Contract.StakeLockPeriodTime(&_Stakers.CallOpts)
}

// StakeTotalAmount is a free data retrieval call binding the contract method 0x3d0317fe.
//
// Solidity: function stakeTotalAmount() view returns(uint256)
func (_Stakers *StakersCaller) StakeTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "stakeTotalAmount")
	return *ret0, err
}

// StakeTotalAmount is a free data retrieval call binding the contract method 0x3d0317fe.
//
// Solidity: function stakeTotalAmount() view returns(uint256)
func (_Stakers *StakersSession) StakeTotalAmount() (*big.Int, error) {
	return _Stakers.Contract.StakeTotalAmount(&_Stakers.CallOpts)
}

// StakeTotalAmount is a free data retrieval call binding the contract method 0x3d0317fe.
//
// Solidity: function stakeTotalAmount() view returns(uint256)
func (_Stakers *StakersCallerSession) StakeTotalAmount() (*big.Int, error) {
	return _Stakers.Contract.StakeTotalAmount(&_Stakers.CallOpts)
}

// StakerMetadata is a free data retrieval call binding the contract method 0x98ec2de5.
//
// Solidity: function stakerMetadata(uint256 ) view returns(bytes)
func (_Stakers *StakersCaller) StakerMetadata(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "stakerMetadata", arg0)
	return *ret0, err
}

// StakerMetadata is a free data retrieval call binding the contract method 0x98ec2de5.
//
// Solidity: function stakerMetadata(uint256 ) view returns(bytes)
func (_Stakers *StakersSession) StakerMetadata(arg0 *big.Int) ([]byte, error) {
	return _Stakers.Contract.StakerMetadata(&_Stakers.CallOpts, arg0)
}

// StakerMetadata is a free data retrieval call binding the contract method 0x98ec2de5.
//
// Solidity: function stakerMetadata(uint256 ) view returns(bytes)
func (_Stakers *StakersCallerSession) StakerMetadata(arg0 *big.Int) ([]byte, error) {
	return _Stakers.Contract.StakerMetadata(&_Stakers.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 stakeAmount, uint256 paidUntilEpoch, uint256 delegatedMe, address dagAddress, address sfcAddress)
func (_Stakers *StakersCaller) Stakers(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
}, error) {
	ret := new(struct {
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
	})
	out := ret
	err := _Stakers.contract.Call(opts, out, "stakers", arg0)
	return *ret, err
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 stakeAmount, uint256 paidUntilEpoch, uint256 delegatedMe, address dagAddress, address sfcAddress)
func (_Stakers *StakersSession) Stakers(arg0 *big.Int) (struct {
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
}, error) {
	return _Stakers.Contract.Stakers(&_Stakers.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 stakeAmount, uint256 paidUntilEpoch, uint256 delegatedMe, address dagAddress, address sfcAddress)
func (_Stakers *StakersCallerSession) Stakers(arg0 *big.Int) (struct {
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
}, error) {
	return _Stakers.Contract.Stakers(&_Stakers.CallOpts, arg0)
}

// StakersLastID is a free data retrieval call binding the contract method 0x81d9dc7a.
//
// Solidity: function stakersLastID() view returns(uint256)
func (_Stakers *StakersCaller) StakersLastID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "stakersLastID")
	return *ret0, err
}

// StakersLastID is a free data retrieval call binding the contract method 0x81d9dc7a.
//
// Solidity: function stakersLastID() view returns(uint256)
func (_Stakers *StakersSession) StakersLastID() (*big.Int, error) {
	return _Stakers.Contract.StakersLastID(&_Stakers.CallOpts)
}

// StakersLastID is a free data retrieval call binding the contract method 0x81d9dc7a.
//
// Solidity: function stakersLastID() view returns(uint256)
func (_Stakers *StakersCallerSession) StakersLastID() (*big.Int, error) {
	return _Stakers.Contract.StakersLastID(&_Stakers.CallOpts)
}

// StakersNum is a free data retrieval call binding the contract method 0x08728f6e.
//
// Solidity: function stakersNum() view returns(uint256)
func (_Stakers *StakersCaller) StakersNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "stakersNum")
	return *ret0, err
}

// StakersNum is a free data retrieval call binding the contract method 0x08728f6e.
//
// Solidity: function stakersNum() view returns(uint256)
func (_Stakers *StakersSession) StakersNum() (*big.Int, error) {
	return _Stakers.Contract.StakersNum(&_Stakers.CallOpts)
}

// StakersNum is a free data retrieval call binding the contract method 0x08728f6e.
//
// Solidity: function stakersNum() view returns(uint256)
func (_Stakers *StakersCallerSession) StakersNum() (*big.Int, error) {
	return _Stakers.Contract.StakersNum(&_Stakers.CallOpts)
}

// TotalBurntLockupRewards is a free data retrieval call binding the contract method 0xa289ad6e.
//
// Solidity: function totalBurntLockupRewards() view returns(uint256)
func (_Stakers *StakersCaller) TotalBurntLockupRewards(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "totalBurntLockupRewards")
	return *ret0, err
}

// TotalBurntLockupRewards is a free data retrieval call binding the contract method 0xa289ad6e.
//
// Solidity: function totalBurntLockupRewards() view returns(uint256)
func (_Stakers *StakersSession) TotalBurntLockupRewards() (*big.Int, error) {
	return _Stakers.Contract.TotalBurntLockupRewards(&_Stakers.CallOpts)
}

// TotalBurntLockupRewards is a free data retrieval call binding the contract method 0xa289ad6e.
//
// Solidity: function totalBurntLockupRewards() view returns(uint256)
func (_Stakers *StakersCallerSession) TotalBurntLockupRewards() (*big.Int, error) {
	return _Stakers.Contract.TotalBurntLockupRewards(&_Stakers.CallOpts)
}

// UnbondingStartDate is a free data retrieval call binding the contract method 0x53a72586.
//
// Solidity: function unbondingStartDate() pure returns(uint256)
func (_Stakers *StakersCaller) UnbondingStartDate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "unbondingStartDate")
	return *ret0, err
}

// UnbondingStartDate is a free data retrieval call binding the contract method 0x53a72586.
//
// Solidity: function unbondingStartDate() pure returns(uint256)
func (_Stakers *StakersSession) UnbondingStartDate() (*big.Int, error) {
	return _Stakers.Contract.UnbondingStartDate(&_Stakers.CallOpts)
}

// UnbondingStartDate is a free data retrieval call binding the contract method 0x53a72586.
//
// Solidity: function unbondingStartDate() pure returns(uint256)
func (_Stakers *StakersCallerSession) UnbondingStartDate() (*big.Int, error) {
	return _Stakers.Contract.UnbondingStartDate(&_Stakers.CallOpts)
}

// UnbondingUnlockPeriod is a free data retrieval call binding the contract method 0x3a0af4d4.
//
// Solidity: function unbondingUnlockPeriod() pure returns(uint256)
func (_Stakers *StakersCaller) UnbondingUnlockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "unbondingUnlockPeriod")
	return *ret0, err
}

// UnbondingUnlockPeriod is a free data retrieval call binding the contract method 0x3a0af4d4.
//
// Solidity: function unbondingUnlockPeriod() pure returns(uint256)
func (_Stakers *StakersSession) UnbondingUnlockPeriod() (*big.Int, error) {
	return _Stakers.Contract.UnbondingUnlockPeriod(&_Stakers.CallOpts)
}

// UnbondingUnlockPeriod is a free data retrieval call binding the contract method 0x3a0af4d4.
//
// Solidity: function unbondingUnlockPeriod() pure returns(uint256)
func (_Stakers *StakersCallerSession) UnbondingUnlockPeriod() (*big.Int, error) {
	return _Stakers.Contract.UnbondingUnlockPeriod(&_Stakers.CallOpts)
}

// UnlockedRewardRatio is a free data retrieval call binding the contract method 0x5e2308d2.
//
// Solidity: function unlockedRewardRatio() pure returns(uint256)
func (_Stakers *StakersCaller) UnlockedRewardRatio(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "unlockedRewardRatio")
	return *ret0, err
}

// UnlockedRewardRatio is a free data retrieval call binding the contract method 0x5e2308d2.
//
// Solidity: function unlockedRewardRatio() pure returns(uint256)
func (_Stakers *StakersSession) UnlockedRewardRatio() (*big.Int, error) {
	return _Stakers.Contract.UnlockedRewardRatio(&_Stakers.CallOpts)
}

// UnlockedRewardRatio is a free data retrieval call binding the contract method 0x5e2308d2.
//
// Solidity: function unlockedRewardRatio() pure returns(uint256)
func (_Stakers *StakersCallerSession) UnlockedRewardRatio() (*big.Int, error) {
	return _Stakers.Contract.UnlockedRewardRatio(&_Stakers.CallOpts)
}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() pure returns(uint256)
func (_Stakers *StakersCaller) ValidatorCommission(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stakers.contract.Call(opts, out, "validatorCommission")
	return *ret0, err
}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() pure returns(uint256)
func (_Stakers *StakersSession) ValidatorCommission() (*big.Int, error) {
	return _Stakers.Contract.ValidatorCommission(&_Stakers.CallOpts)
}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() pure returns(uint256)
func (_Stakers *StakersCallerSession) ValidatorCommission() (*big.Int, error) {
	return _Stakers.Contract.ValidatorCommission(&_Stakers.CallOpts)
}

// WithdrawalRequests is a free data retrieval call binding the contract method 0x4e5a2328.
//
// Solidity: function withdrawalRequests(address , uint256 ) view returns(uint256 stakerID, uint256 epoch, uint256 time, uint256 amount, bool delegation)
func (_Stakers *StakersCaller) WithdrawalRequests(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	StakerID   *big.Int
	Epoch      *big.Int
	Time       *big.Int
	Amount     *big.Int
	Delegation bool
}, error) {
	ret := new(struct {
		StakerID   *big.Int
		Epoch      *big.Int
		Time       *big.Int
		Amount     *big.Int
		Delegation bool
	})
	out := ret
	err := _Stakers.contract.Call(opts, out, "withdrawalRequests", arg0, arg1)
	return *ret, err
}

// WithdrawalRequests is a free data retrieval call binding the contract method 0x4e5a2328.
//
// Solidity: function withdrawalRequests(address , uint256 ) view returns(uint256 stakerID, uint256 epoch, uint256 time, uint256 amount, bool delegation)
func (_Stakers *StakersSession) WithdrawalRequests(arg0 common.Address, arg1 *big.Int) (struct {
	StakerID   *big.Int
	Epoch      *big.Int
	Time       *big.Int
	Amount     *big.Int
	Delegation bool
}, error) {
	return _Stakers.Contract.WithdrawalRequests(&_Stakers.CallOpts, arg0, arg1)
}

// WithdrawalRequests is a free data retrieval call binding the contract method 0x4e5a2328.
//
// Solidity: function withdrawalRequests(address , uint256 ) view returns(uint256 stakerID, uint256 epoch, uint256 time, uint256 amount, bool delegation)
func (_Stakers *StakersCallerSession) WithdrawalRequests(arg0 common.Address, arg1 *big.Int) (struct {
	StakerID   *big.Int
	Epoch      *big.Int
	Time       *big.Int
	Amount     *big.Int
	Delegation bool
}, error) {
	return _Stakers.Contract.WithdrawalRequests(&_Stakers.CallOpts, arg0, arg1)
}

// SyncDelegator is a paid mutator transaction binding the contract method 0x5dc03f1f.
//
// Solidity: function _syncDelegator(address delegator, uint256 stakerID) returns()
func (_Stakers *StakersTransactor) SyncDelegator(opts *bind.TransactOpts, delegator common.Address, stakerID *big.Int) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "_syncDelegator", delegator, stakerID)
}

// SyncDelegator is a paid mutator transaction binding the contract method 0x5dc03f1f.
//
// Solidity: function _syncDelegator(address delegator, uint256 stakerID) returns()
func (_Stakers *StakersSession) SyncDelegator(delegator common.Address, stakerID *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.SyncDelegator(&_Stakers.TransactOpts, delegator, stakerID)
}

// SyncDelegator is a paid mutator transaction binding the contract method 0x5dc03f1f.
//
// Solidity: function _syncDelegator(address delegator, uint256 stakerID) returns()
func (_Stakers *StakersTransactorSession) SyncDelegator(delegator common.Address, stakerID *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.SyncDelegator(&_Stakers.TransactOpts, delegator, stakerID)
}

// SyncStaker is a paid mutator transaction binding the contract method 0xeac3baf2.
//
// Solidity: function _syncStaker(uint256 stakerID) returns()
func (_Stakers *StakersTransactor) SyncStaker(opts *bind.TransactOpts, stakerID *big.Int) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "_syncStaker", stakerID)
}

// SyncStaker is a paid mutator transaction binding the contract method 0xeac3baf2.
//
// Solidity: function _syncStaker(uint256 stakerID) returns()
func (_Stakers *StakersSession) SyncStaker(stakerID *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.SyncStaker(&_Stakers.TransactOpts, stakerID)
}

// SyncStaker is a paid mutator transaction binding the contract method 0xeac3baf2.
//
// Solidity: function _syncStaker(uint256 stakerID) returns()
func (_Stakers *StakersTransactorSession) SyncStaker(stakerID *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.SyncStaker(&_Stakers.TransactOpts, stakerID)
}

// UpgradeStakerStorage is a paid mutator transaction binding the contract method 0x28dca8ff.
//
// Solidity: function _upgradeStakerStorage(uint256 stakerID) returns()
func (_Stakers *StakersTransactor) UpgradeStakerStorage(opts *bind.TransactOpts, stakerID *big.Int) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "_upgradeStakerStorage", stakerID)
}

// UpgradeStakerStorage is a paid mutator transaction binding the contract method 0x28dca8ff.
//
// Solidity: function _upgradeStakerStorage(uint256 stakerID) returns()
func (_Stakers *StakersSession) UpgradeStakerStorage(stakerID *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.UpgradeStakerStorage(&_Stakers.TransactOpts, stakerID)
}

// UpgradeStakerStorage is a paid mutator transaction binding the contract method 0x28dca8ff.
//
// Solidity: function _upgradeStakerStorage(uint256 stakerID) returns()
func (_Stakers *StakersTransactorSession) UpgradeStakerStorage(stakerID *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.UpgradeStakerStorage(&_Stakers.TransactOpts, stakerID)
}

// ClaimDelegationRewards is a paid mutator transaction binding the contract method 0x793c45ce.
//
// Solidity: function claimDelegationRewards(uint256 maxEpochs) returns()
func (_Stakers *StakersTransactor) ClaimDelegationRewards(opts *bind.TransactOpts, maxEpochs *big.Int) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "claimDelegationRewards", maxEpochs)
}

// ClaimDelegationRewards is a paid mutator transaction binding the contract method 0x793c45ce.
//
// Solidity: function claimDelegationRewards(uint256 maxEpochs) returns()
func (_Stakers *StakersSession) ClaimDelegationRewards(maxEpochs *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.ClaimDelegationRewards(&_Stakers.TransactOpts, maxEpochs)
}

// ClaimDelegationRewards is a paid mutator transaction binding the contract method 0x793c45ce.
//
// Solidity: function claimDelegationRewards(uint256 maxEpochs) returns()
func (_Stakers *StakersTransactorSession) ClaimDelegationRewards(maxEpochs *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.ClaimDelegationRewards(&_Stakers.TransactOpts, maxEpochs)
}

// ClaimValidatorRewards is a paid mutator transaction binding the contract method 0x295cccba.
//
// Solidity: function claimValidatorRewards(uint256 maxEpochs) returns()
func (_Stakers *StakersTransactor) ClaimValidatorRewards(opts *bind.TransactOpts, maxEpochs *big.Int) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "claimValidatorRewards", maxEpochs)
}

// ClaimValidatorRewards is a paid mutator transaction binding the contract method 0x295cccba.
//
// Solidity: function claimValidatorRewards(uint256 maxEpochs) returns()
func (_Stakers *StakersSession) ClaimValidatorRewards(maxEpochs *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.ClaimValidatorRewards(&_Stakers.TransactOpts, maxEpochs)
}

// ClaimValidatorRewards is a paid mutator transaction binding the contract method 0x295cccba.
//
// Solidity: function claimValidatorRewards(uint256 maxEpochs) returns()
func (_Stakers *StakersTransactorSession) ClaimValidatorRewards(maxEpochs *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.ClaimValidatorRewards(&_Stakers.TransactOpts, maxEpochs)
}

// CreateDelegation is a paid mutator transaction binding the contract method 0xc312eb07.
//
// Solidity: function createDelegation(uint256 to) payable returns()
func (_Stakers *StakersTransactor) CreateDelegation(opts *bind.TransactOpts, to *big.Int) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "createDelegation", to)
}

// CreateDelegation is a paid mutator transaction binding the contract method 0xc312eb07.
//
// Solidity: function createDelegation(uint256 to) payable returns()
func (_Stakers *StakersSession) CreateDelegation(to *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.CreateDelegation(&_Stakers.TransactOpts, to)
}

// CreateDelegation is a paid mutator transaction binding the contract method 0xc312eb07.
//
// Solidity: function createDelegation(uint256 to) payable returns()
func (_Stakers *StakersTransactorSession) CreateDelegation(to *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.CreateDelegation(&_Stakers.TransactOpts, to)
}

// CreateStake is a paid mutator transaction binding the contract method 0xcc8c2120.
//
// Solidity: function createStake(bytes metadata) payable returns()
func (_Stakers *StakersTransactor) CreateStake(opts *bind.TransactOpts, metadata []byte) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "createStake", metadata)
}

// CreateStake is a paid mutator transaction binding the contract method 0xcc8c2120.
//
// Solidity: function createStake(bytes metadata) payable returns()
func (_Stakers *StakersSession) CreateStake(metadata []byte) (*types.Transaction, error) {
	return _Stakers.Contract.CreateStake(&_Stakers.TransactOpts, metadata)
}

// CreateStake is a paid mutator transaction binding the contract method 0xcc8c2120.
//
// Solidity: function createStake(bytes metadata) payable returns()
func (_Stakers *StakersTransactorSession) CreateStake(metadata []byte) (*types.Transaction, error) {
	return _Stakers.Contract.CreateStake(&_Stakers.TransactOpts, metadata)
}

// CreateStakeWithAddresses is a paid mutator transaction binding the contract method 0x90475ae4.
//
// Solidity: function createStakeWithAddresses(address dagAddress, address sfcAddress, bytes metadata) payable returns()
func (_Stakers *StakersTransactor) CreateStakeWithAddresses(opts *bind.TransactOpts, dagAddress common.Address, sfcAddress common.Address, metadata []byte) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "createStakeWithAddresses", dagAddress, sfcAddress, metadata)
}

// CreateStakeWithAddresses is a paid mutator transaction binding the contract method 0x90475ae4.
//
// Solidity: function createStakeWithAddresses(address dagAddress, address sfcAddress, bytes metadata) payable returns()
func (_Stakers *StakersSession) CreateStakeWithAddresses(dagAddress common.Address, sfcAddress common.Address, metadata []byte) (*types.Transaction, error) {
	return _Stakers.Contract.CreateStakeWithAddresses(&_Stakers.TransactOpts, dagAddress, sfcAddress, metadata)
}

// CreateStakeWithAddresses is a paid mutator transaction binding the contract method 0x90475ae4.
//
// Solidity: function createStakeWithAddresses(address dagAddress, address sfcAddress, bytes metadata) payable returns()
func (_Stakers *StakersTransactorSession) CreateStakeWithAddresses(dagAddress common.Address, sfcAddress common.Address, metadata []byte) (*types.Transaction, error) {
	return _Stakers.Contract.CreateStakeWithAddresses(&_Stakers.TransactOpts, dagAddress, sfcAddress, metadata)
}

// IncreaseDelegation is a paid mutator transaction binding the contract method 0x3a274ff6.
//
// Solidity: function increaseDelegation() payable returns()
func (_Stakers *StakersTransactor) IncreaseDelegation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "increaseDelegation")
}

// IncreaseDelegation is a paid mutator transaction binding the contract method 0x3a274ff6.
//
// Solidity: function increaseDelegation() payable returns()
func (_Stakers *StakersSession) IncreaseDelegation() (*types.Transaction, error) {
	return _Stakers.Contract.IncreaseDelegation(&_Stakers.TransactOpts)
}

// IncreaseDelegation is a paid mutator transaction binding the contract method 0x3a274ff6.
//
// Solidity: function increaseDelegation() payable returns()
func (_Stakers *StakersTransactorSession) IncreaseDelegation() (*types.Transaction, error) {
	return _Stakers.Contract.IncreaseDelegation(&_Stakers.TransactOpts)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0xd9e257ef.
//
// Solidity: function increaseStake() payable returns()
func (_Stakers *StakersTransactor) IncreaseStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "increaseStake")
}

// IncreaseStake is a paid mutator transaction binding the contract method 0xd9e257ef.
//
// Solidity: function increaseStake() payable returns()
func (_Stakers *StakersSession) IncreaseStake() (*types.Transaction, error) {
	return _Stakers.Contract.IncreaseStake(&_Stakers.TransactOpts)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0xd9e257ef.
//
// Solidity: function increaseStake() payable returns()
func (_Stakers *StakersTransactorSession) IncreaseStake() (*types.Transaction, error) {
	return _Stakers.Contract.IncreaseStake(&_Stakers.TransactOpts)
}

// LockUpDelegation is a paid mutator transaction binding the contract method 0xa4b89fab.
//
// Solidity: function lockUpDelegation(uint256 lockDuration, uint256 stakerID) returns()
func (_Stakers *StakersTransactor) LockUpDelegation(opts *bind.TransactOpts, lockDuration *big.Int, stakerID *big.Int) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "lockUpDelegation", lockDuration, stakerID)
}

// LockUpDelegation is a paid mutator transaction binding the contract method 0xa4b89fab.
//
// Solidity: function lockUpDelegation(uint256 lockDuration, uint256 stakerID) returns()
func (_Stakers *StakersSession) LockUpDelegation(lockDuration *big.Int, stakerID *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.LockUpDelegation(&_Stakers.TransactOpts, lockDuration, stakerID)
}

// LockUpDelegation is a paid mutator transaction binding the contract method 0xa4b89fab.
//
// Solidity: function lockUpDelegation(uint256 lockDuration, uint256 stakerID) returns()
func (_Stakers *StakersTransactorSession) LockUpDelegation(lockDuration *big.Int, stakerID *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.LockUpDelegation(&_Stakers.TransactOpts, lockDuration, stakerID)
}

// LockUpStake is a paid mutator transaction binding the contract method 0xf3ae5b1a.
//
// Solidity: function lockUpStake(uint256 lockDuration) returns()
func (_Stakers *StakersTransactor) LockUpStake(opts *bind.TransactOpts, lockDuration *big.Int) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "lockUpStake", lockDuration)
}

// LockUpStake is a paid mutator transaction binding the contract method 0xf3ae5b1a.
//
// Solidity: function lockUpStake(uint256 lockDuration) returns()
func (_Stakers *StakersSession) LockUpStake(lockDuration *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.LockUpStake(&_Stakers.TransactOpts, lockDuration)
}

// LockUpStake is a paid mutator transaction binding the contract method 0xf3ae5b1a.
//
// Solidity: function lockUpStake(uint256 lockDuration) returns()
func (_Stakers *StakersTransactorSession) LockUpStake(lockDuration *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.LockUpStake(&_Stakers.TransactOpts, lockDuration)
}

// PartialWithdrawByRequest is a paid mutator transaction binding the contract method 0xf8b18d8a.
//
// Solidity: function partialWithdrawByRequest(uint256 wrID) returns()
func (_Stakers *StakersTransactor) PartialWithdrawByRequest(opts *bind.TransactOpts, wrID *big.Int) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "partialWithdrawByRequest", wrID)
}

// PartialWithdrawByRequest is a paid mutator transaction binding the contract method 0xf8b18d8a.
//
// Solidity: function partialWithdrawByRequest(uint256 wrID) returns()
func (_Stakers *StakersSession) PartialWithdrawByRequest(wrID *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.PartialWithdrawByRequest(&_Stakers.TransactOpts, wrID)
}

// PartialWithdrawByRequest is a paid mutator transaction binding the contract method 0xf8b18d8a.
//
// Solidity: function partialWithdrawByRequest(uint256 wrID) returns()
func (_Stakers *StakersTransactorSession) PartialWithdrawByRequest(wrID *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.PartialWithdrawByRequest(&_Stakers.TransactOpts, wrID)
}

// PrepareToWithdrawDelegation is a paid mutator transaction binding the contract method 0x1c333318.
//
// Solidity: function prepareToWithdrawDelegation() returns()
func (_Stakers *StakersTransactor) PrepareToWithdrawDelegation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "prepareToWithdrawDelegation")
}

// PrepareToWithdrawDelegation is a paid mutator transaction binding the contract method 0x1c333318.
//
// Solidity: function prepareToWithdrawDelegation() returns()
func (_Stakers *StakersSession) PrepareToWithdrawDelegation() (*types.Transaction, error) {
	return _Stakers.Contract.PrepareToWithdrawDelegation(&_Stakers.TransactOpts)
}

// PrepareToWithdrawDelegation is a paid mutator transaction binding the contract method 0x1c333318.
//
// Solidity: function prepareToWithdrawDelegation() returns()
func (_Stakers *StakersTransactorSession) PrepareToWithdrawDelegation() (*types.Transaction, error) {
	return _Stakers.Contract.PrepareToWithdrawDelegation(&_Stakers.TransactOpts)
}

// PrepareToWithdrawDelegationPartial is a paid mutator transaction binding the contract method 0xe7ff9e78.
//
// Solidity: function prepareToWithdrawDelegationPartial(uint256 wrID, uint256 amount) returns()
func (_Stakers *StakersTransactor) PrepareToWithdrawDelegationPartial(opts *bind.TransactOpts, wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "prepareToWithdrawDelegationPartial", wrID, amount)
}

// PrepareToWithdrawDelegationPartial is a paid mutator transaction binding the contract method 0xe7ff9e78.
//
// Solidity: function prepareToWithdrawDelegationPartial(uint256 wrID, uint256 amount) returns()
func (_Stakers *StakersSession) PrepareToWithdrawDelegationPartial(wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.PrepareToWithdrawDelegationPartial(&_Stakers.TransactOpts, wrID, amount)
}

// PrepareToWithdrawDelegationPartial is a paid mutator transaction binding the contract method 0xe7ff9e78.
//
// Solidity: function prepareToWithdrawDelegationPartial(uint256 wrID, uint256 amount) returns()
func (_Stakers *StakersTransactorSession) PrepareToWithdrawDelegationPartial(wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.PrepareToWithdrawDelegationPartial(&_Stakers.TransactOpts, wrID, amount)
}

// PrepareToWithdrawStake is a paid mutator transaction binding the contract method 0xc41b6405.
//
// Solidity: function prepareToWithdrawStake() returns()
func (_Stakers *StakersTransactor) PrepareToWithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "prepareToWithdrawStake")
}

// PrepareToWithdrawStake is a paid mutator transaction binding the contract method 0xc41b6405.
//
// Solidity: function prepareToWithdrawStake() returns()
func (_Stakers *StakersSession) PrepareToWithdrawStake() (*types.Transaction, error) {
	return _Stakers.Contract.PrepareToWithdrawStake(&_Stakers.TransactOpts)
}

// PrepareToWithdrawStake is a paid mutator transaction binding the contract method 0xc41b6405.
//
// Solidity: function prepareToWithdrawStake() returns()
func (_Stakers *StakersTransactorSession) PrepareToWithdrawStake() (*types.Transaction, error) {
	return _Stakers.Contract.PrepareToWithdrawStake(&_Stakers.TransactOpts)
}

// PrepareToWithdrawStakePartial is a paid mutator transaction binding the contract method 0x26682c71.
//
// Solidity: function prepareToWithdrawStakePartial(uint256 wrID, uint256 amount) returns()
func (_Stakers *StakersTransactor) PrepareToWithdrawStakePartial(opts *bind.TransactOpts, wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "prepareToWithdrawStakePartial", wrID, amount)
}

// PrepareToWithdrawStakePartial is a paid mutator transaction binding the contract method 0x26682c71.
//
// Solidity: function prepareToWithdrawStakePartial(uint256 wrID, uint256 amount) returns()
func (_Stakers *StakersSession) PrepareToWithdrawStakePartial(wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.PrepareToWithdrawStakePartial(&_Stakers.TransactOpts, wrID, amount)
}

// PrepareToWithdrawStakePartial is a paid mutator transaction binding the contract method 0x26682c71.
//
// Solidity: function prepareToWithdrawStakePartial(uint256 wrID, uint256 amount) returns()
func (_Stakers *StakersTransactorSession) PrepareToWithdrawStakePartial(wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.PrepareToWithdrawStakePartial(&_Stakers.TransactOpts, wrID, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stakers *StakersTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stakers *StakersSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stakers.Contract.RenounceOwnership(&_Stakers.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stakers *StakersTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stakers.Contract.RenounceOwnership(&_Stakers.TransactOpts)
}

// StartLockedUp is a paid mutator transaction binding the contract method 0xc9400d4f.
//
// Solidity: function startLockedUp(uint256 epochNum) returns()
func (_Stakers *StakersTransactor) StartLockedUp(opts *bind.TransactOpts, epochNum *big.Int) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "startLockedUp", epochNum)
}

// StartLockedUp is a paid mutator transaction binding the contract method 0xc9400d4f.
//
// Solidity: function startLockedUp(uint256 epochNum) returns()
func (_Stakers *StakersSession) StartLockedUp(epochNum *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.StartLockedUp(&_Stakers.TransactOpts, epochNum)
}

// StartLockedUp is a paid mutator transaction binding the contract method 0xc9400d4f.
//
// Solidity: function startLockedUp(uint256 epochNum) returns()
func (_Stakers *StakersTransactorSession) StartLockedUp(epochNum *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.StartLockedUp(&_Stakers.TransactOpts, epochNum)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stakers *StakersTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stakers *StakersSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stakers.Contract.TransferOwnership(&_Stakers.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stakers *StakersTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stakers.Contract.TransferOwnership(&_Stakers.TransactOpts, newOwner)
}

// UnstashRewards is a paid mutator transaction binding the contract method 0x876f7e2a.
//
// Solidity: function unstashRewards() returns()
func (_Stakers *StakersTransactor) UnstashRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "unstashRewards")
}

// UnstashRewards is a paid mutator transaction binding the contract method 0x876f7e2a.
//
// Solidity: function unstashRewards() returns()
func (_Stakers *StakersSession) UnstashRewards() (*types.Transaction, error) {
	return _Stakers.Contract.UnstashRewards(&_Stakers.TransactOpts)
}

// UnstashRewards is a paid mutator transaction binding the contract method 0x876f7e2a.
//
// Solidity: function unstashRewards() returns()
func (_Stakers *StakersTransactorSession) UnstashRewards() (*types.Transaction, error) {
	return _Stakers.Contract.UnstashRewards(&_Stakers.TransactOpts)
}

// UpdateBaseRewardPerSec is a paid mutator transaction binding the contract method 0x1b593d8a.
//
// Solidity: function updateBaseRewardPerSec(uint256 value) returns()
func (_Stakers *StakersTransactor) UpdateBaseRewardPerSec(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "updateBaseRewardPerSec", value)
}

// UpdateBaseRewardPerSec is a paid mutator transaction binding the contract method 0x1b593d8a.
//
// Solidity: function updateBaseRewardPerSec(uint256 value) returns()
func (_Stakers *StakersSession) UpdateBaseRewardPerSec(value *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.UpdateBaseRewardPerSec(&_Stakers.TransactOpts, value)
}

// UpdateBaseRewardPerSec is a paid mutator transaction binding the contract method 0x1b593d8a.
//
// Solidity: function updateBaseRewardPerSec(uint256 value) returns()
func (_Stakers *StakersTransactorSession) UpdateBaseRewardPerSec(value *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.UpdateBaseRewardPerSec(&_Stakers.TransactOpts, value)
}

// UpdateGasPowerAllocationRate is a paid mutator transaction binding the contract method 0x119e351a.
//
// Solidity: function updateGasPowerAllocationRate(uint256 short, uint256 long) returns()
func (_Stakers *StakersTransactor) UpdateGasPowerAllocationRate(opts *bind.TransactOpts, short *big.Int, long *big.Int) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "updateGasPowerAllocationRate", short, long)
}

// UpdateGasPowerAllocationRate is a paid mutator transaction binding the contract method 0x119e351a.
//
// Solidity: function updateGasPowerAllocationRate(uint256 short, uint256 long) returns()
func (_Stakers *StakersSession) UpdateGasPowerAllocationRate(short *big.Int, long *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.UpdateGasPowerAllocationRate(&_Stakers.TransactOpts, short, long)
}

// UpdateGasPowerAllocationRate is a paid mutator transaction binding the contract method 0x119e351a.
//
// Solidity: function updateGasPowerAllocationRate(uint256 short, uint256 long) returns()
func (_Stakers *StakersTransactorSession) UpdateGasPowerAllocationRate(short *big.Int, long *big.Int) (*types.Transaction, error) {
	return _Stakers.Contract.UpdateGasPowerAllocationRate(&_Stakers.TransactOpts, short, long)
}

// UpdateStakerMetadata is a paid mutator transaction binding the contract method 0x33a14912.
//
// Solidity: function updateStakerMetadata(bytes metadata) returns()
func (_Stakers *StakersTransactor) UpdateStakerMetadata(opts *bind.TransactOpts, metadata []byte) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "updateStakerMetadata", metadata)
}

// UpdateStakerMetadata is a paid mutator transaction binding the contract method 0x33a14912.
//
// Solidity: function updateStakerMetadata(bytes metadata) returns()
func (_Stakers *StakersSession) UpdateStakerMetadata(metadata []byte) (*types.Transaction, error) {
	return _Stakers.Contract.UpdateStakerMetadata(&_Stakers.TransactOpts, metadata)
}

// UpdateStakerMetadata is a paid mutator transaction binding the contract method 0x33a14912.
//
// Solidity: function updateStakerMetadata(bytes metadata) returns()
func (_Stakers *StakersTransactorSession) UpdateStakerMetadata(metadata []byte) (*types.Transaction, error) {
	return _Stakers.Contract.UpdateStakerMetadata(&_Stakers.TransactOpts, metadata)
}

// UpdateStakerSfcAddress is a paid mutator transaction binding the contract method 0xc3d74f1a.
//
// Solidity: function updateStakerSfcAddress(address newSfcAddress) returns()
func (_Stakers *StakersTransactor) UpdateStakerSfcAddress(opts *bind.TransactOpts, newSfcAddress common.Address) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "updateStakerSfcAddress", newSfcAddress)
}

// UpdateStakerSfcAddress is a paid mutator transaction binding the contract method 0xc3d74f1a.
//
// Solidity: function updateStakerSfcAddress(address newSfcAddress) returns()
func (_Stakers *StakersSession) UpdateStakerSfcAddress(newSfcAddress common.Address) (*types.Transaction, error) {
	return _Stakers.Contract.UpdateStakerSfcAddress(&_Stakers.TransactOpts, newSfcAddress)
}

// UpdateStakerSfcAddress is a paid mutator transaction binding the contract method 0xc3d74f1a.
//
// Solidity: function updateStakerSfcAddress(address newSfcAddress) returns()
func (_Stakers *StakersTransactorSession) UpdateStakerSfcAddress(newSfcAddress common.Address) (*types.Transaction, error) {
	return _Stakers.Contract.UpdateStakerSfcAddress(&_Stakers.TransactOpts, newSfcAddress)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0x16bfdd81.
//
// Solidity: function withdrawDelegation() returns()
func (_Stakers *StakersTransactor) WithdrawDelegation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "withdrawDelegation")
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0x16bfdd81.
//
// Solidity: function withdrawDelegation() returns()
func (_Stakers *StakersSession) WithdrawDelegation() (*types.Transaction, error) {
	return _Stakers.Contract.WithdrawDelegation(&_Stakers.TransactOpts)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0x16bfdd81.
//
// Solidity: function withdrawDelegation() returns()
func (_Stakers *StakersTransactorSession) WithdrawDelegation() (*types.Transaction, error) {
	return _Stakers.Contract.WithdrawDelegation(&_Stakers.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Stakers *StakersTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakers.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Stakers *StakersSession) WithdrawStake() (*types.Transaction, error) {
	return _Stakers.Contract.WithdrawStake(&_Stakers.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Stakers *StakersTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _Stakers.Contract.WithdrawStake(&_Stakers.TransactOpts)
}

// StakersBurntRewardStashIterator is returned from FilterBurntRewardStash and is used to iterate over the raw logs and unpacked data for BurntRewardStash events raised by the Stakers contract.
type StakersBurntRewardStashIterator struct {
	Event *StakersBurntRewardStash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersBurntRewardStashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersBurntRewardStash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersBurntRewardStash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersBurntRewardStashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersBurntRewardStashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersBurntRewardStash represents a BurntRewardStash event raised by the Stakers contract.
type StakersBurntRewardStash struct {
	Addr         common.Address
	StakerID     *big.Int
	IsDelegation bool
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBurntRewardStash is a free log retrieval operation binding the contract event 0x0ea92567e76d40ddc52d2c1d74a521a59329a38b50411451de6ad2e565466d0f.
//
// Solidity: event BurntRewardStash(address indexed addr, uint256 indexed stakerID, bool isDelegation, uint256 amount)
func (_Stakers *StakersFilterer) FilterBurntRewardStash(opts *bind.FilterOpts, addr []common.Address, stakerID []*big.Int) (*StakersBurntRewardStashIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "BurntRewardStash", addrRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &StakersBurntRewardStashIterator{contract: _Stakers.contract, event: "BurntRewardStash", logs: logs, sub: sub}, nil
}

// WatchBurntRewardStash is a free log subscription operation binding the contract event 0x0ea92567e76d40ddc52d2c1d74a521a59329a38b50411451de6ad2e565466d0f.
//
// Solidity: event BurntRewardStash(address indexed addr, uint256 indexed stakerID, bool isDelegation, uint256 amount)
func (_Stakers *StakersFilterer) WatchBurntRewardStash(opts *bind.WatchOpts, sink chan<- *StakersBurntRewardStash, addr []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "BurntRewardStash", addrRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersBurntRewardStash)
				if err := _Stakers.contract.UnpackLog(event, "BurntRewardStash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurntRewardStash is a log parse operation binding the contract event 0x0ea92567e76d40ddc52d2c1d74a521a59329a38b50411451de6ad2e565466d0f.
//
// Solidity: event BurntRewardStash(address indexed addr, uint256 indexed stakerID, bool isDelegation, uint256 amount)
func (_Stakers *StakersFilterer) ParseBurntRewardStash(log types.Log) (*StakersBurntRewardStash, error) {
	event := new(StakersBurntRewardStash)
	if err := _Stakers.contract.UnpackLog(event, "BurntRewardStash", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersClaimedDelegationRewardIterator is returned from FilterClaimedDelegationReward and is used to iterate over the raw logs and unpacked data for ClaimedDelegationReward events raised by the Stakers contract.
type StakersClaimedDelegationRewardIterator struct {
	Event *StakersClaimedDelegationReward // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersClaimedDelegationRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersClaimedDelegationReward)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersClaimedDelegationReward)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersClaimedDelegationRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersClaimedDelegationRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersClaimedDelegationReward represents a ClaimedDelegationReward event raised by the Stakers contract.
type StakersClaimedDelegationReward struct {
	From       common.Address
	StakerID   *big.Int
	Reward     *big.Int
	FromEpoch  *big.Int
	UntilEpoch *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimedDelegationReward is a free log retrieval operation binding the contract event 0x2676e1697cf4731b93ddb4ef54e0e5a98c06cccbbbb2202848a3c6286595e6ce.
//
// Solidity: event ClaimedDelegationReward(address indexed from, uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_Stakers *StakersFilterer) FilterClaimedDelegationReward(opts *bind.FilterOpts, from []common.Address, stakerID []*big.Int) (*StakersClaimedDelegationRewardIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "ClaimedDelegationReward", fromRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &StakersClaimedDelegationRewardIterator{contract: _Stakers.contract, event: "ClaimedDelegationReward", logs: logs, sub: sub}, nil
}

// WatchClaimedDelegationReward is a free log subscription operation binding the contract event 0x2676e1697cf4731b93ddb4ef54e0e5a98c06cccbbbb2202848a3c6286595e6ce.
//
// Solidity: event ClaimedDelegationReward(address indexed from, uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_Stakers *StakersFilterer) WatchClaimedDelegationReward(opts *bind.WatchOpts, sink chan<- *StakersClaimedDelegationReward, from []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "ClaimedDelegationReward", fromRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersClaimedDelegationReward)
				if err := _Stakers.contract.UnpackLog(event, "ClaimedDelegationReward", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimedDelegationReward is a log parse operation binding the contract event 0x2676e1697cf4731b93ddb4ef54e0e5a98c06cccbbbb2202848a3c6286595e6ce.
//
// Solidity: event ClaimedDelegationReward(address indexed from, uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_Stakers *StakersFilterer) ParseClaimedDelegationReward(log types.Log) (*StakersClaimedDelegationReward, error) {
	event := new(StakersClaimedDelegationReward)
	if err := _Stakers.contract.UnpackLog(event, "ClaimedDelegationReward", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersClaimedValidatorRewardIterator is returned from FilterClaimedValidatorReward and is used to iterate over the raw logs and unpacked data for ClaimedValidatorReward events raised by the Stakers contract.
type StakersClaimedValidatorRewardIterator struct {
	Event *StakersClaimedValidatorReward // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersClaimedValidatorRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersClaimedValidatorReward)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersClaimedValidatorReward)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersClaimedValidatorRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersClaimedValidatorRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersClaimedValidatorReward represents a ClaimedValidatorReward event raised by the Stakers contract.
type StakersClaimedValidatorReward struct {
	StakerID   *big.Int
	Reward     *big.Int
	FromEpoch  *big.Int
	UntilEpoch *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimedValidatorReward is a free log retrieval operation binding the contract event 0x2ea54c2b22a07549d19fb5eb8e4e48ebe1c653117215e94d5468c5612750d35c.
//
// Solidity: event ClaimedValidatorReward(uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_Stakers *StakersFilterer) FilterClaimedValidatorReward(opts *bind.FilterOpts, stakerID []*big.Int) (*StakersClaimedValidatorRewardIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "ClaimedValidatorReward", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &StakersClaimedValidatorRewardIterator{contract: _Stakers.contract, event: "ClaimedValidatorReward", logs: logs, sub: sub}, nil
}

// WatchClaimedValidatorReward is a free log subscription operation binding the contract event 0x2ea54c2b22a07549d19fb5eb8e4e48ebe1c653117215e94d5468c5612750d35c.
//
// Solidity: event ClaimedValidatorReward(uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_Stakers *StakersFilterer) WatchClaimedValidatorReward(opts *bind.WatchOpts, sink chan<- *StakersClaimedValidatorReward, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "ClaimedValidatorReward", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersClaimedValidatorReward)
				if err := _Stakers.contract.UnpackLog(event, "ClaimedValidatorReward", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimedValidatorReward is a log parse operation binding the contract event 0x2ea54c2b22a07549d19fb5eb8e4e48ebe1c653117215e94d5468c5612750d35c.
//
// Solidity: event ClaimedValidatorReward(uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_Stakers *StakersFilterer) ParseClaimedValidatorReward(log types.Log) (*StakersClaimedValidatorReward, error) {
	event := new(StakersClaimedValidatorReward)
	if err := _Stakers.contract.UnpackLog(event, "ClaimedValidatorReward", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersCreatedDelegationIterator is returned from FilterCreatedDelegation and is used to iterate over the raw logs and unpacked data for CreatedDelegation events raised by the Stakers contract.
type StakersCreatedDelegationIterator struct {
	Event *StakersCreatedDelegation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersCreatedDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersCreatedDelegation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersCreatedDelegation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersCreatedDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersCreatedDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersCreatedDelegation represents a CreatedDelegation event raised by the Stakers contract.
type StakersCreatedDelegation struct {
	Delegator  common.Address
	ToStakerID *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreatedDelegation is a free log retrieval operation binding the contract event 0xfd8c857fb9acd6f4ad59b8621a2a77825168b7b4b76de9586d08e00d4ed462be.
//
// Solidity: event CreatedDelegation(address indexed delegator, uint256 indexed toStakerID, uint256 amount)
func (_Stakers *StakersFilterer) FilterCreatedDelegation(opts *bind.FilterOpts, delegator []common.Address, toStakerID []*big.Int) (*StakersCreatedDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toStakerIDRule []interface{}
	for _, toStakerIDItem := range toStakerID {
		toStakerIDRule = append(toStakerIDRule, toStakerIDItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "CreatedDelegation", delegatorRule, toStakerIDRule)
	if err != nil {
		return nil, err
	}
	return &StakersCreatedDelegationIterator{contract: _Stakers.contract, event: "CreatedDelegation", logs: logs, sub: sub}, nil
}

// WatchCreatedDelegation is a free log subscription operation binding the contract event 0xfd8c857fb9acd6f4ad59b8621a2a77825168b7b4b76de9586d08e00d4ed462be.
//
// Solidity: event CreatedDelegation(address indexed delegator, uint256 indexed toStakerID, uint256 amount)
func (_Stakers *StakersFilterer) WatchCreatedDelegation(opts *bind.WatchOpts, sink chan<- *StakersCreatedDelegation, delegator []common.Address, toStakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toStakerIDRule []interface{}
	for _, toStakerIDItem := range toStakerID {
		toStakerIDRule = append(toStakerIDRule, toStakerIDItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "CreatedDelegation", delegatorRule, toStakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersCreatedDelegation)
				if err := _Stakers.contract.UnpackLog(event, "CreatedDelegation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreatedDelegation is a log parse operation binding the contract event 0xfd8c857fb9acd6f4ad59b8621a2a77825168b7b4b76de9586d08e00d4ed462be.
//
// Solidity: event CreatedDelegation(address indexed delegator, uint256 indexed toStakerID, uint256 amount)
func (_Stakers *StakersFilterer) ParseCreatedDelegation(log types.Log) (*StakersCreatedDelegation, error) {
	event := new(StakersCreatedDelegation)
	if err := _Stakers.contract.UnpackLog(event, "CreatedDelegation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersCreatedStakeIterator is returned from FilterCreatedStake and is used to iterate over the raw logs and unpacked data for CreatedStake events raised by the Stakers contract.
type StakersCreatedStakeIterator struct {
	Event *StakersCreatedStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersCreatedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersCreatedStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersCreatedStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersCreatedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersCreatedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersCreatedStake represents a CreatedStake event raised by the Stakers contract.
type StakersCreatedStake struct {
	StakerID      *big.Int
	DagSfcAddress common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCreatedStake is a free log retrieval operation binding the contract event 0x0697dfe5062b9db8108e4b31254f47a912ae6bbb78837667b2e923a6f5160d39.
//
// Solidity: event CreatedStake(uint256 indexed stakerID, address indexed dagSfcAddress, uint256 amount)
func (_Stakers *StakersFilterer) FilterCreatedStake(opts *bind.FilterOpts, stakerID []*big.Int, dagSfcAddress []common.Address) (*StakersCreatedStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}
	var dagSfcAddressRule []interface{}
	for _, dagSfcAddressItem := range dagSfcAddress {
		dagSfcAddressRule = append(dagSfcAddressRule, dagSfcAddressItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "CreatedStake", stakerIDRule, dagSfcAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakersCreatedStakeIterator{contract: _Stakers.contract, event: "CreatedStake", logs: logs, sub: sub}, nil
}

// WatchCreatedStake is a free log subscription operation binding the contract event 0x0697dfe5062b9db8108e4b31254f47a912ae6bbb78837667b2e923a6f5160d39.
//
// Solidity: event CreatedStake(uint256 indexed stakerID, address indexed dagSfcAddress, uint256 amount)
func (_Stakers *StakersFilterer) WatchCreatedStake(opts *bind.WatchOpts, sink chan<- *StakersCreatedStake, stakerID []*big.Int, dagSfcAddress []common.Address) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}
	var dagSfcAddressRule []interface{}
	for _, dagSfcAddressItem := range dagSfcAddress {
		dagSfcAddressRule = append(dagSfcAddressRule, dagSfcAddressItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "CreatedStake", stakerIDRule, dagSfcAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersCreatedStake)
				if err := _Stakers.contract.UnpackLog(event, "CreatedStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreatedStake is a log parse operation binding the contract event 0x0697dfe5062b9db8108e4b31254f47a912ae6bbb78837667b2e923a6f5160d39.
//
// Solidity: event CreatedStake(uint256 indexed stakerID, address indexed dagSfcAddress, uint256 amount)
func (_Stakers *StakersFilterer) ParseCreatedStake(log types.Log) (*StakersCreatedStake, error) {
	event := new(StakersCreatedStake)
	if err := _Stakers.contract.UnpackLog(event, "CreatedStake", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersCreatedWithdrawRequestIterator is returned from FilterCreatedWithdrawRequest and is used to iterate over the raw logs and unpacked data for CreatedWithdrawRequest events raised by the Stakers contract.
type StakersCreatedWithdrawRequestIterator struct {
	Event *StakersCreatedWithdrawRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersCreatedWithdrawRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersCreatedWithdrawRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersCreatedWithdrawRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersCreatedWithdrawRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersCreatedWithdrawRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersCreatedWithdrawRequest represents a CreatedWithdrawRequest event raised by the Stakers contract.
type StakersCreatedWithdrawRequest struct {
	Auth       common.Address
	Receiver   common.Address
	StakerID   *big.Int
	WrID       *big.Int
	Delegation bool
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreatedWithdrawRequest is a free log retrieval operation binding the contract event 0xde2d2a87af2fa2de55bde86f04143144eb632fa6be266dc224341a371fb8916d.
//
// Solidity: event CreatedWithdrawRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 amount)
func (_Stakers *StakersFilterer) FilterCreatedWithdrawRequest(opts *bind.FilterOpts, auth []common.Address, receiver []common.Address, stakerID []*big.Int) (*StakersCreatedWithdrawRequestIterator, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "CreatedWithdrawRequest", authRule, receiverRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &StakersCreatedWithdrawRequestIterator{contract: _Stakers.contract, event: "CreatedWithdrawRequest", logs: logs, sub: sub}, nil
}

// WatchCreatedWithdrawRequest is a free log subscription operation binding the contract event 0xde2d2a87af2fa2de55bde86f04143144eb632fa6be266dc224341a371fb8916d.
//
// Solidity: event CreatedWithdrawRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 amount)
func (_Stakers *StakersFilterer) WatchCreatedWithdrawRequest(opts *bind.WatchOpts, sink chan<- *StakersCreatedWithdrawRequest, auth []common.Address, receiver []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "CreatedWithdrawRequest", authRule, receiverRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersCreatedWithdrawRequest)
				if err := _Stakers.contract.UnpackLog(event, "CreatedWithdrawRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreatedWithdrawRequest is a log parse operation binding the contract event 0xde2d2a87af2fa2de55bde86f04143144eb632fa6be266dc224341a371fb8916d.
//
// Solidity: event CreatedWithdrawRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 amount)
func (_Stakers *StakersFilterer) ParseCreatedWithdrawRequest(log types.Log) (*StakersCreatedWithdrawRequest, error) {
	event := new(StakersCreatedWithdrawRequest)
	if err := _Stakers.contract.UnpackLog(event, "CreatedWithdrawRequest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersDeactivatedDelegationIterator is returned from FilterDeactivatedDelegation and is used to iterate over the raw logs and unpacked data for DeactivatedDelegation events raised by the Stakers contract.
type StakersDeactivatedDelegationIterator struct {
	Event *StakersDeactivatedDelegation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersDeactivatedDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersDeactivatedDelegation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersDeactivatedDelegation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersDeactivatedDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersDeactivatedDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersDeactivatedDelegation represents a DeactivatedDelegation event raised by the Stakers contract.
type StakersDeactivatedDelegation struct {
	Delegator common.Address
	StakerID  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeactivatedDelegation is a free log retrieval operation binding the contract event 0x912c4125a208704a342cbdc4726795d26556b0170b7fc95bc706d5cb1f506469.
//
// Solidity: event DeactivatedDelegation(address indexed delegator, uint256 indexed stakerID)
func (_Stakers *StakersFilterer) FilterDeactivatedDelegation(opts *bind.FilterOpts, delegator []common.Address, stakerID []*big.Int) (*StakersDeactivatedDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "DeactivatedDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &StakersDeactivatedDelegationIterator{contract: _Stakers.contract, event: "DeactivatedDelegation", logs: logs, sub: sub}, nil
}

// WatchDeactivatedDelegation is a free log subscription operation binding the contract event 0x912c4125a208704a342cbdc4726795d26556b0170b7fc95bc706d5cb1f506469.
//
// Solidity: event DeactivatedDelegation(address indexed delegator, uint256 indexed stakerID)
func (_Stakers *StakersFilterer) WatchDeactivatedDelegation(opts *bind.WatchOpts, sink chan<- *StakersDeactivatedDelegation, delegator []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "DeactivatedDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersDeactivatedDelegation)
				if err := _Stakers.contract.UnpackLog(event, "DeactivatedDelegation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeactivatedDelegation is a log parse operation binding the contract event 0x912c4125a208704a342cbdc4726795d26556b0170b7fc95bc706d5cb1f506469.
//
// Solidity: event DeactivatedDelegation(address indexed delegator, uint256 indexed stakerID)
func (_Stakers *StakersFilterer) ParseDeactivatedDelegation(log types.Log) (*StakersDeactivatedDelegation, error) {
	event := new(StakersDeactivatedDelegation)
	if err := _Stakers.contract.UnpackLog(event, "DeactivatedDelegation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersDeactivatedStakeIterator is returned from FilterDeactivatedStake and is used to iterate over the raw logs and unpacked data for DeactivatedStake events raised by the Stakers contract.
type StakersDeactivatedStakeIterator struct {
	Event *StakersDeactivatedStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersDeactivatedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersDeactivatedStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersDeactivatedStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersDeactivatedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersDeactivatedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersDeactivatedStake represents a DeactivatedStake event raised by the Stakers contract.
type StakersDeactivatedStake struct {
	StakerID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeactivatedStake is a free log retrieval operation binding the contract event 0xf7c308d0d978cce3aec157d1b34e355db4636b4e71ce91b4f5ec9e7a4f5cdc60.
//
// Solidity: event DeactivatedStake(uint256 indexed stakerID)
func (_Stakers *StakersFilterer) FilterDeactivatedStake(opts *bind.FilterOpts, stakerID []*big.Int) (*StakersDeactivatedStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "DeactivatedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &StakersDeactivatedStakeIterator{contract: _Stakers.contract, event: "DeactivatedStake", logs: logs, sub: sub}, nil
}

// WatchDeactivatedStake is a free log subscription operation binding the contract event 0xf7c308d0d978cce3aec157d1b34e355db4636b4e71ce91b4f5ec9e7a4f5cdc60.
//
// Solidity: event DeactivatedStake(uint256 indexed stakerID)
func (_Stakers *StakersFilterer) WatchDeactivatedStake(opts *bind.WatchOpts, sink chan<- *StakersDeactivatedStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "DeactivatedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersDeactivatedStake)
				if err := _Stakers.contract.UnpackLog(event, "DeactivatedStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeactivatedStake is a log parse operation binding the contract event 0xf7c308d0d978cce3aec157d1b34e355db4636b4e71ce91b4f5ec9e7a4f5cdc60.
//
// Solidity: event DeactivatedStake(uint256 indexed stakerID)
func (_Stakers *StakersFilterer) ParseDeactivatedStake(log types.Log) (*StakersDeactivatedStake, error) {
	event := new(StakersDeactivatedStake)
	if err := _Stakers.contract.UnpackLog(event, "DeactivatedStake", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersIncreasedDelegationIterator is returned from FilterIncreasedDelegation and is used to iterate over the raw logs and unpacked data for IncreasedDelegation events raised by the Stakers contract.
type StakersIncreasedDelegationIterator struct {
	Event *StakersIncreasedDelegation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersIncreasedDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersIncreasedDelegation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersIncreasedDelegation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersIncreasedDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersIncreasedDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersIncreasedDelegation represents a IncreasedDelegation event raised by the Stakers contract.
type StakersIncreasedDelegation struct {
	Delegator common.Address
	StakerID  *big.Int
	NewAmount *big.Int
	Diff      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncreasedDelegation is a free log retrieval operation binding the contract event 0x4ca781bfe171e588a2661d5a7f2f5f59df879c53489063552fbad2145b707fc1.
//
// Solidity: event IncreasedDelegation(address indexed delegator, uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_Stakers *StakersFilterer) FilterIncreasedDelegation(opts *bind.FilterOpts, delegator []common.Address, stakerID []*big.Int) (*StakersIncreasedDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "IncreasedDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &StakersIncreasedDelegationIterator{contract: _Stakers.contract, event: "IncreasedDelegation", logs: logs, sub: sub}, nil
}

// WatchIncreasedDelegation is a free log subscription operation binding the contract event 0x4ca781bfe171e588a2661d5a7f2f5f59df879c53489063552fbad2145b707fc1.
//
// Solidity: event IncreasedDelegation(address indexed delegator, uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_Stakers *StakersFilterer) WatchIncreasedDelegation(opts *bind.WatchOpts, sink chan<- *StakersIncreasedDelegation, delegator []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "IncreasedDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersIncreasedDelegation)
				if err := _Stakers.contract.UnpackLog(event, "IncreasedDelegation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncreasedDelegation is a log parse operation binding the contract event 0x4ca781bfe171e588a2661d5a7f2f5f59df879c53489063552fbad2145b707fc1.
//
// Solidity: event IncreasedDelegation(address indexed delegator, uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_Stakers *StakersFilterer) ParseIncreasedDelegation(log types.Log) (*StakersIncreasedDelegation, error) {
	event := new(StakersIncreasedDelegation)
	if err := _Stakers.contract.UnpackLog(event, "IncreasedDelegation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersIncreasedStakeIterator is returned from FilterIncreasedStake and is used to iterate over the raw logs and unpacked data for IncreasedStake events raised by the Stakers contract.
type StakersIncreasedStakeIterator struct {
	Event *StakersIncreasedStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersIncreasedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersIncreasedStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersIncreasedStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersIncreasedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersIncreasedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersIncreasedStake represents a IncreasedStake event raised by the Stakers contract.
type StakersIncreasedStake struct {
	StakerID  *big.Int
	NewAmount *big.Int
	Diff      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncreasedStake is a free log retrieval operation binding the contract event 0xa1d93e9a2a16bf4c2d0cdc6f47fe0fa054c741c96b3dac1297c79eaca31714e9.
//
// Solidity: event IncreasedStake(uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_Stakers *StakersFilterer) FilterIncreasedStake(opts *bind.FilterOpts, stakerID []*big.Int) (*StakersIncreasedStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "IncreasedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &StakersIncreasedStakeIterator{contract: _Stakers.contract, event: "IncreasedStake", logs: logs, sub: sub}, nil
}

// WatchIncreasedStake is a free log subscription operation binding the contract event 0xa1d93e9a2a16bf4c2d0cdc6f47fe0fa054c741c96b3dac1297c79eaca31714e9.
//
// Solidity: event IncreasedStake(uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_Stakers *StakersFilterer) WatchIncreasedStake(opts *bind.WatchOpts, sink chan<- *StakersIncreasedStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "IncreasedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersIncreasedStake)
				if err := _Stakers.contract.UnpackLog(event, "IncreasedStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncreasedStake is a log parse operation binding the contract event 0xa1d93e9a2a16bf4c2d0cdc6f47fe0fa054c741c96b3dac1297c79eaca31714e9.
//
// Solidity: event IncreasedStake(uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_Stakers *StakersFilterer) ParseIncreasedStake(log types.Log) (*StakersIncreasedStake, error) {
	event := new(StakersIncreasedStake)
	if err := _Stakers.contract.UnpackLog(event, "IncreasedStake", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersLockingDelegationIterator is returned from FilterLockingDelegation and is used to iterate over the raw logs and unpacked data for LockingDelegation events raised by the Stakers contract.
type StakersLockingDelegationIterator struct {
	Event *StakersLockingDelegation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersLockingDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersLockingDelegation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersLockingDelegation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersLockingDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersLockingDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersLockingDelegation represents a LockingDelegation event raised by the Stakers contract.
type StakersLockingDelegation struct {
	Delegator common.Address
	StakerID  *big.Int
	FromEpoch *big.Int
	EndTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLockingDelegation is a free log retrieval operation binding the contract event 0x823f252f996e1f519fd0215db7eb4d5a688d78587bf03bfb03d77bfca939806d.
//
// Solidity: event LockingDelegation(address indexed delegator, uint256 indexed stakerID, uint256 fromEpoch, uint256 endTime)
func (_Stakers *StakersFilterer) FilterLockingDelegation(opts *bind.FilterOpts, delegator []common.Address, stakerID []*big.Int) (*StakersLockingDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "LockingDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &StakersLockingDelegationIterator{contract: _Stakers.contract, event: "LockingDelegation", logs: logs, sub: sub}, nil
}

// WatchLockingDelegation is a free log subscription operation binding the contract event 0x823f252f996e1f519fd0215db7eb4d5a688d78587bf03bfb03d77bfca939806d.
//
// Solidity: event LockingDelegation(address indexed delegator, uint256 indexed stakerID, uint256 fromEpoch, uint256 endTime)
func (_Stakers *StakersFilterer) WatchLockingDelegation(opts *bind.WatchOpts, sink chan<- *StakersLockingDelegation, delegator []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "LockingDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersLockingDelegation)
				if err := _Stakers.contract.UnpackLog(event, "LockingDelegation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLockingDelegation is a log parse operation binding the contract event 0x823f252f996e1f519fd0215db7eb4d5a688d78587bf03bfb03d77bfca939806d.
//
// Solidity: event LockingDelegation(address indexed delegator, uint256 indexed stakerID, uint256 fromEpoch, uint256 endTime)
func (_Stakers *StakersFilterer) ParseLockingDelegation(log types.Log) (*StakersLockingDelegation, error) {
	event := new(StakersLockingDelegation)
	if err := _Stakers.contract.UnpackLog(event, "LockingDelegation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersLockingStakeIterator is returned from FilterLockingStake and is used to iterate over the raw logs and unpacked data for LockingStake events raised by the Stakers contract.
type StakersLockingStakeIterator struct {
	Event *StakersLockingStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersLockingStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersLockingStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersLockingStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersLockingStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersLockingStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersLockingStake represents a LockingStake event raised by the Stakers contract.
type StakersLockingStake struct {
	StakerID  *big.Int
	FromEpoch *big.Int
	EndTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLockingStake is a free log retrieval operation binding the contract event 0x71f8e76b11dde805fa567e857c4beba340500f4ca9da003520a25014f542162b.
//
// Solidity: event LockingStake(uint256 indexed stakerID, uint256 fromEpoch, uint256 endTime)
func (_Stakers *StakersFilterer) FilterLockingStake(opts *bind.FilterOpts, stakerID []*big.Int) (*StakersLockingStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "LockingStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &StakersLockingStakeIterator{contract: _Stakers.contract, event: "LockingStake", logs: logs, sub: sub}, nil
}

// WatchLockingStake is a free log subscription operation binding the contract event 0x71f8e76b11dde805fa567e857c4beba340500f4ca9da003520a25014f542162b.
//
// Solidity: event LockingStake(uint256 indexed stakerID, uint256 fromEpoch, uint256 endTime)
func (_Stakers *StakersFilterer) WatchLockingStake(opts *bind.WatchOpts, sink chan<- *StakersLockingStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "LockingStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersLockingStake)
				if err := _Stakers.contract.UnpackLog(event, "LockingStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLockingStake is a log parse operation binding the contract event 0x71f8e76b11dde805fa567e857c4beba340500f4ca9da003520a25014f542162b.
//
// Solidity: event LockingStake(uint256 indexed stakerID, uint256 fromEpoch, uint256 endTime)
func (_Stakers *StakersFilterer) ParseLockingStake(log types.Log) (*StakersLockingStake, error) {
	event := new(StakersLockingStake)
	if err := _Stakers.contract.UnpackLog(event, "LockingStake", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Stakers contract.
type StakersOwnershipTransferredIterator struct {
	Event *StakersOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersOwnershipTransferred represents a OwnershipTransferred event raised by the Stakers contract.
type StakersOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stakers *StakersFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakersOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakersOwnershipTransferredIterator{contract: _Stakers.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stakers *StakersFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakersOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersOwnershipTransferred)
				if err := _Stakers.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stakers *StakersFilterer) ParseOwnershipTransferred(log types.Log) (*StakersOwnershipTransferred, error) {
	event := new(StakersOwnershipTransferred)
	if err := _Stakers.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersPartialWithdrawnByRequestIterator is returned from FilterPartialWithdrawnByRequest and is used to iterate over the raw logs and unpacked data for PartialWithdrawnByRequest events raised by the Stakers contract.
type StakersPartialWithdrawnByRequestIterator struct {
	Event *StakersPartialWithdrawnByRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersPartialWithdrawnByRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersPartialWithdrawnByRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersPartialWithdrawnByRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersPartialWithdrawnByRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersPartialWithdrawnByRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersPartialWithdrawnByRequest represents a PartialWithdrawnByRequest event raised by the Stakers contract.
type StakersPartialWithdrawnByRequest struct {
	Auth       common.Address
	Receiver   common.Address
	StakerID   *big.Int
	WrID       *big.Int
	Delegation bool
	Penalty    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPartialWithdrawnByRequest is a free log retrieval operation binding the contract event 0xd5304dabc5bd47105b6921889d1b528c4b2223250248a916afd129b1c0512ddd.
//
// Solidity: event PartialWithdrawnByRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 penalty)
func (_Stakers *StakersFilterer) FilterPartialWithdrawnByRequest(opts *bind.FilterOpts, auth []common.Address, receiver []common.Address, stakerID []*big.Int) (*StakersPartialWithdrawnByRequestIterator, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "PartialWithdrawnByRequest", authRule, receiverRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &StakersPartialWithdrawnByRequestIterator{contract: _Stakers.contract, event: "PartialWithdrawnByRequest", logs: logs, sub: sub}, nil
}

// WatchPartialWithdrawnByRequest is a free log subscription operation binding the contract event 0xd5304dabc5bd47105b6921889d1b528c4b2223250248a916afd129b1c0512ddd.
//
// Solidity: event PartialWithdrawnByRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 penalty)
func (_Stakers *StakersFilterer) WatchPartialWithdrawnByRequest(opts *bind.WatchOpts, sink chan<- *StakersPartialWithdrawnByRequest, auth []common.Address, receiver []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "PartialWithdrawnByRequest", authRule, receiverRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersPartialWithdrawnByRequest)
				if err := _Stakers.contract.UnpackLog(event, "PartialWithdrawnByRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePartialWithdrawnByRequest is a log parse operation binding the contract event 0xd5304dabc5bd47105b6921889d1b528c4b2223250248a916afd129b1c0512ddd.
//
// Solidity: event PartialWithdrawnByRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 penalty)
func (_Stakers *StakersFilterer) ParsePartialWithdrawnByRequest(log types.Log) (*StakersPartialWithdrawnByRequest, error) {
	event := new(StakersPartialWithdrawnByRequest)
	if err := _Stakers.contract.UnpackLog(event, "PartialWithdrawnByRequest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersPreparedToWithdrawDelegationIterator is returned from FilterPreparedToWithdrawDelegation and is used to iterate over the raw logs and unpacked data for PreparedToWithdrawDelegation events raised by the Stakers contract.
type StakersPreparedToWithdrawDelegationIterator struct {
	Event *StakersPreparedToWithdrawDelegation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersPreparedToWithdrawDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersPreparedToWithdrawDelegation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersPreparedToWithdrawDelegation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersPreparedToWithdrawDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersPreparedToWithdrawDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersPreparedToWithdrawDelegation represents a PreparedToWithdrawDelegation event raised by the Stakers contract.
type StakersPreparedToWithdrawDelegation struct {
	Delegator common.Address
	StakerID  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPreparedToWithdrawDelegation is a free log retrieval operation binding the contract event 0x5b1eea49e405ef6d509836aac841959c30bb0673b1fd70859bfc6ae5e4ee3df2.
//
// Solidity: event PreparedToWithdrawDelegation(address indexed delegator, uint256 indexed stakerID)
func (_Stakers *StakersFilterer) FilterPreparedToWithdrawDelegation(opts *bind.FilterOpts, delegator []common.Address, stakerID []*big.Int) (*StakersPreparedToWithdrawDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "PreparedToWithdrawDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &StakersPreparedToWithdrawDelegationIterator{contract: _Stakers.contract, event: "PreparedToWithdrawDelegation", logs: logs, sub: sub}, nil
}

// WatchPreparedToWithdrawDelegation is a free log subscription operation binding the contract event 0x5b1eea49e405ef6d509836aac841959c30bb0673b1fd70859bfc6ae5e4ee3df2.
//
// Solidity: event PreparedToWithdrawDelegation(address indexed delegator, uint256 indexed stakerID)
func (_Stakers *StakersFilterer) WatchPreparedToWithdrawDelegation(opts *bind.WatchOpts, sink chan<- *StakersPreparedToWithdrawDelegation, delegator []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "PreparedToWithdrawDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersPreparedToWithdrawDelegation)
				if err := _Stakers.contract.UnpackLog(event, "PreparedToWithdrawDelegation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePreparedToWithdrawDelegation is a log parse operation binding the contract event 0x5b1eea49e405ef6d509836aac841959c30bb0673b1fd70859bfc6ae5e4ee3df2.
//
// Solidity: event PreparedToWithdrawDelegation(address indexed delegator, uint256 indexed stakerID)
func (_Stakers *StakersFilterer) ParsePreparedToWithdrawDelegation(log types.Log) (*StakersPreparedToWithdrawDelegation, error) {
	event := new(StakersPreparedToWithdrawDelegation)
	if err := _Stakers.contract.UnpackLog(event, "PreparedToWithdrawDelegation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersPreparedToWithdrawStakeIterator is returned from FilterPreparedToWithdrawStake and is used to iterate over the raw logs and unpacked data for PreparedToWithdrawStake events raised by the Stakers contract.
type StakersPreparedToWithdrawStakeIterator struct {
	Event *StakersPreparedToWithdrawStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersPreparedToWithdrawStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersPreparedToWithdrawStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersPreparedToWithdrawStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersPreparedToWithdrawStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersPreparedToWithdrawStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersPreparedToWithdrawStake represents a PreparedToWithdrawStake event raised by the Stakers contract.
type StakersPreparedToWithdrawStake struct {
	StakerID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPreparedToWithdrawStake is a free log retrieval operation binding the contract event 0x84244546a9da4942f506db48ff90ebc240c73bb399e3e47d58843c6bb60e7185.
//
// Solidity: event PreparedToWithdrawStake(uint256 indexed stakerID)
func (_Stakers *StakersFilterer) FilterPreparedToWithdrawStake(opts *bind.FilterOpts, stakerID []*big.Int) (*StakersPreparedToWithdrawStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "PreparedToWithdrawStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &StakersPreparedToWithdrawStakeIterator{contract: _Stakers.contract, event: "PreparedToWithdrawStake", logs: logs, sub: sub}, nil
}

// WatchPreparedToWithdrawStake is a free log subscription operation binding the contract event 0x84244546a9da4942f506db48ff90ebc240c73bb399e3e47d58843c6bb60e7185.
//
// Solidity: event PreparedToWithdrawStake(uint256 indexed stakerID)
func (_Stakers *StakersFilterer) WatchPreparedToWithdrawStake(opts *bind.WatchOpts, sink chan<- *StakersPreparedToWithdrawStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "PreparedToWithdrawStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersPreparedToWithdrawStake)
				if err := _Stakers.contract.UnpackLog(event, "PreparedToWithdrawStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePreparedToWithdrawStake is a log parse operation binding the contract event 0x84244546a9da4942f506db48ff90ebc240c73bb399e3e47d58843c6bb60e7185.
//
// Solidity: event PreparedToWithdrawStake(uint256 indexed stakerID)
func (_Stakers *StakersFilterer) ParsePreparedToWithdrawStake(log types.Log) (*StakersPreparedToWithdrawStake, error) {
	event := new(StakersPreparedToWithdrawStake)
	if err := _Stakers.contract.UnpackLog(event, "PreparedToWithdrawStake", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersUnstashedRewardsIterator is returned from FilterUnstashedRewards and is used to iterate over the raw logs and unpacked data for UnstashedRewards events raised by the Stakers contract.
type StakersUnstashedRewardsIterator struct {
	Event *StakersUnstashedRewards // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersUnstashedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersUnstashedRewards)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersUnstashedRewards)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersUnstashedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersUnstashedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersUnstashedRewards represents a UnstashedRewards event raised by the Stakers contract.
type StakersUnstashedRewards struct {
	Auth     common.Address
	Receiver common.Address
	Rewards  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnstashedRewards is a free log retrieval operation binding the contract event 0x80b36a0e929d7e7925087e54acfeecf4c6043e451b9d71ac5e908b66f9e5d126.
//
// Solidity: event UnstashedRewards(address indexed auth, address indexed receiver, uint256 rewards)
func (_Stakers *StakersFilterer) FilterUnstashedRewards(opts *bind.FilterOpts, auth []common.Address, receiver []common.Address) (*StakersUnstashedRewardsIterator, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "UnstashedRewards", authRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &StakersUnstashedRewardsIterator{contract: _Stakers.contract, event: "UnstashedRewards", logs: logs, sub: sub}, nil
}

// WatchUnstashedRewards is a free log subscription operation binding the contract event 0x80b36a0e929d7e7925087e54acfeecf4c6043e451b9d71ac5e908b66f9e5d126.
//
// Solidity: event UnstashedRewards(address indexed auth, address indexed receiver, uint256 rewards)
func (_Stakers *StakersFilterer) WatchUnstashedRewards(opts *bind.WatchOpts, sink chan<- *StakersUnstashedRewards, auth []common.Address, receiver []common.Address) (event.Subscription, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "UnstashedRewards", authRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersUnstashedRewards)
				if err := _Stakers.contract.UnpackLog(event, "UnstashedRewards", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstashedRewards is a log parse operation binding the contract event 0x80b36a0e929d7e7925087e54acfeecf4c6043e451b9d71ac5e908b66f9e5d126.
//
// Solidity: event UnstashedRewards(address indexed auth, address indexed receiver, uint256 rewards)
func (_Stakers *StakersFilterer) ParseUnstashedRewards(log types.Log) (*StakersUnstashedRewards, error) {
	event := new(StakersUnstashedRewards)
	if err := _Stakers.contract.UnpackLog(event, "UnstashedRewards", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersUpdatedBaseRewardPerSecIterator is returned from FilterUpdatedBaseRewardPerSec and is used to iterate over the raw logs and unpacked data for UpdatedBaseRewardPerSec events raised by the Stakers contract.
type StakersUpdatedBaseRewardPerSecIterator struct {
	Event *StakersUpdatedBaseRewardPerSec // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersUpdatedBaseRewardPerSecIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersUpdatedBaseRewardPerSec)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersUpdatedBaseRewardPerSec)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersUpdatedBaseRewardPerSecIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersUpdatedBaseRewardPerSecIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersUpdatedBaseRewardPerSec represents a UpdatedBaseRewardPerSec event raised by the Stakers contract.
type StakersUpdatedBaseRewardPerSec struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedBaseRewardPerSec is a free log retrieval operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
//
// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
func (_Stakers *StakersFilterer) FilterUpdatedBaseRewardPerSec(opts *bind.FilterOpts) (*StakersUpdatedBaseRewardPerSecIterator, error) {

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "UpdatedBaseRewardPerSec")
	if err != nil {
		return nil, err
	}
	return &StakersUpdatedBaseRewardPerSecIterator{contract: _Stakers.contract, event: "UpdatedBaseRewardPerSec", logs: logs, sub: sub}, nil
}

// WatchUpdatedBaseRewardPerSec is a free log subscription operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
//
// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
func (_Stakers *StakersFilterer) WatchUpdatedBaseRewardPerSec(opts *bind.WatchOpts, sink chan<- *StakersUpdatedBaseRewardPerSec) (event.Subscription, error) {

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "UpdatedBaseRewardPerSec")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersUpdatedBaseRewardPerSec)
				if err := _Stakers.contract.UnpackLog(event, "UpdatedBaseRewardPerSec", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedBaseRewardPerSec is a log parse operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
//
// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
func (_Stakers *StakersFilterer) ParseUpdatedBaseRewardPerSec(log types.Log) (*StakersUpdatedBaseRewardPerSec, error) {
	event := new(StakersUpdatedBaseRewardPerSec)
	if err := _Stakers.contract.UnpackLog(event, "UpdatedBaseRewardPerSec", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersUpdatedDelegationIterator is returned from FilterUpdatedDelegation and is used to iterate over the raw logs and unpacked data for UpdatedDelegation events raised by the Stakers contract.
type StakersUpdatedDelegationIterator struct {
	Event *StakersUpdatedDelegation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersUpdatedDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersUpdatedDelegation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersUpdatedDelegation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersUpdatedDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersUpdatedDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersUpdatedDelegation represents a UpdatedDelegation event raised by the Stakers contract.
type StakersUpdatedDelegation struct {
	Delegator   common.Address
	OldStakerID *big.Int
	NewStakerID *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedDelegation is a free log retrieval operation binding the contract event 0x19b46b9014e4dc8ca74f505b8921797c6a8a489860217d15b3c7d741637dfcff.
//
// Solidity: event UpdatedDelegation(address indexed delegator, uint256 indexed oldStakerID, uint256 indexed newStakerID, uint256 amount)
func (_Stakers *StakersFilterer) FilterUpdatedDelegation(opts *bind.FilterOpts, delegator []common.Address, oldStakerID []*big.Int, newStakerID []*big.Int) (*StakersUpdatedDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var oldStakerIDRule []interface{}
	for _, oldStakerIDItem := range oldStakerID {
		oldStakerIDRule = append(oldStakerIDRule, oldStakerIDItem)
	}
	var newStakerIDRule []interface{}
	for _, newStakerIDItem := range newStakerID {
		newStakerIDRule = append(newStakerIDRule, newStakerIDItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "UpdatedDelegation", delegatorRule, oldStakerIDRule, newStakerIDRule)
	if err != nil {
		return nil, err
	}
	return &StakersUpdatedDelegationIterator{contract: _Stakers.contract, event: "UpdatedDelegation", logs: logs, sub: sub}, nil
}

// WatchUpdatedDelegation is a free log subscription operation binding the contract event 0x19b46b9014e4dc8ca74f505b8921797c6a8a489860217d15b3c7d741637dfcff.
//
// Solidity: event UpdatedDelegation(address indexed delegator, uint256 indexed oldStakerID, uint256 indexed newStakerID, uint256 amount)
func (_Stakers *StakersFilterer) WatchUpdatedDelegation(opts *bind.WatchOpts, sink chan<- *StakersUpdatedDelegation, delegator []common.Address, oldStakerID []*big.Int, newStakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var oldStakerIDRule []interface{}
	for _, oldStakerIDItem := range oldStakerID {
		oldStakerIDRule = append(oldStakerIDRule, oldStakerIDItem)
	}
	var newStakerIDRule []interface{}
	for _, newStakerIDItem := range newStakerID {
		newStakerIDRule = append(newStakerIDRule, newStakerIDItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "UpdatedDelegation", delegatorRule, oldStakerIDRule, newStakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersUpdatedDelegation)
				if err := _Stakers.contract.UnpackLog(event, "UpdatedDelegation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedDelegation is a log parse operation binding the contract event 0x19b46b9014e4dc8ca74f505b8921797c6a8a489860217d15b3c7d741637dfcff.
//
// Solidity: event UpdatedDelegation(address indexed delegator, uint256 indexed oldStakerID, uint256 indexed newStakerID, uint256 amount)
func (_Stakers *StakersFilterer) ParseUpdatedDelegation(log types.Log) (*StakersUpdatedDelegation, error) {
	event := new(StakersUpdatedDelegation)
	if err := _Stakers.contract.UnpackLog(event, "UpdatedDelegation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersUpdatedGasPowerAllocationRateIterator is returned from FilterUpdatedGasPowerAllocationRate and is used to iterate over the raw logs and unpacked data for UpdatedGasPowerAllocationRate events raised by the Stakers contract.
type StakersUpdatedGasPowerAllocationRateIterator struct {
	Event *StakersUpdatedGasPowerAllocationRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersUpdatedGasPowerAllocationRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersUpdatedGasPowerAllocationRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersUpdatedGasPowerAllocationRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersUpdatedGasPowerAllocationRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersUpdatedGasPowerAllocationRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersUpdatedGasPowerAllocationRate represents a UpdatedGasPowerAllocationRate event raised by the Stakers contract.
type StakersUpdatedGasPowerAllocationRate struct {
	Short *big.Int
	Long  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGasPowerAllocationRate is a free log retrieval operation binding the contract event 0x95ae5488127de4bc98492f4487556e7af9f37eb4b6d5e94f6d849e03ff76cc7c.
//
// Solidity: event UpdatedGasPowerAllocationRate(uint256 short, uint256 long)
func (_Stakers *StakersFilterer) FilterUpdatedGasPowerAllocationRate(opts *bind.FilterOpts) (*StakersUpdatedGasPowerAllocationRateIterator, error) {

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "UpdatedGasPowerAllocationRate")
	if err != nil {
		return nil, err
	}
	return &StakersUpdatedGasPowerAllocationRateIterator{contract: _Stakers.contract, event: "UpdatedGasPowerAllocationRate", logs: logs, sub: sub}, nil
}

// WatchUpdatedGasPowerAllocationRate is a free log subscription operation binding the contract event 0x95ae5488127de4bc98492f4487556e7af9f37eb4b6d5e94f6d849e03ff76cc7c.
//
// Solidity: event UpdatedGasPowerAllocationRate(uint256 short, uint256 long)
func (_Stakers *StakersFilterer) WatchUpdatedGasPowerAllocationRate(opts *bind.WatchOpts, sink chan<- *StakersUpdatedGasPowerAllocationRate) (event.Subscription, error) {

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "UpdatedGasPowerAllocationRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersUpdatedGasPowerAllocationRate)
				if err := _Stakers.contract.UnpackLog(event, "UpdatedGasPowerAllocationRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedGasPowerAllocationRate is a log parse operation binding the contract event 0x95ae5488127de4bc98492f4487556e7af9f37eb4b6d5e94f6d849e03ff76cc7c.
//
// Solidity: event UpdatedGasPowerAllocationRate(uint256 short, uint256 long)
func (_Stakers *StakersFilterer) ParseUpdatedGasPowerAllocationRate(log types.Log) (*StakersUpdatedGasPowerAllocationRate, error) {
	event := new(StakersUpdatedGasPowerAllocationRate)
	if err := _Stakers.contract.UnpackLog(event, "UpdatedGasPowerAllocationRate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersUpdatedStakeIterator is returned from FilterUpdatedStake and is used to iterate over the raw logs and unpacked data for UpdatedStake events raised by the Stakers contract.
type StakersUpdatedStakeIterator struct {
	Event *StakersUpdatedStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersUpdatedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersUpdatedStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersUpdatedStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersUpdatedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersUpdatedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersUpdatedStake represents a UpdatedStake event raised by the Stakers contract.
type StakersUpdatedStake struct {
	StakerID    *big.Int
	Amount      *big.Int
	DelegatedMe *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStake is a free log retrieval operation binding the contract event 0x509404fa75ce234a1273cf9f7918bcf54e0ef19f2772e4f71b6526606a723b7c.
//
// Solidity: event UpdatedStake(uint256 indexed stakerID, uint256 amount, uint256 delegatedMe)
func (_Stakers *StakersFilterer) FilterUpdatedStake(opts *bind.FilterOpts, stakerID []*big.Int) (*StakersUpdatedStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "UpdatedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &StakersUpdatedStakeIterator{contract: _Stakers.contract, event: "UpdatedStake", logs: logs, sub: sub}, nil
}

// WatchUpdatedStake is a free log subscription operation binding the contract event 0x509404fa75ce234a1273cf9f7918bcf54e0ef19f2772e4f71b6526606a723b7c.
//
// Solidity: event UpdatedStake(uint256 indexed stakerID, uint256 amount, uint256 delegatedMe)
func (_Stakers *StakersFilterer) WatchUpdatedStake(opts *bind.WatchOpts, sink chan<- *StakersUpdatedStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "UpdatedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersUpdatedStake)
				if err := _Stakers.contract.UnpackLog(event, "UpdatedStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedStake is a log parse operation binding the contract event 0x509404fa75ce234a1273cf9f7918bcf54e0ef19f2772e4f71b6526606a723b7c.
//
// Solidity: event UpdatedStake(uint256 indexed stakerID, uint256 amount, uint256 delegatedMe)
func (_Stakers *StakersFilterer) ParseUpdatedStake(log types.Log) (*StakersUpdatedStake, error) {
	event := new(StakersUpdatedStake)
	if err := _Stakers.contract.UnpackLog(event, "UpdatedStake", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersUpdatedStakerMetadataIterator is returned from FilterUpdatedStakerMetadata and is used to iterate over the raw logs and unpacked data for UpdatedStakerMetadata events raised by the Stakers contract.
type StakersUpdatedStakerMetadataIterator struct {
	Event *StakersUpdatedStakerMetadata // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersUpdatedStakerMetadataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersUpdatedStakerMetadata)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersUpdatedStakerMetadata)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersUpdatedStakerMetadataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersUpdatedStakerMetadataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersUpdatedStakerMetadata represents a UpdatedStakerMetadata event raised by the Stakers contract.
type StakersUpdatedStakerMetadata struct {
	StakerID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStakerMetadata is a free log retrieval operation binding the contract event 0xb7a99a0df6a9e15c2689e6a55811ef76cdb514c67d4a0e37fcb125ada0e3cd83.
//
// Solidity: event UpdatedStakerMetadata(uint256 indexed stakerID)
func (_Stakers *StakersFilterer) FilterUpdatedStakerMetadata(opts *bind.FilterOpts, stakerID []*big.Int) (*StakersUpdatedStakerMetadataIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "UpdatedStakerMetadata", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &StakersUpdatedStakerMetadataIterator{contract: _Stakers.contract, event: "UpdatedStakerMetadata", logs: logs, sub: sub}, nil
}

// WatchUpdatedStakerMetadata is a free log subscription operation binding the contract event 0xb7a99a0df6a9e15c2689e6a55811ef76cdb514c67d4a0e37fcb125ada0e3cd83.
//
// Solidity: event UpdatedStakerMetadata(uint256 indexed stakerID)
func (_Stakers *StakersFilterer) WatchUpdatedStakerMetadata(opts *bind.WatchOpts, sink chan<- *StakersUpdatedStakerMetadata, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "UpdatedStakerMetadata", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersUpdatedStakerMetadata)
				if err := _Stakers.contract.UnpackLog(event, "UpdatedStakerMetadata", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedStakerMetadata is a log parse operation binding the contract event 0xb7a99a0df6a9e15c2689e6a55811ef76cdb514c67d4a0e37fcb125ada0e3cd83.
//
// Solidity: event UpdatedStakerMetadata(uint256 indexed stakerID)
func (_Stakers *StakersFilterer) ParseUpdatedStakerMetadata(log types.Log) (*StakersUpdatedStakerMetadata, error) {
	event := new(StakersUpdatedStakerMetadata)
	if err := _Stakers.contract.UnpackLog(event, "UpdatedStakerMetadata", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersUpdatedStakerSfcAddressIterator is returned from FilterUpdatedStakerSfcAddress and is used to iterate over the raw logs and unpacked data for UpdatedStakerSfcAddress events raised by the Stakers contract.
type StakersUpdatedStakerSfcAddressIterator struct {
	Event *StakersUpdatedStakerSfcAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersUpdatedStakerSfcAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersUpdatedStakerSfcAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersUpdatedStakerSfcAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersUpdatedStakerSfcAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersUpdatedStakerSfcAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersUpdatedStakerSfcAddress represents a UpdatedStakerSfcAddress event raised by the Stakers contract.
type StakersUpdatedStakerSfcAddress struct {
	StakerID      *big.Int
	OldSfcAddress common.Address
	NewSfcAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStakerSfcAddress is a free log retrieval operation binding the contract event 0x7cc102ee500cbca85691c9642080562e8f012b04d27f5b7f389453672b206946.
//
// Solidity: event UpdatedStakerSfcAddress(uint256 indexed stakerID, address indexed oldSfcAddress, address indexed newSfcAddress)
func (_Stakers *StakersFilterer) FilterUpdatedStakerSfcAddress(opts *bind.FilterOpts, stakerID []*big.Int, oldSfcAddress []common.Address, newSfcAddress []common.Address) (*StakersUpdatedStakerSfcAddressIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}
	var oldSfcAddressRule []interface{}
	for _, oldSfcAddressItem := range oldSfcAddress {
		oldSfcAddressRule = append(oldSfcAddressRule, oldSfcAddressItem)
	}
	var newSfcAddressRule []interface{}
	for _, newSfcAddressItem := range newSfcAddress {
		newSfcAddressRule = append(newSfcAddressRule, newSfcAddressItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "UpdatedStakerSfcAddress", stakerIDRule, oldSfcAddressRule, newSfcAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakersUpdatedStakerSfcAddressIterator{contract: _Stakers.contract, event: "UpdatedStakerSfcAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedStakerSfcAddress is a free log subscription operation binding the contract event 0x7cc102ee500cbca85691c9642080562e8f012b04d27f5b7f389453672b206946.
//
// Solidity: event UpdatedStakerSfcAddress(uint256 indexed stakerID, address indexed oldSfcAddress, address indexed newSfcAddress)
func (_Stakers *StakersFilterer) WatchUpdatedStakerSfcAddress(opts *bind.WatchOpts, sink chan<- *StakersUpdatedStakerSfcAddress, stakerID []*big.Int, oldSfcAddress []common.Address, newSfcAddress []common.Address) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}
	var oldSfcAddressRule []interface{}
	for _, oldSfcAddressItem := range oldSfcAddress {
		oldSfcAddressRule = append(oldSfcAddressRule, oldSfcAddressItem)
	}
	var newSfcAddressRule []interface{}
	for _, newSfcAddressItem := range newSfcAddress {
		newSfcAddressRule = append(newSfcAddressRule, newSfcAddressItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "UpdatedStakerSfcAddress", stakerIDRule, oldSfcAddressRule, newSfcAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersUpdatedStakerSfcAddress)
				if err := _Stakers.contract.UnpackLog(event, "UpdatedStakerSfcAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedStakerSfcAddress is a log parse operation binding the contract event 0x7cc102ee500cbca85691c9642080562e8f012b04d27f5b7f389453672b206946.
//
// Solidity: event UpdatedStakerSfcAddress(uint256 indexed stakerID, address indexed oldSfcAddress, address indexed newSfcAddress)
func (_Stakers *StakersFilterer) ParseUpdatedStakerSfcAddress(log types.Log) (*StakersUpdatedStakerSfcAddress, error) {
	event := new(StakersUpdatedStakerSfcAddress)
	if err := _Stakers.contract.UnpackLog(event, "UpdatedStakerSfcAddress", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersWithdrawnDelegationIterator is returned from FilterWithdrawnDelegation and is used to iterate over the raw logs and unpacked data for WithdrawnDelegation events raised by the Stakers contract.
type StakersWithdrawnDelegationIterator struct {
	Event *StakersWithdrawnDelegation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersWithdrawnDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersWithdrawnDelegation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersWithdrawnDelegation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersWithdrawnDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersWithdrawnDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersWithdrawnDelegation represents a WithdrawnDelegation event raised by the Stakers contract.
type StakersWithdrawnDelegation struct {
	Delegator common.Address
	StakerID  *big.Int
	Penalty   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnDelegation is a free log retrieval operation binding the contract event 0x87e86b3710b72c10173ca52c6a9f9cf2df27e77ed177741a8b4feb12bb7a606f.
//
// Solidity: event WithdrawnDelegation(address indexed delegator, uint256 indexed stakerID, uint256 penalty)
func (_Stakers *StakersFilterer) FilterWithdrawnDelegation(opts *bind.FilterOpts, delegator []common.Address, stakerID []*big.Int) (*StakersWithdrawnDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "WithdrawnDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &StakersWithdrawnDelegationIterator{contract: _Stakers.contract, event: "WithdrawnDelegation", logs: logs, sub: sub}, nil
}

// WatchWithdrawnDelegation is a free log subscription operation binding the contract event 0x87e86b3710b72c10173ca52c6a9f9cf2df27e77ed177741a8b4feb12bb7a606f.
//
// Solidity: event WithdrawnDelegation(address indexed delegator, uint256 indexed stakerID, uint256 penalty)
func (_Stakers *StakersFilterer) WatchWithdrawnDelegation(opts *bind.WatchOpts, sink chan<- *StakersWithdrawnDelegation, delegator []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "WithdrawnDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersWithdrawnDelegation)
				if err := _Stakers.contract.UnpackLog(event, "WithdrawnDelegation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawnDelegation is a log parse operation binding the contract event 0x87e86b3710b72c10173ca52c6a9f9cf2df27e77ed177741a8b4feb12bb7a606f.
//
// Solidity: event WithdrawnDelegation(address indexed delegator, uint256 indexed stakerID, uint256 penalty)
func (_Stakers *StakersFilterer) ParseWithdrawnDelegation(log types.Log) (*StakersWithdrawnDelegation, error) {
	event := new(StakersWithdrawnDelegation)
	if err := _Stakers.contract.UnpackLog(event, "WithdrawnDelegation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakersWithdrawnStakeIterator is returned from FilterWithdrawnStake and is used to iterate over the raw logs and unpacked data for WithdrawnStake events raised by the Stakers contract.
type StakersWithdrawnStakeIterator struct {
	Event *StakersWithdrawnStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakersWithdrawnStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakersWithdrawnStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakersWithdrawnStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakersWithdrawnStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakersWithdrawnStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakersWithdrawnStake represents a WithdrawnStake event raised by the Stakers contract.
type StakersWithdrawnStake struct {
	StakerID *big.Int
	Penalty  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnStake is a free log retrieval operation binding the contract event 0x8c6548258f8f12a9d4b593fa89a223417ed901d4ee9712ba09beb4d56f5262b6.
//
// Solidity: event WithdrawnStake(uint256 indexed stakerID, uint256 penalty)
func (_Stakers *StakersFilterer) FilterWithdrawnStake(opts *bind.FilterOpts, stakerID []*big.Int) (*StakersWithdrawnStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.FilterLogs(opts, "WithdrawnStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &StakersWithdrawnStakeIterator{contract: _Stakers.contract, event: "WithdrawnStake", logs: logs, sub: sub}, nil
}

// WatchWithdrawnStake is a free log subscription operation binding the contract event 0x8c6548258f8f12a9d4b593fa89a223417ed901d4ee9712ba09beb4d56f5262b6.
//
// Solidity: event WithdrawnStake(uint256 indexed stakerID, uint256 penalty)
func (_Stakers *StakersFilterer) WatchWithdrawnStake(opts *bind.WatchOpts, sink chan<- *StakersWithdrawnStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _Stakers.contract.WatchLogs(opts, "WithdrawnStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakersWithdrawnStake)
				if err := _Stakers.contract.UnpackLog(event, "WithdrawnStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawnStake is a log parse operation binding the contract event 0x8c6548258f8f12a9d4b593fa89a223417ed901d4ee9712ba09beb4d56f5262b6.
//
// Solidity: event WithdrawnStake(uint256 indexed stakerID, uint256 penalty)
func (_Stakers *StakersFilterer) ParseWithdrawnStake(log types.Log) (*StakersWithdrawnStake, error) {
	event := new(StakersWithdrawnStake)
	if err := _Stakers.contract.UnpackLog(event, "WithdrawnStake", log); err != nil {
		return nil, err
	}
	return event, nil
}
