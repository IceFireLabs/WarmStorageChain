package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/BlockCraftsman/WarmStorageChain/testutil/keeper"
	"github.com/BlockCraftsman/WarmStorageChain/testutil/nullify"
	"github.com/BlockCraftsman/WarmStorageChain/x/warmstoragechain/keeper"
	"github.com/BlockCraftsman/WarmStorageChain/x/warmstoragechain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNResource(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Resource {
	items := make([]types.Resource, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetResource(ctx, items[i])
	}
	return items
}

func TestResourceGet(t *testing.T) {
	keeper, ctx := keepertest.WarmstoragechainKeeper(t)
	items := createNResource(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetResource(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestResourceRemove(t *testing.T) {
	keeper, ctx := keepertest.WarmstoragechainKeeper(t)
	items := createNResource(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveResource(ctx,
			item.Index,
		)
		_, found := keeper.GetResource(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestResourceGetAll(t *testing.T) {
	keeper, ctx := keepertest.WarmstoragechainKeeper(t)
	items := createNResource(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllResource(ctx)),
	)
}
