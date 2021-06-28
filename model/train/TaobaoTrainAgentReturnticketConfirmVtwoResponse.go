package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
退票通知 APIResponse
taobao.train.agent.returnticket.confirm.vtwo

火车票代理商接口——退票通知回调
*/
type TaobaoTrainAgentReturnticketConfirmVtwoAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTrainAgentReturnticketConfirmVtwoResponse `json:"train_agent_returnticket_confirm_vtwo_response,omitempty"` 
    TaobaoTrainAgentReturnticketConfirmVtwoResponse
}

/* model for simplify = false
type TaobaoTrainAgentReturnticketConfirmVtwoResponse struct {

    // resultCode
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // resultMsg
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

}
*/

type TaobaoTrainAgentReturnticketConfirmVtwoResponse struct {

    // resultCode
    ResultCode   string `json:"result_code,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

    // resultMsg
    ResultMsg   string `json:"result_msg,omitempty"`

}
