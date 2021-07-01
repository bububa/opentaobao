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

// NewTaobaoAlitripCarOrderRefundRequest 初始化TaobaoAlitripCarOrderRefundAPIRequest对象
func NewTaobaoAlitripCarOrderRefundRequest() *TaobaoAlitripCarOrderRefundAPIRequest {
	return &TaobaoAlitripCarOrderRefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarOrderRefundAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.order.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarOrderRefundAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamOrderRefund Setter
// 退款对象
func (r *TaobaoAlitripCarOrderRefundAPIRequest) SetParamOrderRefund(_paramOrderRefund *OrderRefund) error {
	r._paramOrderRefund = _paramOrderRefund
	r.Set("param_order_refund", _paramOrderRefund)
	return nil
}

// Get ParamOrderRefund Getter
func (r TaobaoAlitripCarOrderRefundAPIRequest) GetParamOrderRefund() *OrderRefund {
	return r._paramOrderRefund
}
