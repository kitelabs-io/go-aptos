package common

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/sha3"
)

var (
	objectFromSeedAddressScheme = []byte{0xfe}
)

func CreateObjectAddress(
	account AccountAddress,
	seed string,
) string {
	// sha3_256(address + seed + objectFromSeedAddressScheme)

	input := []byte{}
	input = append(input, account[:]...)
	input = append(input, []byte(seed)...)
	input = append(input, objectFromSeedAddressScheme...)

	hash := sha3.New256()
	_, _ = hash.Write([]byte(input))

	return fmt.Sprintf("0x%s", hex.EncodeToString(hash.Sum(nil)))
}
