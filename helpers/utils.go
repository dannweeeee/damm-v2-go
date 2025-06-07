package helpers

import (
	"github.com/dannwee/dbc-go/common"
	"lukechampine.com/uint128"
)

// GetUnclaimReward calculates the unclaimed rewards for a position
func GetUnclaimReward(poolState *common.Pool, positionState *common.PositionState) (*common.UnclaimReward, error) {
	// Calculate total position liquidity
	totalPositionLiquidity := positionState.UnlockedLiquidity.
		Add(positionState.VestedLiquidity).
		Add(positionState.PermanentLockedLiquidity)

	// Convert byte arrays to uint128
	feeAPerLiquidity := uint128.FromBytes(poolState.FeeAPerLiquidity[:])
	feeBPerLiquidity := uint128.FromBytes(poolState.FeeBPerLiquidity[:])
	feeAPerTokenCheckpoint := uint128.FromBytes(positionState.FeeAPerTokenCheckpoint[:])
	feeBPerTokenCheckpoint := uint128.FromBytes(positionState.FeeBPerTokenCheckpoint[:])

	// Calculate fee per token stored for token A
	feeAPerTokenStored := feeAPerLiquidity.Sub(feeAPerTokenCheckpoint)

	// Calculate fee per token stored for token B
	feeBPerTokenStored := feeBPerLiquidity.Sub(feeBPerTokenCheckpoint)

	// Calculate fees with safe multiplication
	// First divide by 2^32 to prevent overflow, then multiply
	feeA := totalPositionLiquidity.Rsh(32).Mul(feeAPerTokenStored).Rsh(common.LIQUIDITY_SCALE - 32)
	feeB := totalPositionLiquidity.Rsh(32).Mul(feeBPerTokenStored).Rsh(common.LIQUIDITY_SCALE - 32)

	// Convert pending fees to uint128
	feeAPending := uint128.From64(positionState.FeeAPending)
	feeBPending := uint128.From64(positionState.FeeBPending)

	// Calculate total fees including pending
	totalFeeA := feeAPending.Add(feeA)
	totalFeeB := feeBPending.Add(feeB)

	// Get rewards from reward infos
	rewards := make([]uint128.Uint128, 0)
	if len(positionState.RewardInfos) > 0 {
		for _, info := range positionState.RewardInfos {
			rewards = append(rewards, uint128.From64(info.RewardPendings))
		}
	}

	return &common.UnclaimReward{
		FeeTokenA: totalFeeA,
		FeeTokenB: totalFeeB,
		Rewards:   rewards,
	}, nil
}
