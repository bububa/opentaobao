package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
拒绝退票申请 APIRequest
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

func NewTaobaoAlitripIeAgentRefundRefuseRequest() *TaobaoAlitripIeAgentRefundRefuseRequest{
    return &TaobaoAlitripIeAgentRefundRefuseRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripIeAgentRefundRefuseRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.refund.refuse"
}

func (r TaobaoAlitripIeAgentRefundRefuseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripIeAgentRefundRefuseRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r TaobaoAlitripIeAgentRefundRefuseRequest) GetApplyId() int64 {
    return r.applyId
}

func (r *TaobaoAlitripIeAgentRefundRefuseRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoAlitripIeAgentRefundRefuseRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TaobaoAlitripIeAgentRefundRefuseRequest) SetAgentAnswer(agentAnswer string) error {
    r.agentAnswer = agentAnswer
    r.Set("agent_answer", agentAnswer)
    return nil
}

func (r TaobaoAlitripIeAgentRefundRefuseRequest) GetAgentAnswer() string {
    return r.agentAnswer
}

func (r *TaobaoAlitripIeAgentRefundRefuseRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoAlitripIeAgentRefundRefuseRequest) GetAgentId() int64 {
    return r.agentId
}

