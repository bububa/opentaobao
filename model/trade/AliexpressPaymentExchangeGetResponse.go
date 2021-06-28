package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
getExchange APIResponse
aliexpress.payment.exchange.get

提供国际汇率服务
*/
type AliexpressPaymentExchangeGetAPIResponse struct {
    model.CommonResponse
    AliexpressPaymentExchangeGetResponse
}

type AliexpressPaymentExchangeGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_payment_exchange_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *AliexpressPaymentExchangeGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
