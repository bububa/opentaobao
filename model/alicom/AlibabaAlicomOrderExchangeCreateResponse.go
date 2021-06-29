package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商积分兑换接口 APIResponse
alibaba.alicom.order.exchange.create

代理商调用该接口来进行积分兑换
*/
type AlibabaAlicomOrderExchangeCreateAPIResponse struct {
    model.CommonResponse
    AlibabaAlicomOrderExchangeCreateResponse
}

type AlibabaAlicomOrderExchangeCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_alicom_order_exchange_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
