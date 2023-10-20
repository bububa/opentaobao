package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionMerchantProfileGet aliexpress.solution.merchant.profile.get
// aliexpress.solution.merchant.profile.get
//
// API for oversea sellers to obtain the normal information, e.g. store id, registration country code.
func AliexpressSolutionMerchantProfileGet(clt *core.SDKClient, req *aesolution.AliexpressSolutionMerchantProfileGetAPIRequest, resp *aesolution.AliexpressSolutionMerchantProfileGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
