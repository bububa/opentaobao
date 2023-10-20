package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophymerchantchannelorderprecheck 前置校验商品是否可下单作业
// alibaba.tcls.aelophy.merchant.channel.order.precheck
//
// 前置校验商品是否可下单作业
func Alibabatclsaelophymerchantchannelorderprecheck(clt *core.SDKClient, req *wdk.AlibabatclsaelophymerchantchannelorderprecheckAPIRequest, session string) (*wdk.AlibabatclsaelophymerchantchannelorderprecheckAPIResponse, error) {
	var resp wdk.AlibabatclsaelophymerchantchannelorderprecheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
