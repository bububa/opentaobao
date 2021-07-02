package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyMerchantIdMix 商家用户id混淆
// alibaba.tcls.aelophy.merchant.id.mix
//
// 商家用户id混淆
func AlibabaTclsAelophyMerchantIdMix(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantIdMixAPIRequest, session string) (*wdk.AlibabaTclsAelophyMerchantIdMixAPIResponse, error) {
	var resp wdk.AlibabaTclsAelophyMerchantIdMixAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
