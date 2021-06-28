package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建订单 APIResponse
alibaba.wdk.channel.order.create

外部商家创建订单
*/
type AlibabaWdkChannelOrderCreateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkChannelOrderCreateResponse `json:"alibaba_wdk_channel_order_create_response,omitempty"` 
    AlibabaWdkChannelOrderCreateResponse
}

/* model for simplify = false
type AlibabaWdkChannelOrderCreateResponse struct {

    // 返回结果
    
    ApiResult  *struct {
        AlibabaWdkChannelOrderCreateApiResult  *AlibabaWdkChannelOrderCreateApiResult `json:"alibaba_wdk_channel_order_create_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaWdkChannelOrderCreateResponse struct {

    // 返回结果
    ApiResult   *AlibabaWdkChannelOrderCreateApiResult `json:"api_result,omitempty"`

}
