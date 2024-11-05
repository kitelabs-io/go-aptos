package aptos

import (
	"context"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/kitelabs-io/go-aptos/client"
	"github.com/stretchr/testify/assert"
)

func TestCoinClient_GetPairedFA(t *testing.T) {
	coinClient := newCoinClient()

	faAddress, err := coinClient.GetPairedFA(context.Background(), "0x1::aptos_coin::AptosCoin")
	assert.NoError(t, err)
	assert.Equal(t, "0xa", faAddress)
}

func TestCoinClient_GetPairedCoin(t *testing.T) {
	coinClient := newCoinClient()

	// Fungible asset with paired coin
	coinType, err := coinClient.GetPairedCoin(context.Background(), "0xa")
	assert.NoError(t, err)
	assert.Equal(t, "0x1::aptos_coin::AptosCoin", coinType)

	coinType, err = coinClient.GetPairedCoin(context.Background(), "0x1f783324e5de19165863de35c3c432528e2007d5336058d22a34ca28397a8943")
	assert.NoError(t, err)
	assert.Equal(t, "0x43417434fd869edee76cca2a4d2301e528a1551b1d719b75c350c3c97d15b8b9::coins::USDT", coinType)

	// Fungible asset without paired coin
	coinType, err = coinClient.GetPairedCoin(context.Background(), "0x0e550aeb585046fb2b63633fe8834bd199cbf3756b5cf2a8d693883d67f2eb3e")
	assert.NoError(t, err)
	assert.Equal(t, "", coinType)
}

func newCoinClient() ICoinClient {
	restyClient := resty.New()
	restyClient.SetBaseURL("https://api.testnet.aptoslabs.com/v1")
	restyClient.SetHeader("Accept", "application/json")
	restyClient.SetHeader("Content-Type", "application/json")
	nodeClient := client.NewClient(client.WithClient(restyClient))

	return NewCoinClient(nodeClient)
}
