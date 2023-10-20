package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyMerchantChannelRefundCancel 翱象商家自有渠道 逆向单申请取消
// alibaba.tcls.aelophy.merchant.channel.refund.cancel
//
// 翱象小程序 用户逆向申请取消
func AlibabaTclsAelophyMerchantChannelRefundCancel(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest, resp *wdk.AlibabaTclsAelophyMerchantChannelRefundCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
