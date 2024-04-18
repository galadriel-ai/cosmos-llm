package keeper

import (
	"context"
	"cosmos-llm/x/inference/types"
	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) JoinInferencePool(goCtx context.Context, msg *types.MsgJoinInferencePool) (*types.MsgJoinInferencePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, exists := GetModelById(msg.ModelId)
	if !exists {
		//sdkerrors.Wrapf(types.ErrCreatorNotPlayer, "%s", msg.Creator)
		return nil, sdkerrors.Wrapf(types.ErrInvalidModel, "%d", int(msg.ModelId))
	}

	address, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidAddress, "%s", msg.Creator)
	}

	// Coin denomination
	denom := "token"
	// Amount = 1stake
	coins := sdk.NewCoin(denom, math.NewInt(1))
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, types.ModuleName, sdk.NewCoins(coins))
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrFailedToPay, "%s", msg.Creator)
	}

	var staked = types.StakedGpuNode{
		Owner:        msg.Creator,
		Stake:        1,
		Denom:        denom,
		SuccessCount: 0,
		FailCount:    0,
		ModelId:      msg.ModelId,
	}
	k.AddGpuNode(ctx, staked)

	return &types.MsgJoinInferencePoolResponse{}, nil
}
