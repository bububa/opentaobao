package icbudropshipping

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDropshippingOrderPayAPIRequest alibaba dropshipping 支付代扣 API请求
// alibaba.dropshipping.order.pay
//
// alibaba dropshipping 支付代扣
type AlibabaDropshippingOrderPayAPIRequest struct {
	model.Params
	// request model
	_paramOrderPayRequest *OrderPayRequest
}

// NewAlibabaDropshippingOrderPayRequest 初始化AlibabaDropshippingOrderPayAPIRequest对象
func NewAlibabaDropshippingOrderPayRequest() *AlibabaDropshippingOrderPayAPIRequest {
	return &AlibabaDropshippingOrderPayAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDropshippingOrderPayAPIRequest) Reset() {
	r._paramOrderPayRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDropshippingOrderPayAPIRequest) GetApiMethodName() string {
	return "alibaba.dropshipping.order.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDropshippingOrderPayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDropshippingOrderPayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderPayRequest is ParamOrderPayRequest Setter
// request model
func (r *AlibabaDropshippingOrderPayAPIRequest) SetParamOrderPayRequest(_paramOrderPayRequest *OrderPayRequest) error {
	r._paramOrderPayRequest = _paramOrderPayRequest
	r.Set("param_order_pay_request", _paramOrderPayRequest)
	return nil
}

// GetParamOrderPayRequest ParamOrderPayRequest Getter
func (r AlibabaDropshippingOrderPayAPIRequest) GetParamOrderPayRequest() *OrderPayRequest {
	return r._paramOrderPayRequest
}

var poolAlibabaDropshippingOrderPayAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDropshippingOrderPayRequest()
	},
}

// GetAlibabaDropshippingOrderPayRequest 从 sync.Pool 获取 AlibabaDropshippingOrderPayAPIRequest
func GetAlibabaDropshippingOrderPayAPIRequest() *AlibabaDropshippingOrderPayAPIRequest {
	return poolAlibabaDropshippingOrderPayAPIRequest.Get().(*AlibabaDropshippingOrderPayAPIRequest)
}

// ReleaseAlibabaDropshippingOrderPayAPIRequest 将 AlibabaDropshippingOrderPayAPIRequest 放入 sync.Pool
func ReleaseAlibabaDropshippingOrderPayAPIRequest(v *AlibabaDropshippingOrderPayAPIRequest) {
	v.Reset()
	poolAlibabaDropshippingOrderPayAPIRequest.Put(v)
}
