package keeper

import (
	"cosmos-llm/x/inference/types"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) AppendInferenceRun(ctx sdk.Context, run types.Inferencerun) uint64 {
	count := k.GetInferenceCount(ctx)

	run.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InferenceKey))
	appendedValue := k.cdc.MustMarshal(&run)
	store.Set(GetIDBytes(run.Id), appendedValue)
	k.SetInferencerunCount(ctx, count+1)
	return count
}

func (k Keeper) GetInferenceCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.InferenceCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetInferencerunCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.InferenceCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetInferencerun(ctx sdk.Context, id uint64) (val types.Inferencerun, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InferenceKey))
	b := store.Get(GetIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

//// TODO: validate this
//func (k Keeper) UpdateAgentRun(ctx sdk.Context, agentRun types.Agentrun) {
//	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
//	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AgentrunKey))
//	b := k.cdc.MustMarshal(&agentRun)
//	store.Set(GetIDBytes(agentRun.Id), b)
//}
