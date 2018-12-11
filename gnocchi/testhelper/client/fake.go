package client

import (
	"github.com/samuelbernardolip/gophercloud"
	"github.com/samuelbernardolip/gophercloud/testhelper/client"
)

// TokenID is a fake Identity service token.
const TokenID = client.TokenID

// ServiceClient returns a generic service client for use in tests.
func ServiceClient() *gophercloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v1/"
	return sc
}
