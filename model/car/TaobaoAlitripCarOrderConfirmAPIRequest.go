package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripcarorderconfirmAPIRequest 司机应答API API请求
// taobao.alitrip.car.order.confirm
//
// 航旅事业群-度假事业部-旅行用车项目组对外部服务商提供的司机应答回调接口
type TaobaoalitripcarorderconfirmAPIRequest struct {
	model.Params
	// 请求对象
	_paramOrderConfirm *OrderConfirm
}

// NewTaobaoalitripcarorderconfirmRequest 初始化TaobaoalitripcarorderconfirmAPIRequest对象
func NewTaobaoalitripcarorderconfirmRequest() *TaobaoalitripcarorderconfirmAPIRequest {
	return &TaobaoalitripcarorderconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripcarorderconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.order.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripcarorderconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripcarorderconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderConfirm is ParamOrderConfirm Setter
// 请求对象
func (r *TaobaoalitripcarorderconfirmAPIRequest) SetParamOrderConfirm(_paramOrderConfirm *OrderConfirm) error {
	r._paramOrderConfirm = _paramOrderConfirm
	r.Set("param_order_confirm", _paramOrderConfirm)
	return nil
}

// GetParamOrderConfirm ParamOrderConfirm Getter
func (r TaobaoalitripcarorderconfirmAPIRequest) GetParamOrderConfirm() *OrderConfirm {
	return r._paramOrderConfirm
}
