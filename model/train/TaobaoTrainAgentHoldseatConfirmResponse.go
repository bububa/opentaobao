package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——确认占座是否成功 APIResponse
taobao.train.agent.holdseat.confirm

火车票代理商接口——确认占座是否成功
*/
type TaobaoTrainAgentHoldseatConfirmAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTrainAgentHoldseatConfirmResponse `json:"train_agent_holdseat_confirm_response,omitempty"` 
    TaobaoTrainAgentHoldseatConfirmResponse
}

/* model for simplify = false
type TaobaoTrainAgentHoldseatConfirmResponse struct {

    // resultMsg
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

    // resultCode
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoTrainAgentHoldseatConfirmResponse struct {

    // resultMsg
    ResultMsg   string `json:"result_msg,omitempty"`

    // resultCode
    ResultCode   string `json:"result_code,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
