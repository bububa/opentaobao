package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认退款 APIRequest
taobao.alitrip.ie.agent.refund.refundmoney

卖家进行退款操作
*/
type TaobaoAlitripIeAgentRefundRefundmoneyRequest struct {
    model.Params

    // 退票申请单id
    applyId   int64 

    // 代理商id
    agentId   int64 

}

func NewTaobaoAlitripIeAgentRefundRefundmoneyRequest() *TaobaoAlitripIeAgentRefundRefundmoneyRequest{
    return &TaobaoAlitripIeAgentRefundRefundmoneyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripIeAgentRefundRefundmoneyRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.refund.refundmoney"
}

func (r TaobaoAlitripIeAgentRefundRefundmoneyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripIeAgentRefundRefundmoneyRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r TaobaoAlitripIeAgentRefundRefundmoneyRequest) GetApplyId() int64 {
    return r.applyId
}

func (r *TaobaoAlitripIeAgentRefundRefundmoneyRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoAlitripIeAgentRefundRefundmoneyRequest) GetAgentId() int64 {
    return r.agentId
}

