package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetWagerCoin(stakeInput uint64) (stake sdk.Coin) {
	return sdk.NewCoin(sdk.DefaultBondDenom, math.NewIntFromUint64(stakeInput))
}
