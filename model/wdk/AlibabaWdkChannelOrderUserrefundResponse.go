package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户发起售后退款(整单/部分) APIResponse
alibaba.wdk.channel.order.userrefund

用户发起售后退款(整单/部分)
*/
type AlibabaWdkChannelOrderUserrefundAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_channel_order_userrefund_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    ApiResult   *AlibabaWdkChannelOrderUserrefundApiResult `json:"api_result,omitempty" xml:"