package keeper

import (
	"cosmos-llm/x/inference/types"
	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	types2 "cosmossdk.io/store/types"
	"fmt"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
)

func (k Keeper) InsertInferenceRunResponse(ctx sdk.Context, runId uint64, modelId uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InferenceResponseKey))

	run := types.InferencerunResponse{
		RunId:      runId,
		ModelId:    modelId,
		Responses:  []string{},
		Responders: []string{},
		Isfinished: false,
	}
	runDbID := uuid.New().String()
	run.Id = runDbID

	appendedValue := k.cdc.MustMarshal(&run)
	key := append(types.KeyPrefix(types.InferenceResponseKey), []byte(run.Id)...)
	store.Set(key, appendedValue)
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

func (k Keeper) GetUnansweredInferencerunResponse(ctx sdk.Context, modelId uint64) (val types.Inferencerun, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InferenceResponseKey))
	iterator := types2.KVStorePrefixIterator(store, types.KeyPrefix(types.InferenceResponseKey))
	defer iterator.Close()

	fmt.Println("======================================================== Iterating through responses")
	for ; iterator.Valid(); iterator.Next() {
		var run types.InferencerunResponse
		k.cdc.MustUnmarshal(iterator.Value(), &run)
		fmt.Printf("======================================================== %d\n", run.ModelId)
		if run.ModelId == modelId && !run.Isfinished {
			return k.GetInferencerun(ctx, run.RunId)
		}
	}
	return types.Inferencerun{}, false
}
