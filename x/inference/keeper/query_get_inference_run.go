package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"cosmos-llm/x/inference/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetInferenceRun(goCtx context.Context, req *types.QueryGetInferenceRunRequest) (*types.QueryGetInferenceRunResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	run, found := k.GetInferencerun(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	formattedRun := types.Inferencerun{
		Id:         run.Id,
		Prompt:     run.Prompt,
		Response:   run.Response,
		Isfinished: run.Isfinished,
	}
	return &types.QueryGetInferenceRunResponse{Inference: &formattedRun}, nil
}
