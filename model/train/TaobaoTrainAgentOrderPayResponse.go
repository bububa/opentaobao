package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
代购订单代付接口 APIResponse
taobao.train.agent.order.pay

代购订单代付接口
*/
type TaobaoTrainAgentOrderPayAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTrainAgentOrderPayResponse `json:"train_agent_order_pay_response,omitempty"` 
    TaobaoTrainAgentOrderPayResponse
}

/* model for simplify = false
type TaobaoTrainAgentOrderPayResponse struct {

    // 成功返回
    
    ExtendParams   string `json:"extend_params,omitempty"`
    

}
*/

type TaobaoTrainAgentOrderPayResponse struct {

    // 成功返回
    ExtendParams   string `json:"extend_params,omitempty"`

}
