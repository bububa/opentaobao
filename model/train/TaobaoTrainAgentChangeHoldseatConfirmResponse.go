package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——确认改签占座是否成功 APIResponse
taobao.train.agent.change.holdseat.confirm

火车票代理商接口——确认改签占座是否成功
*/
type TaobaoTrainAgentChangeHoldseatConfirmAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTrainAgentChangeHoldseatConfirmResponse `json:"train_agent_change_holdseat_confirm_response,omitempty"` 
    TaobaoTrainAgentChangeHoldseatConfirmResponse
}

/* model for simplify = false
type TaobaoTrainAgentChangeHoldseatConfirmResponse struct {

    // 是否成功标记
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // errorCode
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // errorMsg
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

}
*/

type TaobaoTrainAgentChangeHoldseatConfirmResponse struct {

    // 是否成功标记
    IsSuccess   bool `json:"is_success,omitempty"`

    // errorCode
    ResultCode   string `json:"result_code,omitempty"`

    // errorMsg
    ResultMsg   string `json:"result_msg,omitempty"`

}
