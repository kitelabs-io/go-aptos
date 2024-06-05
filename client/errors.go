package client

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var (
	ErrUnknownError = errors.New("unknown error")
	ErrRPCError     = errors.New("rpc error")
)

func handleErrResp(data []byte) error {
	var resp ErrorResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return errors.Wrap(ErrUnknownError, string(data))
	}

	return errors.Wrap(ErrRPCError, resp.ErrorCode)
}
