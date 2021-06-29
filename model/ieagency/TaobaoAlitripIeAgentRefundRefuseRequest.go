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
    _applyId   int64
    // 订单id
    _orderId   int64
    // 代理商回复
    _agentAnswer   string
    // 代理商id
    _agentId   int64
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
func (r *TaobaoAlitripIeAgentRefundRefuseRequest) SetApplyId(_applyId int64) error {
    r._applyId = _applyId
    r.Set("apply_id", _applyId)
    return nil
}

// ApplyId Getter
func (r TaobaoAlitripIeAgentRefundRefuseRequest) GetApplyId() int64 {
    return r._applyId
}
// OrderId Setter
// 订单id
func (r *TaobaoAlitripIeAgentRefundRefuseRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoAlitripIeAgentRefundRefuseRequest) GetOrderId() int64 {
    return r._orderId
}
// AgentAnswer Setter
// 代理商回复
func (r *TaobaoAlitripIeAgentRefundRefuseRequest) SetAgentAnswer(_agentAnswer string) error {
    r._agentAnswer = _agentAnswer
    r.Set("agent_answer", _agentAnswer)
    return nil
}

// AgentAnswer Getter
func (r TaobaoAlitripIeAgentRefundRefuseRequest) GetAgentAnswer() string {
    return r._agentAnswer
}
// AgentId Setter
// 代理商id
func (r *TaobaoAlitripIeAgentRefundRefuseRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoAlitripIeAgentRefundRefuseRequest) GetAgentId() int64 {
    return r._agentId
}
