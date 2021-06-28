package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
代理商手动退款接口 APIResponse
taobao.train.agent.handrefund.refundfee

火车票代理商手动退款接口
*/
type TaobaoTrainAgentHandrefundRefundfeeAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTrainAgentHandrefundRefundfeeResponse `json:"train_agent_handrefund_refundfee_response,omitempty"` 
    TaobaoTrainAgentHandrefundRefundfeeResponse
}

/* model for simplify = false
type TaobaoTrainAgentHandrefundRefundfeeResponse struct {

    // 是否成功标记
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 失败code
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // 失败文案
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

}
*/

type TaobaoTrainAgentHandrefundRefundfeeResponse struct {

    // 是否成功标记
    IsSuccess   bool `json:"is_success,omitempty"`

    // 失败code
    ResultCode   string `json:"result_code,omitempty"`

    // 失败文案
    ResultMsg   string `json:"result_msg,omitempty"`

}
