package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家退票受理申请(对外) APIRequest
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

func NewTaobaoAlitripIeAgentRefundNewReceiveRequest() *TaobaoAlitripIeAgentRefundNewReceiveRequest{
    return &TaobaoAlitripIeAgentRefundNewReceiveRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripIeAgentRefundNewReceiveRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.refund.new.receive"
}

func (r TaobaoAlitripIeAgentRefundNewReceiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripIeAgentRefundNewReceiveRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r TaobaoAlitripIeAgentRefundNewReceiveRequest) GetApplyId() int64 {
    return r.applyId
}

func (r *TaobaoAlitripIeAgentRefundNewReceiveRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoAlitripIeAgentRefundNewReceiveRequest) GetAgentId() int64 {
    return r.agentId
}

func (r *TaobaoAlitripIeAgentRefundNewReceiveRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoAlitripIeAgentRefundNewReceiveRequest) GetOrderId() int64 {
    return r.orderId
}

