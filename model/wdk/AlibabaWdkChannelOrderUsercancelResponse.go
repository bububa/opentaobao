package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
用户发起售中取消 APIResponse
alibaba.wdk.channel.order.usercancel

用户发起售中取消
*/
type AlibabaWdkChannelOrderUsercancelAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkChannelOrderUsercancelResponse `json:"alibaba_wdk_channel_order_usercancel_response,omitempty"` 
    AlibabaWdkChannelOrderUsercancelResponse
}

/* model for simplify = false
type AlibabaWdkChannelOrderUsercancelResponse struct {

    // 返回结果
    
    ApiResult  *struct {
        AlibabaWdkChannelOrderUsercancelApiResult  *AlibabaWdkChannelOrderUsercancelApiResult `json:"alibaba_wdk_channel_order_usercancel_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaWdkChannelOrderUsercancelResponse struct {

    // 返回结果
    ApiResult   *AlibabaWdkChannelOrderUsercancelApiResult `json:"api_result,omitempty"`

}
