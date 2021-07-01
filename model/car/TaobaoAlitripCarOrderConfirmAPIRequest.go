package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripCarOrderConfirmAPIRequest
司机应答API API请求
taobao.alitrip.car.order.confirm

航旅事业群-度假事业部-旅行用车项目组对外部服务商提供的司机应答回调接口 */
type TaobaoAlitripCarOrderConfirmAPIRequest struct {
	model.Params
	// 请求对象
	_paramOrderConfirm *OrderConfirm
}

// NewTaobaoAlitripCarOrderConfirmRequest 初始化TaobaoAlitripCarOrderConfirmAPIRequest对象
func NewTaobaoAlitripCarOrderConfirmRequest() *TaobaoAlitripCarOrderConfirmAPIRequest {
	return &TaobaoAlitripCarOrderConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarOrderConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.order.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarOrderConfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamOrderConfirm Setter
// 请求对象
func (r *TaobaoAlitripCarOrderConfirmAPIRequest) SetParamOrderConfirm(_paramOrderConfirm *OrderConfirm) error {
	r._paramOrderConfirm = _paramOrderConfirm
	r.Set("param_order_confirm", _paramOrderConfirm)
	return nil
}

// Get ParamOrderConfirm Getter
func (r TaobaoAlitripCarOrderConfirmAPIRequest) GetParamOrderConfirm() *OrderConfirm {
	return r._paramOrderConfirm
}
