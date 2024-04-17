package keeper

import (
	"cosmos-llm/x/inference/types"
	"cosmossdk.io/store/prefix"
	types2 "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
)

func (k Keeper) AddGpuNode(ctx sdk.Context, run types.StakedGpuNode) string {
	// Generate a UUID for the new run.
	runID := uuid.New().String()
	run.Id = runID

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GpuNodeKey))
	b := k.cdc.MustMarshal(&run)
	store.Set([]byte(runID), b)

	return runID
}

func (k Keeper) DeleteInferenceRun(ctx sdk.Context, id string) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GpuNodeKey))
	store.Delete([]byte(id))
}

func (k Keeper) ListAllInferenceRuns(ctx sdk.Context) []types.Inferencerun {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GpuNodeKey))
	iterator := types2.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	var runs []types.Inferencerun
	for ; iterator.Valid(); iterator.Next() {
		var run types.Inferencerun
		k.cdc.MustUnmarshal(iterator.Value(), &run)
		runs = append(runs, run)
	}
	return runs
}
