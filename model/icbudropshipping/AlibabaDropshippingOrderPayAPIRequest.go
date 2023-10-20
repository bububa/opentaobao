package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadropshippingorderpayAPIRequest alibaba dropshipping 支付代扣 API请求
// alibaba.dropshipping.order.pay
//
// alibaba dropshipping 支付代扣
type AlibabadropshippingorderpayAPIRequest struct {
	model.Params
	// request model
	_paramOrderPayRequest *OrderPayRequest
}

// NewAlibabadropshippingorderpayRequest 初始化AlibabadropshippingorderpayAPIRequest对象
func NewAlibabadropshippingorderpayRequest() *AlibabadropshippingorderpayAPIRequest {
	return &AlibabadropshippingorderpayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadropshippingorderpayAPIRequest) GetApiMethodName() string {
	return "alibaba.dropshipping.order.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadropshippingorderpayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadropshippingorderpayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderPayRequest is ParamOrderPayRequest Setter
// request model
func (r *AlibabadropshippingorderpayAPIRequest) SetParamOrderPayRequest(_paramOrderPayRequest *OrderPayRequest) error {
	r._paramOrderPayRequest = _paramOrderPayRequest
	r.Set("param_order_pay_request", _paramOrderPayRequest)
	return nil
}

// GetParamOrderPayRequest ParamOrderPayRequest Getter
func (r AlibabadropshippingorderpayAPIRequest) GetParamOrderPayRequest() *OrderPayRequest {
	return r._paramOrderPayRequest
}
