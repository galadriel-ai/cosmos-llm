package keeper

import (
	"context"
	"fmt"

	"cosmos-llm/x/inference/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetPoolSize(goCtx context.Context, req *types.QueryGetPoolSizeRequest) (*types.QueryGetPoolSizeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	nodes := k.ListAllGpuNodes(ctx)
	fmt.Printf("NODES COUNT: %d\n", len(nodes))

	return &types.QueryGetPoolSizeResponse{
		Size_: uint64(len(nodes)),
	}, nil
}
