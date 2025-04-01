package client

import (
	"context"
)

func (ts *ClientTestSuite) TestGetNodeInfo() {
	ctx := context.Background()

	nodeInfo, err := ts.client.GetNodeInfo(ctx)

	ts.Assert().NoError(err)
	ts.Assert().NotNil(nodeInfo)
	ts.Assert().NotEmpty(nodeInfo.LedgerVersion)
	ts.Assert().NotEmpty(nodeInfo.LedgerTimestamp)
	ts.Assert().NotEmpty(nodeInfo.BlockHeight)
}
