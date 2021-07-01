package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户发起售后退款(整单/部分) API返回值 
alibaba.wdk.channel.order.userrefund

用户发起售后退款(整单/部分)
*/
type AlibabaWdkChannelOrderUserrefundAPIResponse struct {
    model.CommonResponse
    AlibabaWdkChannelOrderUserrefundAPIResponseModel
}

// 用户发起售后退款(整单/部分) 成功返回结果
type AlibabaWdkChannelOrderUserrefundAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_channel_order_userrefund_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    ApiResult   *AlibabaWdkChannelOrderUserrefundApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
