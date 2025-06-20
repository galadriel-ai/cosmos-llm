package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "cosmos-llm/testutil/keeper"
	"cosmos-llm/x/inference/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.InferenceKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
