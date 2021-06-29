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
    _applyId   int64
    // 代理商id
    _agentId   int64
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
func (r *TaobaoAlitripIeAgentRefundGetRequest) SetApplyId(_applyId int64) error {
    r._applyId = _applyId
    r.Set("apply_id", _applyId)
    return nil
}

// ApplyId Getter
func (r TaobaoAlitripIeAgentRefundGetRequest) GetApplyId() int64 {
    return r._applyId
}
// AgentId Setter
// 代理商id
func (r *TaobaoAlitripIeAgentRefundGetRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoAlitripIeAgentRefundGetRequest) GetAgentId() int64 {
    return r._agentId
}
