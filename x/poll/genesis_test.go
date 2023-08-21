package poll_test

import (
	"testing"

	keepertest "github.com/ashishkhuraishy/poll/testutil/keeper"
	"github.com/ashishkhuraishy/poll/testutil/nullify"
	"github.com/ashishkhuraishy/poll/x/poll"
	"github.com/ashishkhuraishy/poll/x/poll/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PollKeeper(t)
	poll.InitGenesis(ctx, *k, genesisState)
	got := poll.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
