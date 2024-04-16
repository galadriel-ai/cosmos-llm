package keeper

import (
	"cosmos-llm/x/inference/types"
)

var _ types.QueryServer = Keeper{}
