package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkChannelOrderRefundConfirm 退款确认
// alibaba.wdk.channel.order.refund.confirm
//
// 退款确认
func AlibabaWdkChannelOrderRefundConfirm(clt *core.SDKClient, req *wdk.AlibabaWdkChannelOrderRefundConfirmAPIRequest, resp *wdk.AlibabaWdkChannelOrderRefundConfirmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
