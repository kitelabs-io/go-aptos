package types

import "encoding/json"

type Resource struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}
