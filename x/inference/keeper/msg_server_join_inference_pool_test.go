package keeper_test

import (
	"context"
	keepertest "cosmos-llm/testutil/keeper"
	"cosmos-llm/x/inference/keeper"
	inference "cosmos-llm/x/inference/module"
	"cosmos-llm/x/inference/testutil"
	"cosmos-llm/x/inference/types"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func setupMsgServerCreateDependencies(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context, *testutil.MockBankKeeper) {
	server, k, context, _, bank := setupMsgServerJoinPoolWithMock(t)
	bank.ExpectAny(context)
	return server, k, context, bank
}

func setupMsgServerJoinPoolWithMock(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context,
	*gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.InferenceKeeperWithMocks(t, bankMock)
	inference.InitGenesis(ctx, k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(k)
	context := sdk.WrapSDKContext(ctx)
	return server, k, context, ctrl, bankMock
}

func TestJoinPool(t *testing.T) {
	msgServer, _, ctx, _ := setupMsgServerCreateDependencies(t)
	joinResponse, err := msgServer.JoinInferencePool(ctx, &types.MsgJoinInferencePool{
		Creator: alice,
		ModelId: 1,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgJoinInferencePoolResponse{}, *joinResponse)
}

func TestJoinPoolInvalidModelId(t *testing.T) {
	msgServer, _, ctx, _ := setupMsgServerCreateDependencies(t)
	joinResponse, err := msgServer.JoinInferencePool(ctx, &types.MsgJoinInferencePool{
		Creator: alice,
		ModelId: 1337,
	})
	require.Nil(t, joinResponse)
	require.Equal(t,
		"1337: model id invalid",
		err.Error())
}

func TestJoinPoolTakesStake(t *testing.T) {
	msgServer, _, ctx, _, bank := setupMsgServerJoinPoolWithMock(t)
	bank.ExpectPayWithDenom(ctx, alice, keeper.StakeAmount, keeper.StakeDenom).Times(1)

	_, err := msgServer.JoinInferencePool(ctx, &types.MsgJoinInferencePool{
		Creator: alice,
		ModelId: 1,
	})
	require.Nil(t, err)
}

func TestJoinPoolTwiceErr(t *testing.T) {
	msgServer, _, ctx, _ := setupMsgServerCreateDependencies(t)
	msgServer.JoinInferencePool(ctx, &types.MsgJoinInferencePool{
		Creator: alice,
		ModelId: 1,
	})
	joinResponse, err := msgServer.JoinInferencePool(ctx, &types.MsgJoinInferencePool{
		Creator: alice,
		ModelId: 1,
	})
	require.Nil(t, joinResponse)
	require.Equal(t,
		fmt.Sprintf("%s 1: model id already registered with given account", alice),
		err.Error())
}
