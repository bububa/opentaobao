package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionProductInfoGet Get Single Product Info
// aliexpress.solution.product.info.get
//
// Get Single Product Info
func AliexpressSolutionProductInfoGet(clt *core.SDKClient, req *aesolution.AliexpressSolutionProductInfoGetAPIRequest, resp *aesolution.AliexpressSolutionProductInfoGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
