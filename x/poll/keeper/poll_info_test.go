package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/ashishkhuraishy/poll/testutil/keeper"
	"github.com/ashishkhuraishy/poll/testutil/nullify"
	"github.com/ashishkhuraishy/poll/x/poll/keeper"
	"github.com/ashishkhuraishy/poll/x/poll/types"
)

func createTestPollInfo(keeper *keeper.Keeper, ctx sdk.Context) types.PollInfo {
	item := types.PollInfo{}
	keeper.SetPollInfo(ctx, item)
	return item
}

func TestPollInfoGet(t *testing.T) {
	keeper, ctx := keepertest.PollKeeper(t)
	item := createTestPollInfo(keeper, ctx)
	rst, found := keeper.GetPollInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestPollInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.PollKeeper(t)
	createTestPollInfo(keeper, ctx)
	keeper.RemovePollInfo(ctx)
	_, found := keeper.GetPollInfo(ctx)
	require.False(t, found)
}
