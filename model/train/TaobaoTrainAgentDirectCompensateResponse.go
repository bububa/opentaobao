package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——订单关闭实际出票成功审计接口 APIResponse
taobao.train.agent.direct.compensate

代购直连订单平台关单但是代理商出票成功补偿接口
*/
type TaobaoTrainAgentDirectCompensateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTrainAgentDirectCompensateResponse `json:"train_agent_direct_compensate_response,omitempty"` 
    TaobaoTrainAgentDirectCompensateResponse
}

/* model for simplify = false
type TaobaoTrainAgentDirectCompensateResponse struct {

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // resultCode
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // resultMsg
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

}
*/

type TaobaoTrainAgentDirectCompensateResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // resultCode
    ResultCode   string `json:"result_code,omitempty"`

    // resultMsg
    ResultMsg   string `json:"result_msg,omitempty"`

}
