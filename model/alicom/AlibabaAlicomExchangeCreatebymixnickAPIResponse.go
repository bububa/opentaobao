package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商积分兑换接口tae API返回值 
alibaba.alicom.exchange.createbymixnick

代理商调用该接口来进行积分兑换，tae
*/
type AlibabaAlicomExchangeCreatebymixnickAPIResponse struct {
    model.CommonResponse
    AlibabaAlicomExchangeCreatebymixnickAPIResponseModel
}

// 代理商积分兑换接口tae 成功返回结果
type AlibabaAlicomExchangeCreatebymixnickAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alicom_exchange_createbymixnick_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
