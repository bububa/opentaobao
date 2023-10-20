package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyMerchantChannelOrderUpdatestatus 翱象商家自有渠道 订单状态更新
// alibaba.tcls.aelophy.merchant.channel.order.updatestatus
//
// 订单状态变更
func AlibabaTclsAelophyMerchantChannelOrderUpdatestatus(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantChannelOrderUpdatestatusAPIRequest, resp *wdk.AlibabaTclsAelophyMerchantChannelOrderUpdatestatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
