package client

import "github.com/go-resty/resty/v2"

type Option func(client *Client)

func WithClient(restyClient *resty.Client) Option {
	return func(c *Client) {
		c.client = restyClient
	}
}
