package types

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type Option[T any] struct {
	Vec []T `json:"vec"`
}

type Object struct {
	Inner string `mapstructure:"inner" json:"inner"`
}

type TypeInfo struct {
	AccountAddress string `mapstructure:"account_address" json:"account_address"`
	ModuleName     string `mapstructure:"module_name" json:"module_name"`
	StructName     string `mapstructure:"struct_name" json:"struct_name"`
}

func (t *TypeInfo) GetTypeName() string {
	return fmt.Sprintf("%s::%s::%s", t.AccountAddress, hexToString(t.ModuleName), hexToString(t.StructName))
}

func hexToString(hexStr string) string {
	h, err := hex.DecodeString(strings.TrimPrefix(hexStr, "0x"))
	if err != nil {
		return ""
	}

	return string(h)
}
