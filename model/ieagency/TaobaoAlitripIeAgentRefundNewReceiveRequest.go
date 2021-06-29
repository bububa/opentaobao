package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家退票受理申请(对外) API请求
taobao.alitrip.ie.agent.refund.new.receive

允许代理商通过top接口受理退票申请
*/
type TaobaoAlitripIeAgentRefundNewReceiveRequest struct {
    model.Params
    // 退票申请单id
    applyId   int64
    // 代理商id
    agentId   int64
    // 订单id
    orderId   int64
}

// 初始化TaobaoAlitripIeAgentRefundNewReceiveRequest对象
func NewTaobaoAlitripIeAgentRefundNewReceiveRequest() *TaobaoAlitripIeAgentRefundNewReceiveRequest{
    return &TaobaoAlitripIeAgentRefundNewReceiveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundNewReceiveRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.refund.new.receive"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundNewReceiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApplyId Setter
// 退票申请单id
func (r *TaobaoAlitripIeAgentRefundNewReceiveRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

// ApplyId Getter
func (r TaobaoAlitripIeAgentRefundNewReceiveRequest) GetApplyId() int64 {
    return r.applyId
}
// AgentId Setter
// 代理商id
func (r *TaobaoAlitripIeAgentRefundNewReceiveRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

// AgentId Getter
func (r TaobaoAlitripIeAgentRefundNewReceiveRequest) GetAgentId() int64 {
    return r.agentId
}
// OrderId Setter
// 订单id
func (r *TaobaoAlitripIeAgentRefundNewReceiveRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TaobaoAlitripIeAgentRefundNewReceiveRequest) GetOrderId() int64 {
    return r.orderId
}
