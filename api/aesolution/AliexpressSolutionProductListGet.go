package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionProductListGet Get product list
// aliexpress.solution.product.list.get
//
// Get product list
func AliexpressSolutionProductListGet(clt *core.SDKClient, req *aesolution.AliexpressSolutionProductListGetAPIRequest, resp *aesolution.AliexpressSolutionProductListGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
