package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发出奇门事件 APIResponse
taobao.qimen.event.produce

当订单被处理时，用于通知奇门系统。
*/
type TaobaoQimenEventProduceAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_event_produce_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"