package keeper

import (
	"github.com/BlockCraftsman/WarmStorageChain/x/warmstoragechain/types"
)

var _ types.QueryServer = Keeper{}
