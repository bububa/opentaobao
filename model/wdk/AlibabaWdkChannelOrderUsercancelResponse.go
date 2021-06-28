package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户发起售中取消 APIResponse
alibaba.wdk.channel.order.usercancel

用户发起售中取消
*/
type AlibabaWdkChannelOrderUsercancelAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_channel_order_usercancel_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    ApiResult   *AlibabaWdkChannelOrderUsercancelApiResult `json:"api_result,omitempty" xml:"