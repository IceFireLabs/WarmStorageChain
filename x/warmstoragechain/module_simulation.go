package warmstoragechain

import (
	"math/rand"

	"github.com/BlockCraftsman/WarmStorageChain/testutil/sample"
	warmstoragechainsimulation "github.com/BlockCraftsman/WarmStorageChain/x/warmstoragechain/simulation"
	"github.com/BlockCraftsman/WarmStorageChain/x/warmstoragechain/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = warmstoragechainsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateResource = "op_weight_msg_resource"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateResource int = 100

	opWeightMsgUpdateResource = "op_weight_msg_resource"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateResource int = 100

	opWeightMsgDeleteResource = "op_weight_msg_resource"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteResource int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	warmstoragechainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ResourceList: []types.Resource{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&warmstoragechainGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateResource int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateResource, &weightMsgCreateResource, nil,
		func(_ *rand.Rand) {
			weightMsgCreateResource = defaultWeightMsgCreateResource
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateResource,
		warmstoragechainsimulation.SimulateMsgCreateResource(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateResource int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateResource, &weightMsgUpdateResource, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateResource = defaultWeightMsgUpdateResource
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateResource,
		warmstoragechainsimulation.SimulateMsgUpdateResource(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteResource int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteResource, &weightMsgDeleteResource, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteResource = defaultWeightMsgDeleteResource
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteResource,
		warmstoragechainsimulation.SimulateMsgDeleteResource(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
