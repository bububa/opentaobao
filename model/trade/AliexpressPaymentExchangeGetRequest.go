package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
getExchange API请求
aliexpress.payment.exchange.get

提供国际汇率服务
*/
type AliexpressPaymentExchangeGetRequest struct {
    model.Params
    // 系统自动生成
    checkoutExchangeRequest   *CheckoutExchangeRequest
}

// 初始化AliexpressPaymentExchangeGetRequest对象
func NewAliexpressPaymentExchangeGetRequest() *AliexpressPaymentExchangeGetRequest{
    return &AliexpressPaymentExchangeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressPaymentExchangeGetRequest) GetApiMethodName() string {
    return "aliexpress.payment.exchange.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressPaymentExchangeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CheckoutExchangeRequest Setter
// 系统自动生成
func (r *AliexpressPaymentExchangeGetRequest) SetCheckoutExchangeRequest(checkoutExchangeRequest *CheckoutExchangeRequest) error {
    r.checkoutExchangeRequest = checkoutExchangeRequest
    r.Set("checkout_exchange_request", checkoutExchangeRequest)
    return nil
}

// CheckoutExchangeRequest Getter
func (r AliexpressPaymentExchangeGetRequest) GetCheckoutExchangeRequest() *CheckoutExchangeRequest {
    return r.checkoutExchangeRequest
}
