package poll

import (
	"github.com/ashishkhuraishy/poll/x/poll/keeper"
	"github.com/ashishkhuraishy/poll/x/poll/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.PollInfo != nil {
		k.SetPollInfo(ctx, *genState.PollInfo)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all pollInfo
	pollInfo, found := k.GetPollInfo(ctx)
	if found {
		genesis.PollInfo = &pollInfo
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
