package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyMerchantChannelOrderPrecheck 前置校验商品是否可下单作业
// alibaba.tcls.aelophy.merchant.channel.order.precheck
//
// 前置校验商品是否可下单作业
func AlibabaTclsAelophyMerchantChannelOrderPrecheck(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest, resp *wdk.AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
