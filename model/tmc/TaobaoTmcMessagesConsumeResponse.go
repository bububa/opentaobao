package tmc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
消费多条消息 APIResponse
taobao.tmc.messages.consume

消费多条消息。消费时如果没有返回消息，建议做控制，不要一直调api，浪费应用的流量。如对程序做好优化，若没有消息则，sleep 100ms 等。
*/
type TaobaoTmcMessagesConsumeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmc_messages_consume_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 消息列表
    
    Messages   []TmcMessage `json:"messages,omitempty" xml:"