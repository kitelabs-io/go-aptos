package client

// NodeInfo information retrieved about the current state of the blockchain on API requests
type NodeInfo struct {
	ChainId             uint8  `json:"chain_id"`
	Epoch               string `json:"epoch"`
	LedgerTimestamp     string `json:"ledger_timestamp"`
	LedgerVersion       string `json:"ledger_version"`
	OldestLedgerVersion string `json:"oldest_ledger_version"`
	NodeRole            string `json:"node_role"`
	BlockHeight         string `json:"block_height"`
	OldestBlockHeight   string `json:"oldest_block_height"`
	GitHash             string `json:"git_hash"`
}

type GetAccountResourcesQueryParams struct {
	LedgerVersion uint64
	Limit         int
	Start         string
}

type GetAccountResourceQueryParams struct {
	LedgerVersion uint64
}

type ViewQueryParams struct {
	LedgerVersion uint64
}

type ViewBodyParams struct {
	Function      string   `json:"function"`
	TypeArguments []string `json:"type_arguments"`
	Arguments     []any    `json:"arguments"`
}

type ErrorResponse struct {
	Message     string `json:"message"`
	ErrorCode   string `json:"error_code"`
	VMErrorCode int    `json:"vm_error_code"`
}

type Metadata struct {
	BlockHeight         int
	ChainID             int
	EPoch               int
	LedgerOldestVersion int
	LedgerTimestampUSec int
	LedgerVersion       int
	OldestBlockHeight   int
	Cursor              string
}
