package tmc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
确认消费消息的状态 APIResponse
taobao.tmc.messages.confirm

确认消费消息的状态
*/
type TaobaoTmcMessagesConfirmAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTmcMessagesConfirmResponse `json:"tmc_messages_confirm_response,omitempty"` 
    TaobaoTmcMessagesConfirmResponse
}

/* model for simplify = false
type TaobaoTmcMessagesConfirmResponse struct {

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoTmcMessagesConfirmResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
