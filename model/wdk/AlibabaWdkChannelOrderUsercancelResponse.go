package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户发起售中取消 API返回值 
alibaba.wdk.channel.order.usercancel

用户发起售中取消
*/
type AlibabaWdkChannelOrderUsercancelAPIResponse struct {
    model.CommonResponse
    AlibabaWdkChannelOrderUsercancelResponse
}

// 用户发起售中取消 成功返回结果
type AlibabaWdkChannelOrderUsercancelResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_channel_order_usercancel_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    ApiResult   *AlibabaWdkChannelOrderUsercancelApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
