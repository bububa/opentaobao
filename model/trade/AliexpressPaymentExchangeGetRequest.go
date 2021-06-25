package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
getExchange APIRequest
aliexpress.payment.exchange.get

提供国际汇率服务
*/
type AliexpressPaymentExchangeGetRequest struct {
    model.Params

    // 系统自动生成
    checkoutExchangeRequest   *CheckoutExchangeRequest 

}

func NewAliexpressPaymentExchangeGetRequest() *AliexpressPaymentExchangeGetRequest{
    return &AliexpressPaymentExchangeGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressPaymentExchangeGetRequest) GetApiMethodName() string {
    return "aliexpress.payment.exchange.get"
}

func (r AliexpressPaymentExchangeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressPaymentExchangeGetRequest) SetCheckoutExchangeRequest(checkoutExchangeRequest *CheckoutExchangeRequest) error {
    r.checkoutExchangeRequest = checkoutExchangeRequest
    r.Set("checkout_exchange_request", checkoutExchangeRequest)
    return nil
}

func (r AliexpressPaymentExchangeGetRequest) GetCheckoutExchangeRequest() *CheckoutExchangeRequest {
    return r.checkoutExchangeRequest
}

