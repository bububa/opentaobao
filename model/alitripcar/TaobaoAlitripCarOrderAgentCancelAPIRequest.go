package alitripcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripcarorderagentcancelAPIRequest 司机或客服取消订单 API请求
// taobao.alitrip.car.order.agent.cancel
//
// 司机或客服取消订单后通知飞猪订单取消
type TaobaoalitripcarorderagentcancelAPIRequest struct {
	model.Params
	// 取消对象
	_paramOrderCancel *OrderCancel
}

// NewTaobaoalitripcarorderagentcancelRequest 初始化TaobaoalitripcarorderagentcancelAPIRequest对象
func NewTaobaoalitripcarorderagentcancelRequest() *TaobaoalitripcarorderagentcancelAPIRequest {
	return &TaobaoalitripcarorderagentcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripcarorderagentcancelAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.order.agent.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripcarorderagentcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripcarorderagentcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderCancel is ParamOrderCancel Setter
// 取消对象
func (r *TaobaoalitripcarorderagentcancelAPIRequest) SetParamOrderCancel(_paramOrderCancel *OrderCancel) error {
	r._paramOrderCancel = _paramOrderCancel
	r.Set("param_order_cancel", _paramOrderCancel)
	return nil
}

// GetParamOrderCancel ParamOrderCancel Getter
func (r TaobaoalitripcarorderagentcancelAPIRequest) GetParamOrderCancel() *OrderCancel {
	return r._paramOrderCancel
}
