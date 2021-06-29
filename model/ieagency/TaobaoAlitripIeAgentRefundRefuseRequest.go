package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
拒绝退票申请 API请求
taobao.alitrip.ie.agent.refund.refuse

卖家拒绝退票退票申请
*/
type TaobaoAlitripIeAgentRefundRefuseRequest struct {
    model.Params
    // 退票申请单id
    applyId   int64
    // 订单id
    orderId   int64
    // 代理商回复
    agentAnswer   string
    // 代理商id
    agentId   int64
}

// 初始化TaobaoAlitripIeAgentRefundRefuseRequest对象
func NewTaobaoAlitripIeAgentRefundRefuseRequest() *TaobaoAlitripIeAgentRefundRefuseRequest{
    return &TaobaoAlitripIeAgentRefundRefuseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundRefuseRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.refund.refuse"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundRefuseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApplyId Setter
// 退票申请单id
func (r *TaobaoAlitripIeAgentRefundRefuseRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

// ApplyId Getter
func (r TaobaoAlitripIeAgentRefundRefuseRequest) GetApplyId() int64 {
    return r.applyId
}
// OrderId Setter
// 订单id
func (r *TaobaoAlitripIeAgentRefundRefuseRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TaobaoAlitripIeAgentRefundRefuseRequest) GetOrderId() int64 {
    return r.orderId
}
// AgentAnswer Setter
// 代理商回复
func (r *TaobaoAlitripIeAgentRefundRefuseRequest) SetAgentAnswer(agentAnswer string) error {
    r.agentAnswer = agentAnswer
    r.Set("agent_answer", agentAnswer)
    return nil
}

// AgentAnswer Getter
func (r TaobaoAlitripIeAgentRefundRefuseRequest) GetAgentAnswer() string {
    return r.agentAnswer
}
// AgentId Setter
// 代理商id
func (r *TaobaoAlitripIeAgentRefundRefuseRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

// AgentId Getter
func (r TaobaoAlitripIeAgentRefundRefuseRequest) GetAgentId() int64 {
    return r.agentId
}
