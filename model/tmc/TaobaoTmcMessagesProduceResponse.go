package tmc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量发送消息 APIResponse
taobao.tmc.messages.produce

批量发送消息
*/
type TaobaoTmcMessagesProduceAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmc_messages_produce_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否全部成功
    
    IsAllSuccess   bool `json:"is_all_success,omitempty" xml:"