package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionOrderGet get order list
// aliexpress.solution.order.get
//
// Get Order List from AliExpress
func AliexpressSolutionOrderGet(clt *core.SDKClient, req *aesolution.AliexpressSolutionOrderGetAPIRequest, resp *aesolution.AliexpressSolutionOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
