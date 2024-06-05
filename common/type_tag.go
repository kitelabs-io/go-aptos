package common

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	TypeTagStructPattern = regexp.MustCompile(`^(.+?)::(.+?)::(.+?)(?:<(.+)>)?$`)
	NullStructName       = "Null"
	TypeParamsSeparator  = ", "
)

type TypeTag interface {
	fmt.Stringer
}

type TypeTagStruct struct {
	ModuleAddress AccountAddress
	ModuleName    string
	StructName    string
	TypeParams    []TypeTag
}

func AsTypeTagStruct(tag string) (TypeTagStruct, bool) {
	matches := TypeTagStructPattern.FindStringSubmatch(tag)

	if len(matches) < 4 {
		return TypeTagStruct{}, false
	}

	moduleAddress, err := HexToAccountAddress(matches[1])
	if err != nil {
		return TypeTagStruct{}, false
	}

	var typeParams []TypeTag
	if len(matches) == 5 {
		typeParamsStr := matches[4]
		countAngleBracket := 0
		startPartIdx := 0

		for idx, ch := range typeParamsStr {
			if ch == '<' {
				countAngleBracket += 1
				continue
			}

			if ch == '>' {
				countAngleBracket -= 1
				// We don't continue in this case since this character
				// might be the last character of typeParamsStr.
			}

			if countAngleBracket != 0 {
				continue
			}

			if idx == len(typeParamsStr)-1 || (ch == ',' && typeParamsStr[idx+1] == ' ') {
				endPartIdx := idx
				if idx == len(typeParamsStr)-1 {
					endPartIdx += 1 // Include the last character in the part
				}

				partStr := typeParamsStr[startPartIdx:endPartIdx]
				typeParam, ok := AsTypeTagStruct(partStr)
				if !ok {
					continue
				}

				typeParams = append(typeParams, typeParam)
				startPartIdx = idx + 2
			}
		}
	}

	return TypeTagStruct{
		ModuleAddress: moduleAddress,
		ModuleName:    matches[2],
		StructName:    matches[3],
		TypeParams:    typeParams,
	}, true
}

func (t TypeTagStruct) String() string {
	structType := fmt.Sprintf("%s::%s::%s", t.ModuleAddress.PrefixZeroTrimmedHex(), t.ModuleName, t.StructName)

	if len(t.TypeParams) == 0 {
		return structType
	}

	typeParamsStr := make([]string, 0, len(t.TypeParams))
	for _, param := range t.TypeParams {
		typeParamsStr = append(typeParamsStr, param.String())
	}

	return fmt.Sprintf("%s<%s>", structType, strings.Join(typeParamsStr, TypeParamsSeparator))
}

func (t TypeTagStruct) IsNull() bool {
	return t.StructName == NullStructName
}
