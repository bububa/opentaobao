package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量发送奇门事件 APIResponse
taobao.qimen.events.produce

批量发送消息
*/
type TaobaoQimenEventsProduceAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_events_produce_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否全部成功
    
    IsAllSuccess   bool `json:"is_all_success,omitempty" xml:"