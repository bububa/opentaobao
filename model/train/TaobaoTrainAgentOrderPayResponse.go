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
    Response *TaobaoTrainAgentOrderPayResponse `json:"taobao_train_agent_order_pay_response,omitempty"`
}

type TaobaoTrainAgentOrderPayResponse struct {

    // 成功返回
    ExtendParams   string `json:"extend_params,omitempty"`

}
