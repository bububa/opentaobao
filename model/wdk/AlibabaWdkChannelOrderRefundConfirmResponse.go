package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
退款确认 APIResponse
alibaba.wdk.channel.order.refund.confirm

退款确认
*/
type AlibabaWdkChannelOrderRefundConfirmAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkChannelOrderRefundConfirmResponse `json:"alibaba_wdk_channel_order_refund_confirm_response,omitempty"` 
    AlibabaWdkChannelOrderRefundConfirmResponse
}

/* model for simplify = false
type AlibabaWdkChannelOrderRefundConfirmResponse struct {

    // 返回结果
    
    ApiResult  *struct {
        AlibabaWdkChannelOrderRefundConfirmApiResult  *AlibabaWdkChannelOrderRefundConfirmApiResult `json:"alibaba_wdk_channel_order_refund_confirm_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaWdkChannelOrderRefundConfirmResponse struct {

    // 返回结果
    ApiResult   *AlibabaWdkChannelOrderRefundConfirmApiResult `json:"api_result,omitempty"`

}
