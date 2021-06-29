package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退款确认 APIResponse
alibaba.wdk.channel.order.refund.confirm

退款确认
*/
type AlibabaWdkChannelOrderRefundConfirmAPIResponse struct {
    model.CommonResponse
    AlibabaWdkChannelOrderRefundConfirmResponse
}

type AlibabaWdkChannelOrderRefundConfirmResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_channel_order_refund_confirm_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    ApiResult   *AlibabaWdkChannelOrderRefundConfirmApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`

    
}
