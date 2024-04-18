package inference

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "cosmos-llm/api/cosmosllm/inference"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "GetInferenceRun",
					Use:            "get-inference-run [id]",
					Short:          "Query get-inference-run",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "GetPoolSize",
					Use:            "get-pool-size",
					Short:          "Query get-pool-size",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "GetUnansweredInferenceRun",
					Use:            "get-unanswered-inference-run [model-id]",
					Short:          "Query get-unanswered-inference-run",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "modelId"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "RunInference",
					Use:            "run-inference [prompt] [modelid]",
					Short:          "Send a runInference tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "prompt"}, {ProtoField: "modelid"}},
				},
				{
					RpcMethod:      "JoinInferencePool",
					Use:            "join-inference-pool [model-id]",
					Short:          "Send a joinInferencePool tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "modelId"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
