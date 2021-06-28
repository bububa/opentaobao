package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
用户发起售后退款(整单/部分) APIResponse
alibaba.wdk.channel.order.userrefund

用户发起售后退款(整单/部分)
*/
type AlibabaWdkChannelOrderUserrefundAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkChannelOrderUserrefundResponse `json:"alibaba_wdk_channel_order_userrefund_response,omitempty"` 
    AlibabaWdkChannelOrderUserrefundResponse
}

/* model for simplify = false
type AlibabaWdkChannelOrderUserrefundResponse struct {

    // 返回结果
    
    ApiResult  *struct {
        AlibabaWdkChannelOrderUserrefundApiResult  *AlibabaWdkChannelOrderUserrefundApiResult `json:"alibaba_wdk_channel_order_userrefund_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaWdkChannelOrderUserrefundResponse struct {

    // 返回结果
    ApiResult   *AlibabaWdkChannelOrderUserrefundApiResult `json:"api_result,omitempty"`

}
