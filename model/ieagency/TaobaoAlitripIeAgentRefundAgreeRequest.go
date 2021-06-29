package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同意退票 APIRequest
taobao.alitrip.ie.agent.refund.agree

卖家同意买家退票申请
*/
type TaobaoAlitripIeAgentRefundAgreeRequest struct {
    model.Params

    // 退款金额
    refundMoney   int64 

    // 申请单id
    applyId   int64 

    // 订单id
    orderId   int64 

    // 回复信息
    agentAnswer   string 

    // 代理商id
    agentId   int64 

}

func NewTaobaoAlitripIeAgentRefundAgreeRequest() *TaobaoAlitripIeAgentRefundAgreeRequest{
    return &TaobaoAlitripIeAgentRefundAgreeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripIeAgentRefundAgreeRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.refund.agree"
}

func (r TaobaoAlitripIeAgentRefundAgreeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripIeAgentRefundAgreeRequest) SetRefundMoney(refundMoney int64) error {
    r.refundMoney = refundMoney
    r.Set("refund_money", refundMoney)
    return nil
}

func (r TaobaoAlitripIeAgentRefundAgreeRequest) GetRefundMoney() int64 {
    return r.refundMoney
}

func (r *TaobaoAlitripIeAgentRefundAgreeRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r TaobaoAlitripIeAgentRefundAgreeRequest) GetApplyId() int64 {
    return r.applyId
}

func (r *TaobaoAlitripIeAgentRefundAgreeRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoAlitripIeAgentRefundAgreeRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TaobaoAlitripIeAgentRefundAgreeRequest) SetAgentAnswer(agentAnswer string) error {
    r.agentAnswer = agentAnswer
    r.Set("agent_answer", agentAnswer)
    return nil
}

func (r TaobaoAlitripIeAgentRefundAgreeRequest) GetAgentAnswer() string {
    return r.agentAnswer
}

func (r *TaobaoAlitripIeAgentRefundAgreeRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoAlitripIeAgentRefundAgreeRequest) GetAgentId() int64 {
    return r.agentId
}

