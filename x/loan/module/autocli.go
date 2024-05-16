package loan

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "loan/api/loan/loan"
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
					RpcMethod: "LoanAll",
					Use:       "list-loan",
					Short:     "List all loan",
				},
				{
					RpcMethod:      "Loan",
					Use:            "show-loan [id]",
					Short:          "Shows a loan by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
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
					RpcMethod:      "RequestLoan",
					Use:            "request-loan [amount] [fee] [collateral] [deadline]",
					Short:          "Send a request-loan tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}, {ProtoField: "fee"}, {ProtoField: "collateral"}, {ProtoField: "deadline"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
