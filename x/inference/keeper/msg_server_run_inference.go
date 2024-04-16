package keeper

import (
	"context"

	"cosmos-llm/x/inference/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RunInference(goCtx context.Context, msg *types.MsgRunInference) (*types.MsgRunInferenceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, exists := GetModelById(msg.Modelid)
	if !exists {
		return &types.MsgRunInferenceResponse{}, nil
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
	return &types.MsgRunInferenceResponse{
		Id: id,
	}, nil
}
