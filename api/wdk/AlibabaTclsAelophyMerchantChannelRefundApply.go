package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyMerchantChannelRefundApply 翱象商家自有渠道 逆向单申请
// alibaba.tcls.aelophy.merchant.channel.refund.apply
//
// 翱象小程序 用户逆向单申请
func AlibabaTclsAelophyMerchantChannelRefundApply(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest, resp *wdk.AlibabaTclsAelophyMerchantChannelRefundApplyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
