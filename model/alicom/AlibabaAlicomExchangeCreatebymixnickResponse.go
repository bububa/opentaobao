package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商积分兑换接口tae APIResponse
alibaba.alicom.exchange.createbymixnick

代理商调用该接口来进行积分兑换，tae
*/
type AlibabaAlicomExchangeCreatebymixnickAPIResponse struct {
    model.CommonResponse
    AlibabaAlicomExchangeCreatebymixnickResponse
}

type AlibabaAlicomExchangeCreatebymixnickResponse struct {
    XMLName xml.Name `xml:"alibaba_alicom_exchange_createbymixnick_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
