package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyMerchantChannelOrderPrecheck 前置校验商品是否可下单作业
// alibaba.tcls.aelophy.merchant.channel.order.precheck
//
// 前置校验商品是否可下单作业
func AlibabaTclsAelophyMerchantChannelOrderPrecheck(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest, session string) (*wdk.AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponse, error) {
	var resp wdk.AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
