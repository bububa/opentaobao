package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建订单 APIResponse
alibaba.wdk.channel.order.create

外部商家创建订单
*/
type AlibabaWdkChannelOrderCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_channel_order_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    ApiResult   *AlibabaWdkChannelOrderCreateApiResult `json:"api_result,omitempty" xml:"