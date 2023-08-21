package keeper_test

import (
	"testing"

	testkeeper "github.com/ashishkhuraishy/poll/testutil/keeper"
	"github.com/ashishkhuraishy/poll/x/poll/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PollKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
