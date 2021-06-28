package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单状态变更 APIResponse
alibaba.wdk.channel.order.status.update

订单状态变更
*/
type AlibabaWdkChannelOrderStatusUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_channel_order_status_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    ApiResult   *AlibabaWdkChannelOrderStatusUpdateApiResult `json:"api_result,omitempty" xml:"