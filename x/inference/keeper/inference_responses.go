package keeper

import (
	"cosmos-llm/x/inference/types"
	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) InsertInferenceRunResponse(ctx sdk.Context, runId uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InferenceResponseKey))

	run := types.InferencerunResponse{
		RunId:      runId,
		Responses:  []string{},
		Responders: []string{},
		Isfinished: false,
	}

	appendedValue := k.cdc.MustMarshal(&run)
	store.Set(GetIDBytes(run.RunId), appendedValue)
}

func (k Keeper) AppendInferenceRunResponse(
	ctx sdk.Context,
	runId uint64,
	response string,
	responder string,
) error {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InferenceResponseKey))

	runResponse, found := k.GetInferencerunResponse(ctx, runId)
	if !found {
		return sdkerrors.Wrapf(types.ErrRunNotFound, "%d", runId)
	}
	for _, element := range runResponse.Responders {
		if element == responder {
			return sdkerrors.Wrapf(types.ErrAlreadyRegistered, "%s %d", responder, runId)
		}
	}

	_ = append(runResponse.Responses, response)
	_ = append(runResponse.Responses, responder)

	appendedValue := k.cdc.MustMarshal(&runResponse)
	store.Set(GetIDBytes(runResponse.RunId), appendedValue)
	return nil
}

func (k Keeper) GetInferencerunResponse(ctx sdk.Context, runId uint64) (val types.InferencerunResponse, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InferenceResponseKey))
	b := store.Get(GetIDBytes(runId))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
