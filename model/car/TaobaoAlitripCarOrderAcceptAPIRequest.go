package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripcarorderacceptAPIRequest 确认接单 API请求
// taobao.alitrip.car.order.accept
//
// 用来接收服务商确认接单信息
type TaobaoalitripcarorderacceptAPIRequest struct {
	model.Params
	// 确认订单请求
	_paramOrderAccept *OrderAccept
}

// NewTaobaoalitripcarorderacceptRequest 初始化TaobaoalitripcarorderacceptAPIRequest对象
func NewTaobaoalitripcarorderacceptRequest() *TaobaoalitripcarorderacceptAPIRequest {
	return &TaobaoalitripcarorderacceptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripcarorderacceptAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.order.accept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripcarorderacceptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripcarorderacceptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderAccept is ParamOrderAccept Setter
// 确认订单请求
func (r *TaobaoalitripcarorderacceptAPIRequest) SetParamOrderAccept(_paramOrderAccept *OrderAccept) error {
	r._paramOrderAccept = _paramOrderAccept
	r.Set("param_order_accept", _paramOrderAccept)
	return nil
}

// GetParamOrderAccept ParamOrderAccept Getter
func (r TaobaoalitripcarorderacceptAPIRequest) GetParamOrderAccept() *OrderAccept {
	return r._paramOrderAccept
}
