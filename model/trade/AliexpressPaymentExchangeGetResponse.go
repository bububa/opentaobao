package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
getExchange APIResponse
aliexpress.payment.exchange.get

提供国际汇率服务
*/
type AliexpressPaymentExchangeGetAPIResponse struct {
    model.CommonResponse
    // Response *AliexpressPaymentExchangeGetResponse `json:"aliexpress_payment_exchange_get_response,omitempty"` 
    AliexpressPaymentExchangeGetResponse
}

/* model for simplify = false
type AliexpressPaymentExchangeGetResponse struct {

    // 接口返回model
    
    Result  *struct {
        AliexpressPaymentExchangeGetResult  *AliexpressPaymentExchangeGetResult `json:"aliexpress_payment_exchange_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AliexpressPaymentExchangeGetResponse struct {

    // 接口返回model
    Result   *AliexpressPaymentExchangeGetResult `json:"result,omitempty"`

}
