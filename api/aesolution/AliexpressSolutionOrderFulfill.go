package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionOrderFulfill fulfill order
// aliexpress.solution.order.fulfill
//
// fulfill order for seller
func AliexpressSolutionOrderFulfill(clt *core.SDKClient, req *aesolution.AliexpressSolutionOrderFulfillAPIRequest, resp *aesolution.AliexpressSolutionOrderFulfillAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
