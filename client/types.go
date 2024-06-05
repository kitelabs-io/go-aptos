package client

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
	Arguments     []string `json:"arguments"`
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
