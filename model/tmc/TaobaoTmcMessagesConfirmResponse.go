package tmc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
确认消费消息的状态 APIResponse
taobao.tmc.messages.confirm

确认消费消息的状态
*/
type TaobaoTmcMessagesConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoTmcMessagesConfirmResponse
}

type TaobaoTmcMessagesConfirmResponse struct {
    XMLName xml.Name `xml:"tmc_messages_confirm_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
