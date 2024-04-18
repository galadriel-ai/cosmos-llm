package keeper

import (
	"context"
	"cosmos-llm/x/inference/types"
	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const StakeAmount = 1000
const StakeDenom = "token"

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
	coins := sdk.NewCoin(StakeDenom, math.NewInt(StakeAmount))
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, types.ModuleName, sdk.NewCoins(coins))
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrFailedToPay, "%s", msg.Creator)
	}

	nodes := k.ListAllGpuNodes(ctx)
	for _, element := range nodes {
		if element.Owner == msg.Creator && element.ModelId == msg.ModelId {
			return nil, sdkerrors.Wrapf(types.ErrAlreadyRegistered, "%s %d", msg.Creator, msg.ModelId)
		}
	}

	var staked = types.StakedGpuNode{
		Owner:        msg.Creator,
		Stake:        1,
		Denom:        StakeDenom,
		SuccessCount: 0,
		FailCount:    0,
		ModelId:      msg.ModelId,
	}
	k.AddGpuNode(ctx, staked)

	return &types.MsgJoinInferencePoolResponse{}, nil
}
