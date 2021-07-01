package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同意退票 API请求
taobao.alitrip.ie.agent.refund.agree

卖家同意买家退票申请
*/
type TaobaoAlitripIeAgentRefundAgreeAPIRequest struct {
    model.Params
    // 退款金额
    _refundMoney   int64
    // 申请单id
    _applyId   int64
    // 订单id
    _orderId   int64
    // 回复信息
    _agentAnswer   string
    // 代理商id
    _agentId   int64
}

// 初始化TaobaoAlitripIeAgentRefundAgreeAPIRequest对象
func NewTaobaoAlitripIeAgentRefundAgreeRequest() *TaobaoAlitripIeAgentRefundAgreeAPIRequest{
    return &TaobaoAlitripIeAgentRefundAgreeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundAgreeAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.refund.agree"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundAgreeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundMoney Setter
// 退款金额
func (r *TaobaoAlitripIeAgentRefundAgreeAPIRequest) SetRefundMoney(_refundMoney int64) error {
    r._refundMoney = _refundMoney
    r.Set("refund_money", _refundMoney)
    return nil
}

// RefundMoney Getter
func (r TaobaoAlitripIeAgentRefundAgreeAPIRequest) GetRefundMoney() int64 {
    return r._refundMoney
}
// ApplyId Setter
// 申请单id
func (r *TaobaoAlitripIeAgentRefundAgreeAPIRequest) SetApplyId(_applyId int64) error {
    r._applyId = _applyId
    r.Set("apply_id", _applyId)
    return nil
}

// ApplyId Getter
func (r TaobaoAlitripIeAgentRefundAgreeAPIRequest) GetApplyId() int64 {
    return r._applyId
}
// OrderId Setter
// 订单id
func (r *TaobaoAlitripIeAgentRefundAgreeAPIRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoAlitripIeAgentRefundAgreeAPIRequest) GetOrderId() int64 {
    return r._orderId
}
// AgentAnswer Setter
// 回复信息
func (r *TaobaoAlitripIeAgentRefundAgreeAPIRequest) SetAgentAnswer(_agentAnswer string) error {
    r._agentAnswer = _agentAnswer
    r.Set("agent_answer", _agentAnswer)
    return nil
}

// AgentAnswer Getter
func (r TaobaoAlitripIeAgentRefundAgreeAPIRequest) GetAgentAnswer() string {
    return r._agentAnswer
}
// AgentId Setter
// 代理商id
func (r *TaobaoAlitripIeAgentRefundAgreeAPIRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoAlitripIeAgentRefundAgreeAPIRequest) GetAgentId() int64 {
    return r._agentId
}
