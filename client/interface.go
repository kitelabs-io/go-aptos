package client

import (
	"context"

	"github.com/anqa-ag/go-aptos/types"
)

type IClient interface {
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
