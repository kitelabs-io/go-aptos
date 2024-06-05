package client

import "github.com/go-resty/resty/v2"

type Client struct {
	client *resty.Client
}

func NewClient(options ...Option) *Client {
	c := &Client{}

	for _, opt := range options {
		opt(c)
	}

	return c
}
