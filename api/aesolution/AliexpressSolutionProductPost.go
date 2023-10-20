package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionProductPost Product posting API
// aliexpress.solution.product.post
//
// Product posting API for Oversea merchants, simplifying the complexity of integration that sellers and merchants face. For example, these sellers can use their own category and attributes instead of mapping those from AE.
func AliexpressSolutionProductPost(clt *core.SDKClient, req *aesolution.AliexpressSolutionProductPostAPIRequest, resp *aesolution.AliexpressSolutionProductPostAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
