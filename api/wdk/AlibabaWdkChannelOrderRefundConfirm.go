package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkChannelOrderRefundConfirm 退款确认
// alibaba.wdk.channel.order.refund.confirm
//
// 退款确认
func AlibabaWdkChannelOrderRefundConfirm(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkChannelOrderRefundConfirmAPIRequest, resp *wdk.AlibabaWdkChannelOrderRefundConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
