package client

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

type ClientTestSuite struct {
	suite.Suite

	client *Client
}

func (ts *ClientTestSuite) SetupTest() {
	restyClient := resty.New()
	restyClient.
		SetBaseURL("https://api.mainnet.aptoslabs.com/v1").
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json")

	ts.client = NewClient(WithClient(restyClient))
}
