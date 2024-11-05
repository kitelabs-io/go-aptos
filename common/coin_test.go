package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePairedFAAddress(t *testing.T) {
	testCases := []struct {
		coinType          string
		expectedFAAddress string
	}{
		{
			coinType:          "0x275f508689de8756169d1ee02d889c777de1cebda3a7bbcce63ba8a27c563c6f::tokens::USDT",
			expectedFAAddress: "0xe161897670a0ee5a0e3c79c3b894a0c46e4ba54c6d2ca32e285ab4b01eb74b66",
		},
		{
			coinType:          "0x275f508689de8756169d1ee02d889c777de1cebda3a7bbcce63ba8a27c563c6f::tokens::USDC",
			expectedFAAddress: "0x1e74c3312b1a7a08eb7cf61310787597ea6609d6d99ce86c0e48399144ea4ce9",
		},
		{
			coinType:          "0x275f508689de8756169d1ee02d889c777de1cebda3a7bbcce63ba8a27c563c6f::tokens::WETH",
			expectedFAAddress: "0xa484a866e1bfcb76e8057939d6944539070b53c511813d7b21c76cae9e6a6e26",
		},
		{
			coinType:          "0xc7efb4076dbe143cbcd98cfaaa929ecfc8f299203dfff63b95ccb6bfe19850fa::swap::LPToken<0x1::aptos_coin::AptosCoin, 0x3c27315fb69ba6e4b960f1507d1cefcc9a4247869f26a8d59d6b7869d23782c::test_coins::CAKE>",
			expectedFAAddress: "0x626eb18033951b49e19127feaf13af0ce7c5e8c761b1f12d152fe70ea0ef10fe",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.coinType, func(t *testing.T) {
			faAddress := CreatePairedFAAddress(tc.coinType)
			assert.Equal(t, tc.expectedFAAddress, faAddress)
		})
	}
}
