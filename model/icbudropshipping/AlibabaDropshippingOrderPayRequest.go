package icbudropshipping

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
alibaba dropshipping 支付代扣 APIRequest
alibaba.dropshipping.order.pay

alibaba dropshipping 支付代扣
*/
type AlibabaDropshippingOrderPayRequest struct {
    model.Params

    // request model
    paramOrderPayRequest   *OrderPayRequest 

}

func NewAlibabaDropshippingOrderPayRequest() *AlibabaDropshippingOrderPayRequest{
    return &AlibabaDropshippingOrderPayRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDropshippingOrderPayRequest) GetApiMethodName() string {
    return "alibaba.dropshipping.order.pay"
}

func (r AlibabaDropshippingOrderPayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDropshippingOrderPayRequest) SetParamOrderPayRequest(paramOrderPayRequest *OrderPayRequest) error {
    r.paramOrderPayRequest = paramOrderPayRequest
    r.Set("param_order_pay_request", paramOrderPayRequest)
    return nil
}

func (r AlibabaDropshippingOrderPayRequest) GetParamOrderPayRequest() *OrderPayRequest {
    return r.paramOrderPayRequest
}

