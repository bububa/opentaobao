package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代购订单代付接口 API请求
taobao.train.agent.order.pay

代购订单代付接口
*/
type TaobaoTrainAgentOrderPayRequest struct {
    model.Params
    // 入参对象
    _agentPayOrderParam   *AgentPayOrderParam
}

// 初始化TaobaoTrainAgentOrderPayRequest对象
func NewTaobaoTrainAgentOrderPayRequest() *TaobaoTrainAgentOrderPayRequest{
    return &TaobaoTrainAgentOrderPayRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentOrderPayRequest) GetApiMethodName() string {
    return "taobao.train.agent.order.pay"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentOrderPayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AgentPayOrderParam Setter
// 入参对象
func (r *TaobaoTrainAgentOrderPayRequest) SetAgentPayOrderParam(_agentPayOrderParam *AgentPayOrderParam) error {
    r._agentPayOrderParam = _agentPayOrderParam
    r.Set("agent_pay_order_param", _agentPayOrderParam)
    return nil
}

// AgentPayOrderParam Getter
func (r TaobaoTrainAgentOrderPayRequest) GetAgentPayOrderParam() *AgentPayOrderParam {
    return r._agentPayOrderParam
}
