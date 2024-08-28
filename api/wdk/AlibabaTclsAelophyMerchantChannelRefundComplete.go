package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyMerchantChannelRefundComplete 翱象商家自有渠道 逆向单完成
// alibaba.tcls.aelophy.merchant.channel.refund.complete
//
// 翱象小程序 退款完成
func AlibabaTclsAelophyMerchantChannelRefundComplete(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest, resp *wdk.AlibabaTclsAelophyMerchantChannelRefundCompleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
