package client

import (
	"context"
	"encoding/json"
	"strconv"
)

var (
	pathView = "/view"
)

func (c *Client) View(ctx context.Context, bodyParams ViewBodyParams, queryParams ViewQueryParams) ([]any, *Metadata, error) {
	body, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, nil, err
	}

	req := c.client.R().
		SetContext(ctx).
		SetBody(body)

	if queryParams.LedgerVersion > 0 {
		req.SetQueryParam("ledger_version", strconv.FormatUint(queryParams.LedgerVersion, 64))
	}

	var results []any
	req.SetResult(&results)
	resp, err := req.Post(pathView)
	if err != nil {
		return nil, nil, err
	}

	if resp.IsError() {
		return nil, nil, handleErrResp(resp.Body())
	}

	return results, handleRspHdr(resp.Header()), nil
}
