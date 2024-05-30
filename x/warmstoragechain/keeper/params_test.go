package keeper_test

import (
	"testing"

	testkeeper "github.com/BlockCraftsman/WarmStorageChain/testutil/keeper"
	"github.com/BlockCraftsman/WarmStorageChain/x/warmstoragechain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.WarmstoragechainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
