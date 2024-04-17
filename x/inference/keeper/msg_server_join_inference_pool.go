package keeper

import (
	"context"
	"cosmos-llm/x/inference/types"
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) JoinInferencePool(goCtx context.Context, msg *types.MsgJoinInferencePool) (*types.MsgJoinInferencePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, exists := GetModelById(msg.ModelId)
	if !exists {
		//sdkerrors.Wrapf(types.ErrCreatorNotPlayer, "%s", msg.Creator)
		return nil, sdkerrors.Wrapf(types.ErrInvalidModel, "%d", int(msg.ModelId))
	}
	var staked = types.StakedGpuNode{
		Owner:        msg.Creator,
		Stake:        1,
		Denom:        "stake",
		SuccessCount: 0,
		FailCount:    0,
		ModelId:      msg.ModelId,
	}
	k.AddGpuNode(ctx, staked)

	return &types.MsgJoinInferencePoolResponse{}, nil
}
