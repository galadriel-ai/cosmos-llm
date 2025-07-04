package simulation

import (
	"math/rand"

	"cosmos-llm/x/inference/keeper"
	"cosmos-llm/x/inference/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgJoinInferencePool(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgJoinInferencePool{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the JoinInferencePool simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "JoinInferencePool simulation not implemented"), nil, nil
	}
}
