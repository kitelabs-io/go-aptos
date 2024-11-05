package coin

import (
	"context"
	"fmt"

	node "github.com/kitelabs-io/go-aptos/client"
	"github.com/kitelabs-io/go-aptos/types"
	"github.com/mitchellh/mapstructure"
)

const (
	coinModuleName = "coin"
)

type IClient interface {
	GetPairedCoin(ctx context.Context, faAddress string) (string, error)
	GetPairedFA(ctx context.Context, coinType string) (string, error)
}

type client struct {
	nodeClient node.IClient
}

func NewClient(nodeClient node.IClient) IClient {
	return &client{
		nodeClient: nodeClient,
	}
}

// GetPairedFA returns the paired fungible asset address for a given coin type.
func (c *client) GetPairedFA(ctx context.Context, coinType string) (string, error) {
	results, _, err := c.nodeClient.View(
		ctx,
		node.ViewBodyParams{
			Function:      c.getCoinModuleFunction("paired_metadata"),
			TypeArguments: []string{coinType},
			Arguments:     []string{},
		},
		node.ViewQueryParams{},
	)
	if err != nil {
		return "", err
	}

	if len(results) != 1 {
		return "", nil
	}

	metadataOption := types.Option[types.Object]{}
	mapstructure.Decode(results[0], &metadataOption)

	if len(metadataOption.Vec) == 0 {
		return "", nil
	}

	return metadataOption.Vec[0].Inner, nil
}

// GetPairedCoin returns the paired coin type with the given fungible asset address.
func (c *client) GetPairedCoin(ctx context.Context, faAddress string) (string, error) {
	results, _, err := c.nodeClient.View(
		ctx,
		node.ViewBodyParams{
			Function:      c.getCoinModuleFunction("paired_coin"),
			TypeArguments: []string{},
			Arguments:     []string{faAddress},
		},
		node.ViewQueryParams{},
	)
	if err != nil {
		return "", err
	}

	if len(results) != 1 {
		return "", nil
	}

	metadataOption := types.Option[types.TypeInfo]{}
	mapstructure.Decode(results[0], &metadataOption)

	if len(metadataOption.Vec) == 0 {
		return "", nil
	}

	return metadataOption.Vec[0].GetTypeName(), nil
}

func (r *client) getCoinModuleFunction(functionName string) string {
	return fmt.Sprintf("0x1::%s::%s", coinModuleName, functionName)
}
