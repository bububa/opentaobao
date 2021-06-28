package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发出奇门事件 APIResponse
taobao.qimen.event.produce

当订单被处理时，用于通知奇门系统。
*/
type TaobaoQimenEventProduceAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenEventProduceResponse `json:"qimen_event_produce_response,omitempty"` 
    TaobaoQimenEventProduceResponse
}

/* model for simplify = false
type TaobaoQimenEventProduceResponse struct {

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoQimenEventProduceResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
