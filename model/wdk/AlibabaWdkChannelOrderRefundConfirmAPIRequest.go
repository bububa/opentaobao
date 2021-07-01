package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkChannelOrderRefundConfirmAPIRequest
退款确认 API请求
alibaba.wdk.channel.order.refund.confirm

退款确认 */
type AlibabaWdkChannelOrderRefundConfirmAPIRequest struct {
	model.Params
	// 退款确认信息
	_orderRefundConfirmInfo *OrderRefundConfirmInfo
}

// New
