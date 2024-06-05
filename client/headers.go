package client

import (
	"net/http"
	"strconv"
)

var (
	hdrXAptosBlockHeight         = http.CanonicalHeaderKey("X-APTOS-BLOCK-HEIGHT")
	hdrXAptosChainID             = http.CanonicalHeaderKey("X-APTOS-CHAIN-ID")
	hdrXAptosEPoch               = http.CanonicalHeaderKey("X-APTOS-EPOCH")
	hdrXAptosLedgerOldestVersion = http.CanonicalHeaderKey("X-APTOS-LEDGER-OLDEST-VERSION")
	hdrXAptosLedgerTimestampUSec = http.CanonicalHeaderKey("X-APTOS-LEDGER-TIMESTAMPUSEC")
	hdrXAptosLedgerVersion       = http.CanonicalHeaderKey("X-APTOS-LEDGER-VERSION")
	hdrXAptosOldestBlockHeight   = http.CanonicalHeaderKey("X-APTOS-OLDEST-BLOCK-HEIGHT")
	hdrXAptosCursor              = http.CanonicalHeaderKey("X-APTOS-CURSOR")
)

func handleRspHdr(headers http.Header) *Metadata {
	metadata := &Metadata{}

	for key, values := range headers {
		switch key {
		case hdrXAptosBlockHeight:
			metadata.BlockHeight, _ = strconv.Atoi(values[0])
		case hdrXAptosChainID:
			metadata.ChainID, _ = strconv.Atoi(values[0])
		case hdrXAptosEPoch:
			metadata.EPoch, _ = strconv.Atoi(values[0])
		case hdrXAptosLedgerOldestVersion:
			metadata.LedgerOldestVersion, _ = strconv.Atoi(values[0])
		case hdrXAptosLedgerTimestampUSec:
			metadata.LedgerTimestampUSec, _ = strconv.Atoi(values[0])
		case hdrXAptosLedgerVersion:
			metadata.LedgerTimestampUSec, _ = strconv.Atoi(values[0])
		case hdrXAptosOldestBlockHeight:
			metadata.LedgerTimestampUSec, _ = strconv.Atoi(values[0])
		case hdrXAptosCursor:
			metadata.Cursor = values[0]
		}
	}

	return metadata
}
