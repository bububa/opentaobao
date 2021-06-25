package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代购订单代付接口 APIRequest
taobao.train.agent.order.pay

代购订单代付接口
*/
type TaobaoTrainAgentOrderPayRequest struct {
    model.Params

    // 入参对象
    agentPayOrderParam   *AgentPayOrderParam 

}

func NewTaobaoTrainAgentOrderPayRequest() *TaobaoTrainAgentOrderPayRequest{
    return &TaobaoTrainAgentOrderPayRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentOrderPayRequest) GetApiMethodName() string {
    return "taobao.train.agent.order.pay"
}

func (r TaobaoTrainAgentOrderPayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentOrderPayRequest) SetAgentPayOrderParam(agentPayOrderParam *AgentPayOrderParam) error {
    r.agentPayOrderParam = agentPayOrderParam
    r.Set("agent_pay_order_param", agentPayOrderParam)
    return nil
}

func (r TaobaoTrainAgentOrderPayRequest) GetAgentPayOrderParam() *AgentPayOrderParam {
    return r.agentPayOrderParam
}

