package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/ashishkhuraishy/poll/testutil/keeper"
	"github.com/ashishkhuraishy/poll/x/poll/keeper"
	"github.com/ashishkhuraishy/poll/x/poll/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.PollKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
