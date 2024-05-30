package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/BlockCraftsman/WarmStorageChain/testutil/keeper"
	"github.com/BlockCraftsman/WarmStorageChain/x/warmstoragechain/keeper"
	"github.com/BlockCraftsman/WarmStorageChain/x/warmstoragechain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.WarmstoragechainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
