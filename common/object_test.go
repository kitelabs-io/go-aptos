package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateObjectAddress(t *testing.T) {
	addr, _ := HexToAccountAddress("0xcafe")

	testCases := []struct {
		address  AccountAddress
		seed     string
		expected string
	}{
		{
			address:  addr,
			seed:     "hello_aptos",
			expected: "0xb84eec59931c742ce81e3cf9f1c7cd977c24f8882d6e9e5641d2c624115cc495",
		},
		{
			address:  addr,
			seed:     "0x1::aptos_coin::AptosCoin",
			expected: "0x2b923e259c24fa528ae7e5d3679a80dc5ab6db15896a7a06c9c25bef8a78dbc6",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.seed, func(t *testing.T) {
			objectAddress := CreateObjectAddress(tc.address, tc.seed)
			assert.Equal(t, tc.expected, objectAddress)
		})
	}
}
