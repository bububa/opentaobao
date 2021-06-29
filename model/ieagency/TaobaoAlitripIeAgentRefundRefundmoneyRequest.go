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
    _applyId   int64
    // 代理商id
    _agentId   int64
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
func (r *TaobaoAlitripIeAgentRefundRefundmoneyRequest) SetApplyId(_applyId int64) error {
    r._applyId = _applyId
    r.Set("apply_id", _applyId)
    return nil
}

// ApplyId Getter
func (r TaobaoAlitripIeAgentRefundRefundmoneyRequest) GetApplyId() int64 {
    return r._applyId
}
// AgentId Setter
// 代理商id
func (r *TaobaoAlitripIeAgentRefundRefundmoneyRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoAlitripIeAgentRefundRefundmoneyRequest) GetAgentId() int64 {
    return r._agentId
}
