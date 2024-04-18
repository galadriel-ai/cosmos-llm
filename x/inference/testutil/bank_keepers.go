package testutil

import (
	"context"
	"cosmos-llm/x/inference/types"
	"cosmossdk.io/math"
	"go.uber.org/mock/gomock"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func coinsOf(amount uint64, denom string) sdk.Coins {
	return sdk.Coins{
		sdk.Coin{
			Denom:  denom,
			Amount: math.NewInt(int64(amount)),
		},
	}
}

func (escrow *MockBankKeeper) ExpectAny(context context.Context) {
	escrow.EXPECT().SendCoinsFromAccountToModule(sdk.UnwrapSDKContext(context), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	escrow.EXPECT().SendCoinsFromModuleToAccount(sdk.UnwrapSDKContext(context), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
}

func (escrow *MockBankKeeper) ExpectPayWithDenom(context context.Context, who string, amount uint64, denom string) *gomock.Call {
	whoAddr, err := sdk.AccAddressFromBech32(who)
	if err != nil {
		panic(err)
	}
	return escrow.EXPECT().SendCoinsFromAccountToModule(sdk.UnwrapSDKContext(context), whoAddr, types.ModuleName, coinsOf(amount, denom))
}
