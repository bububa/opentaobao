package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取退票申请详情 API请求
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

// 初始化TaobaoAlitripIeAgentRefundGetRequest对象
func NewTaobaoAlitripIeAgentRefundGetRequest() *TaobaoAlitripIeAgentRefundGetRequest{
    return &TaobaoAlitripIeAgentRefundGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundGetRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.refund.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApplyId Setter
// 退票申请单id
func (r *TaobaoAlitripIeAgentRefundGetRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

// ApplyId Getter
func (r TaobaoAlitripIeAgentRefundGetRequest) GetApplyId() int64 {
    return r.applyId
}
// AgentId Setter
// 代理商id
func (r *TaobaoAlitripIeAgentRefundGetRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

// AgentId Getter
func (r TaobaoAlitripIeAgentRefundGetRequest) GetAgentId() int64 {
    return r.agentId
}
