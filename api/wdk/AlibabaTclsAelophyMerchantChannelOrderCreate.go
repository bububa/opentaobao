package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyMerchantChannelOrderCreate 翱象商家自有渠道 订单创建
// alibaba.tcls.aelophy.merchant.channel.order.create
//
// 翱象小程序渠道订单创建
func AlibabaTclsAelophyMerchantChannelOrderCreate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest, resp *wdk.AlibabaTclsAelophyMerchantChannelOrderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
