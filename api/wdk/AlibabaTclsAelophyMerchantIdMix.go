package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophymerchantidmix 商家用户id混淆
// alibaba.tcls.aelophy.merchant.id.mix
//
// 商家用户id混淆
func Alibabatclsaelophymerchantidmix(clt *core.SDKClient, req *wdk.AlibabatclsaelophymerchantidmixAPIRequest, session string) (*wdk.AlibabatclsaelophymerchantidmixAPIResponse, error) {
	var resp wdk.AlibabatclsaelophymerchantidmixAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
