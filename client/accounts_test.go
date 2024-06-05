package client

import (
	"context"

	"github.com/goccy/go-json"
)

var (
	addressPancake = "0xc7efb4076dbe143cbcd98cfaaa929ecfc8f299203dfff63b95ccb6bfe19850fa"
)

func (ts *ClientTestSuite) TestGetAccountResources() {
	ctx := context.Background()

	resources, metadata, err := ts.client.GetAccountResources(
		ctx,
		addressPancake,
		GetAccountResourcesQueryParams{Limit: 3},
	)

	ts.Assert().NoError(err)
	ts.Assert().NotNil(metadata)
	ts.Assert().NotNil(resources)
	ts.Assert().Equal(3, len(resources))
}

func (ts *ClientTestSuite) TestGetAccountResource() {
	ctx := context.Background()

	resourceType := `0xc7efb4076dbe143cbcd98cfaaa929ecfc8f299203dfff63b95ccb6bfe19850fa::swap::TokenPairReserve<0xa72a97e872be9ee3d2f14d56fd511eb7e4a53f4055be3a267d8602e7685b41c0::coin::T, 0xced3e98a47279b4d39a75fa8da867e2e8383d5d8ce7e59b2964a9469616a761b::coin::T>`
	resource, metadata, err := ts.client.GetAccountResource(
		ctx,
		addressPancake,
		resourceType,
		GetAccountResourceQueryParams{},
	)

	ts.Assert().NoError(err)
	ts.Assert().NotNil(metadata)
	ts.Assert().NotNil(resource)
	ts.Assert().Equal(resourceType, resource.Type)

	var data struct {
		ReserveX string `json:"reserve_x"`
		ReserveY string `json:"reserve_y"`
	}
	if err = json.Unmarshal(resource.Data, &data); err != nil {
		ts.Assert().NoError(err)
	}

	ts.Assert().True(len(data.ReserveX) > 0)
	ts.Assert().True(len(data.ReserveY) > 0)
}
