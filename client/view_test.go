package client

import (
	"context"

	"github.com/stretchr/testify/assert"
)

func (ts *ClientTestSuite) TestView() {
	ts.Run("it should view correctly without arguments", func() {
		ctx := context.Background()

		results, metadata, err := ts.client.View(
			ctx,
			ViewBodyParams{
				Function:      "0x4bf51972879e3b95c4781a5cdcb9e1ee24ef483e7d22f2d903626f126df62bd1::liquidity_pool::all_pool_addresses",
				TypeArguments: []string{},
				Arguments:     []any{},
			},
			ViewQueryParams{},
		)

		ts.Assert().NoError(err)
		ts.Assert().NotNil(metadata)
		ts.Assert().NotNil(results)

		result0, ok := results[0].([]any)
		assert.True(ts.T(), ok)

		poolIds := make([]string, 0, len(result0))
		for _, obj := range result0 {
			objMap, ok := obj.(map[string]any)
			if !ok {
				continue
			}

			poolID, ok := objMap["inner"]
			if !ok {
				continue
			}

			poolIds = append(poolIds, poolID.(string))
		}

		assert.GreaterOrEqual(ts.T(), len(poolIds), 15)
	})

	ts.Run("it should view correctly with arguments - 1", func() {
		ctx := context.Background()

		results, metadata, err := ts.client.View(
			ctx,
			ViewBodyParams{
				Function:      "0x4bf51972879e3b95c4781a5cdcb9e1ee24ef483e7d22f2d903626f126df62bd1::liquidity_pool::pool_reserves",
				TypeArguments: []string{"0x1::object::ObjectCore"},
				Arguments:     []any{"0x85d3337c4ca94612f278c5164d2b21d0d83354648bf9555272b5f9d8f1f33b2a"},
			},
			ViewQueryParams{},
		)

		ts.Assert().NoError(err)
		ts.Assert().NotNil(metadata)
		ts.Assert().Equal(2, len(results))
	})

	ts.Run("it should view correctly with arguments - 2", func() {
		ctx := context.Background()

		results, metadata, err := ts.client.View(
			ctx,
			ViewBodyParams{
				Function:      "0x4bf51972879e3b95c4781a5cdcb9e1ee24ef483e7d22f2d903626f126df62bd1::liquidity_pool::liquidity_pool_address",
				TypeArguments: []string{},
				Arguments: []any{
					"0x2ebb2ccac5e027a87fa0e2e5f656a3a4238d6a48d93ec9b610d570fc0aa0df12",
					"0xedc2704f2cef417a06d1756a04a16a9fa6faaed13af469be9cdfcac5a21a8e2e",
					false,
				},
			},
			ViewQueryParams{},
		)

		ts.Assert().NoError(err)
		ts.Assert().NotNil(metadata)
		ts.Assert().Equal(1, len(results))
	})

	ts.Run("it should view correctly with arguments - 3", func() {
		ctx := context.Background()

		results, metadata, err := ts.client.View(
			ctx,
			ViewBodyParams{
				Function:      "0x4bf51972879e3b95c4781a5cdcb9e1ee24ef483e7d22f2d903626f126df62bd1::liquidity_pool::get_amount_out",
				TypeArguments: []string{},
				Arguments: []any{
					"0x85d3337c4ca94612f278c5164d2b21d0d83354648bf9555272b5f9d8f1f33b2a",
					"0x2ebb2ccac5e027a87fa0e2e5f656a3a4238d6a48d93ec9b610d570fc0aa0df12",
					"100000000",
				},
			},
			ViewQueryParams{},
		)

		ts.Assert().NoError(err)
		ts.Assert().NotNil(metadata)
		ts.Assert().Equal(2, len(results))
	})
}
