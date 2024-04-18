package keeper

import (
	"context"
	"cosmos-llm/x/inference/types"
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RunInference(goCtx context.Context, msg *types.MsgRunInference) (*types.MsgRunInferenceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, exists := GetModelById(msg.Modelid)
	if !exists {
		return nil, sdkerrors.Wrapf(types.ErrInvalidModel, "%d", msg.Modelid)
	}
	var question = types.Inferencerun{
		ModelId:    msg.Modelid,
		Prompt:     msg.Prompt,
		Isfinished: false,
	}
	id := k.AppendInferenceRun(
		ctx,
		question,
	)
	k.InsertInferenceRunResponse(ctx, id, msg.Modelid)
	// Need ante handler for this..
	//ctx.GasMeter().ConsumeGas(types.RunInferenceGas, "Run inference")
	return &types.MsgRunInferenceResponse{
		Id: id,
	}, nil
}
