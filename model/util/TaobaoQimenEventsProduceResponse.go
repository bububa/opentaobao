package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量发送奇门事件 API返回值 
taobao.qimen.events.produce

批量发送消息
*/
type TaobaoQimenEventsProduceAPIResponse struct {
    model.CommonResponse
    TaobaoQimenEventsProduceResponse
}

// 批量发送奇门事件 成功返回结果
type TaobaoQimenEventsProduceResponse struct {
    XMLName xml.Name `xml:"qimen_events_produce_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否全部成功
    IsAllSuccess   bool `json:"is_all_success,omitempty" xml:"is_all_success,omitempty"`
    // 发送结果，与发送时的参数顺序一致。如果is_all_success为true时，不用校验result是否成功
    Results   []QimenResult `json:"results,omitempty" xml:"results>qimen_result,omitempty"`
}
