package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
订单状态变更 APIResponse
alibaba.wdk.channel.order.status.update

订单状态变更
*/
type AlibabaWdkChannelOrderStatusUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkChannelOrderStatusUpdateResponse `json:"alibaba_wdk_channel_order_status_update_response,omitempty"` 
    AlibabaWdkChannelOrderStatusUpdateResponse
}

/* model for simplify = false
type AlibabaWdkChannelOrderStatusUpdateResponse struct {

    // 返回结果
    
    ApiResult  *struct {
        AlibabaWdkChannelOrderStatusUpdateApiResult  *AlibabaWdkChannelOrderStatusUpdateApiResult `json:"alibaba_wdk_channel_order_status_update_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaWdkChannelOrderStatusUpdateResponse struct {

    // 返回结果
    ApiResult   *AlibabaWdkChannelOrderStatusUpdateApiResult `json:"api_result,omitempty"`

}
