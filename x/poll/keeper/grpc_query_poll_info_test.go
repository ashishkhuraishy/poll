package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/ashishkhuraishy/poll/testutil/keeper"
	"github.com/ashishkhuraishy/poll/testutil/nullify"
	"github.com/ashishkhuraishy/poll/x/poll/types"
)

func TestPollInfoQuery(t *testing.T) {
	keeper, ctx := keepertest.PollKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestPollInfo(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetPollInfoRequest
		response *types.QueryGetPollInfoResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetPollInfoRequest{},
			response: &types.QueryGetPollInfoResponse{PollInfo: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.PollInfo(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
