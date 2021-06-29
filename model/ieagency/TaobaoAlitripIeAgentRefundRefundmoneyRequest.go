package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认退款 API请求
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

// 初始化TaobaoAlitripIeAgentRefundRefundmoneyRequest对象
func NewTaobaoAlitripIeAgentRefundRefundmoneyRequest() *TaobaoAlitripIeAgentRefundRefundmoneyRequest{
    return &TaobaoAlitripIeAgentRefundRefundmoneyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundRefundmoneyRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.refund.refundmoney"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundRefundmoneyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApplyId Setter
// 退票申请单id
func (r *TaobaoAlitripIeAgentRefundRefundmoneyRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

// ApplyId Getter
func (r TaobaoAlitripIeAgentRefundRefundmoneyRequest) GetApplyId() int64 {
    return r.applyId
}
// AgentId Setter
// 代理商id
func (r *TaobaoAlitripIeAgentRefundRefundmoneyRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

// AgentId Getter
func (r TaobaoAlitripIeAgentRefundRefundmoneyRequest) GetAgentId() int64 {
    return r.agentId
}
