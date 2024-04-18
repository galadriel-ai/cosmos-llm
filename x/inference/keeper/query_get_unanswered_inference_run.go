package keeper

import (
	"context"
	sdkerrors "cosmossdk.io/errors"

	"cosmos-llm/x/inference/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetUnansweredInferenceRun(
	goCtx context.Context, req *types.QueryGetUnansweredInferenceRunRequest,
) (*types.QueryGetUnansweredInferenceRunResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	run, found := k.GetUnansweredInferencerunResponse(ctx, req.ModelId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNoUnansweredRuns, "%d", req.ModelId)
	}

	formattedRun := types.Inferencerun{
		Id:         run.Id,
		Prompt:     run.Prompt,
		ModelId:    run.ModelId,
		Response:   run.Response,
		Isfinished: run.Isfinished,
	}
	return &types.QueryGetUnansweredInferenceRunResponse{Inference: &formattedRun}, nil
}
