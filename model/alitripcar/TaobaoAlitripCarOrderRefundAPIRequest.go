package alitripcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripCarOrderRefundAPIRequest
用户投诉达成一致后给用户退款 API请求
taobao.alitrip.car.order.refund

用户投诉后，供应商客服与客户沟通达成一致后通知飞猪给客户退款。退款金额以接口回调金额为准。 */
type TaobaoAlitripCarOrderRefundAPIRequest struct {
	model.Params
	// 退款对象
	_paramOrderRefund *OrderRefund
}

// New
