package keeper

import (
	"context"

	"github.com/ashishkhuraishy/poll/x/poll/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PollInfo(c context.Context, req *types.QueryGetPollInfoRequest) (*types.QueryGetPollInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPollInfo(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPollInfoResponse{PollInfo: val}, nil
}
