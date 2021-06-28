package tmc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
消费多条消息 APIResponse
taobao.tmc.messages.consume

消费多条消息。消费时如果没有返回消息，建议做控制，不要一直调api，浪费应用的流量。如对程序做好优化，若没有消息则，sleep 100ms 等。
*/
type TaobaoTmcMessagesConsumeAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTmcMessagesConsumeResponse `json:"tmc_messages_consume_response,omitempty"` 
    TaobaoTmcMessagesConsumeResponse
}

/* model for simplify = false
type TaobaoTmcMessagesConsumeResponse struct {

    // 消息列表
    
    Messages  struct {
        TmcMessage  []TmcMessage `json:"tmc_message,omitempty"`
    } `json:"messages,omitempty"`
    

}
*/

type TaobaoTmcMessagesConsumeResponse struct {

    // 消息列表
    Messages   []TmcMessage `json:"messages,omitempty"`

}
