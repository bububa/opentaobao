package icbudropshipping

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
alibaba dropshipping 支付代扣 API请求
alibaba.dropshipping.order.pay

alibaba dropshipping 支付代扣
*/
type AlibabaDropshippingOrderPayRequest struct {
    model.Params
    // request model
    _paramOrderPayRequest   *OrderPayRequest
}

// 初始化AlibabaDropshippingOrderPayRequest对象
func NewAlibabaDropshippingOrderPayRequest() *AlibabaDropshippingOrderPayRequest{
    return &AlibabaDropshippingOrderPayRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDropshippingOrderPayRequest) GetApiMethodName() string {
    return "alibaba.dropshipping.order.pay"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDropshippingOrderPayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamOrderPayRequest Setter
// request model
func (r *AlibabaDropshippingOrderPayRequest) SetParamOrderPayRequest(_paramOrderPayRequest *OrderPayRequest) error {
    r._paramOrderPayRequest = _paramOrderPayRequest
    r.Set("param_order_pay_request", _paramOrderPayRequest)
    return nil
}

// ParamOrderPayRequest Getter
func (r AlibabaDropshippingOrderPayRequest) GetParamOrderPayRequest() *OrderPayRequest {
    return r._paramOrderPayRequest
}
