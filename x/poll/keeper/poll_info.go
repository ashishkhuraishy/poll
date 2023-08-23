package keeper

import (
	"github.com/ashishkhuraishy/poll/x/poll/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPollInfo set pollInfo in the store
func (k Keeper) SetPollInfo(ctx sdk.Context, pollInfo types.PollInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollInfoKey))
	b := k.cdc.MustMarshal(&pollInfo)
	store.Set([]byte{0}, b)
}

// GetPollInfo returns pollInfo
func (k Keeper) GetPollInfo(ctx sdk.Context) (val types.PollInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePollInfo removes pollInfo from the store
func (k Keeper) RemovePollInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollInfoKey))
	store.Delete([]byte{0})
}
