package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyMerchantIdMix 商家用户id混淆
// alibaba.tcls.aelophy.merchant.id.mix
//
// 商家用户id混淆
func AlibabaTclsAelophyMerchantIdMix(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantIdMixAPIRequest, resp *wdk.AlibabaTclsAelophyMerchantIdMixAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
