package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/BlockCraftsman/WarmStorageChain/testutil/keeper"
	"github.com/BlockCraftsman/WarmStorageChain/testutil/nullify"
	"github.com/BlockCraftsman/WarmStorageChain/x/warmstoragechain/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestResourceQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.WarmstoragechainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNResource(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetResourceRequest
		response *types.QueryGetResourceResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetResourceRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetResourceResponse{Resource: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetResourceRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetResourceResponse{Resource: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetResourceRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Resource(wctx, tc.request)
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

func TestResourceQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.WarmstoragechainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNResource(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllResourceRequest {
		return &types.QueryAllResourceRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ResourceAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Resource), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Resource),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ResourceAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Resource), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Resource),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ResourceAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Resource),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ResourceAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
