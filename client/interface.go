package client

import (
	"context"

	"github.com/kitelabs-io/go-aptos/types"
)

type IClient interface {
	GetNodeInfo(ctx context.Context) (*NodeInfo, error)

	GetAccountResources(
		ctx context.Context,
		address string,
		queryParams GetAccountResourcesQueryParams,
	) ([]types.Resource, *Metadata, error)

	GetAccountResource(
		ctx context.Context,
		address string,
		resourceType string,
		queryParams GetAccountResourceQueryParams,
	) (*types.Resource, *Metadata, error)

	View(
		ctx context.Context,
		bodyParams ViewBodyParams,
		queryParams ViewQueryParams,
	) ([]any, *Metadata, error)
}
