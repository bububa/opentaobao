package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——确认出票是否成功 APIResponse
taobao.train.agent.bookticket.confirm

火车票代理商接口——确认出票是否成功
*/
type TaobaoTrainAgentBookticketConfirmAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTrainAgentBookticketConfirmResponse `json:"taobao_train_agent_bookticket_confirm_response,omitempty"`
}

type TaobaoTrainAgentBookticketConfirmResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
