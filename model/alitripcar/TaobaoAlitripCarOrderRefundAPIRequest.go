package alitripcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripcarorderrefundAPIRequest 用户投诉达成一致后给用户退款 API请求
// taobao.alitrip.car.order.refund
//
// 用户投诉后，供应商客服与客户沟通达成一致后通知飞猪给客户退款。退款金额以接口回调金额为准。
type TaobaoalitripcarorderrefundAPIRequest struct {
	model.Params
	// 退款对象
	_paramOrderRefund *OrderRefund
}

// NewTaobaoalitripcarorderrefundRequest 初始化TaobaoalitripcarorderrefundAPIRequest对象
func NewTaobaoalitripcarorderrefundRequest() *TaobaoalitripcarorderrefundAPIRequest {
	return &TaobaoalitripcarorderrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripcarorderrefundAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.order.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripcarorderrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripcarorderrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderRefund is ParamOrderRefund Setter
// 退款对象
func (r *TaobaoalitripcarorderrefundAPIRequest) SetParamOrderRefund(_paramOrderRefund *OrderRefund) error {
	r._paramOrderRefund = _paramOrderRefund
	r.Set("param_order_refund", _paramOrderRefund)
	return nil
}

// GetParamOrderRefund ParamOrderRefund Getter
func (r TaobaoalitripcarorderrefundAPIRequest) GetParamOrderRefund() *OrderRefund {
	return r._paramOrderRefund
}
