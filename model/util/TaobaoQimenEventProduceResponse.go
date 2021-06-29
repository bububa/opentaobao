package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发出奇门事件 API返回值 
taobao.qimen.event.produce

当订单被处理时，用于通知奇门系统。
*/
type TaobaoQimenEventProduceAPIResponse struct {
    model.CommonResponse
    TaobaoQimenEventProduceResponse
}

// 发出奇门事件 成功返回结果
type TaobaoQimenEventProduceResponse struct {
    XMLName xml.Name `xml:"qimen_event_produce_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
