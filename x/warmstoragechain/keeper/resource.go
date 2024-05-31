package keeper

import (
	"log"
	"os"

	"github.com/BlockCraftsman/WarmStorageChain/pkg/ipfs"
	"github.com/BlockCraftsman/WarmStorageChain/x/warmstoragechain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetResource set a specific resource in the store from its index
func (k Keeper) SetResource(ctx sdk.Context, resource types.Resource) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResourceKeyPrefix))

	// Filecoin CID retrieval
	CID := resource.Cid
	// dataS, err := ipfs.GetDATAFromIPFSCID(ipfs.GetReuseIPFSClient(), CID)

	// if err != nil {
	// 	log.Println("SetResource Error", err)
	// 	return
	// }

	// //Just store the first CID file
	// if len(dataS) < 1 {
	// 	log.Println(errors.New("data of CID is nil"))
	// 	return
	// }

	// CIDdata := dataS[0]

	// err = KVDB.Put([]byte(resource.Index), CIDdata, nil)

	// if err != nil {
	// 	log.Println("SetResource NoSQLDB.Set Fail ", err)
	// 	return
	// }
	//MemKV.Set(resource.Index, CIDdata, cache.NoExpiration)
	CIDExtraOutPutFullDir := CIDExtraOutPutDir + "/" + resource.Index

	err := os.MkdirAll(CIDExtraOutPutFullDir, 0755)
	if err != nil {
		log.Println("SetResource Fail", "Key", resource.Index, err)
		return
	}

	CIDExtraOutPutFullPath := CIDExtraOutPutDir + "/" + resource.Index + "/" + CID
	_, err = ipfs.ExtraCIDToDir(ipfs.GetReuseIPFSClient(), CID, CIDExtraOutPutFullPath)
	if err != nil {
		log.Println("SetResource Fail", "Key", resource.Index, "stroage Path", CIDExtraOutPutFullPath, err)
		return
	}

	log.Println("SetResource Success", "Key", resource.Index, "stroage Path", CIDExtraOutPutFullPath)

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

	log.Println("342134234")
	for ; iterator.Valid(); iterator.Next() {
		var val types.Resource
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
