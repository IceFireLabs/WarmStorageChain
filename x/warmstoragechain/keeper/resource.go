package keeper

import (
	"github.com/BlockCraftsman/WarmStorageChain/x/warmstoragechain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetResource set a specific resource in the store from its index
func (k Keeper) SetResource(ctx sdk.Context, resource types.Resource) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResourceKeyPrefix))
	b := k.cdc.MustMarshal(&resource)
	store.Set(types.ResourceKey(
		resource.Index,
	), b)
}

// GetResource returns a resource from its index
func (k Keeper) GetResource(
	ctx sdk.Context,
	index string,

) (val types.Resource, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResourceKeyPrefix))

	b := store.Get(types.ResourceKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveResource removes a resource from the store
func (k Keeper) RemoveResource(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResourceKeyPrefix))
	store.Delete(types.ResourceKey(
		index,
	))
}

// GetAllResource returns all resource
func (k Keeper) GetAllResource(ctx sdk.Context) (list []types.Resource) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResourceKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Resource
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
