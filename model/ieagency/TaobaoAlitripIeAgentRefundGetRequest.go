package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取退票申请详情 APIRequest
taobao.alitrip.ie.agent.refund.get

获取退票申请详情
*/
type TaobaoAlitripIeAgentRefundGetRequest struct {
    model.Params

    // 退票申请单id
    applyId   int64 

    // 代理商id
    agentId   int64 

}

func NewTaobaoAlitripIeAgentRefundGetRequest() *TaobaoAlitripIeAgentRefundGetRequest{
    return &TaobaoAlitripIeAgentRefundGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripIeAgentRefundGetRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.refund.get"
}

func (r TaobaoAlitripIeAgentRefundGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripIeAgentRefundGetRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r TaobaoAlitripIeAgentRefundGetRequest) GetApplyId() int64 {
    return r.applyId
}

func (r *TaobaoAlitripIeAgentRefundGetRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoAlitripIeAgentRefundGetRequest) GetAgentId() int64 {
    return r.agentId
}

