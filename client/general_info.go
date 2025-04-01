package client

import (
	"context"
)

func (c *Client) GetNodeInfo(ctx context.Context) (*NodeInfo, error) {
	var nodeInfo NodeInfo
	req := c.client.R().SetContext(ctx).SetResult(&nodeInfo)

	resp, err := req.Get("/")
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, handleErrResp(resp.Body())
	}

	return &nodeInfo, nil
}
