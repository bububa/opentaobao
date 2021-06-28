package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——确认出票是否成功v2--增加鉴权校验 APIResponse
taobao.train.agent.bookticket.confirm.vtwo

火车票代理商接口——确认出票是否成功
*/
type TaobaoTrainAgentBookticketConfirmVtwoAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTrainAgentBookticketConfirmVtwoResponse `json:"train_agent_bookticket_confirm_vtwo_response,omitempty"` 
    TaobaoTrainAgentBookticketConfirmVtwoResponse
}

/* model for simplify = false
type TaobaoTrainAgentBookticketConfirmVtwoResponse struct {

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoTrainAgentBookticketConfirmVtwoResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
