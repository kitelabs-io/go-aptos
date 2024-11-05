package common

var (
	aptosFungibleAsset, _ = HexToAccountAddress("0xa")
)

// CreatePairedFAAddress creates a paired fungible asset address from a coin type.
func CreatePairedFAAddress(coinType string) string {
	return CreateObjectAddress(aptosFungibleAsset, coinType)
}
