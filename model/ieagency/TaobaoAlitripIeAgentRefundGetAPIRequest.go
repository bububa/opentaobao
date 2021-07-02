package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentRefundGetAPIRequest 获取退票申请详情 API请求
// taobao.alitrip.ie.agent.refund.get
//
// 获取退票申请详情
type TaobaoAlitripIeAgentRefundGetAPIRequest struct {
	model.Params
	// 退票申请单id
	_applyId int64
	// 代理商id
	_agentId int64
}

// NewTaobaoAlitripIeAgentRefundGetRequest 初始化TaobaoAlitripIeAgentRefundGetAPIRequest对象
func NewTaobaoAlitripIeAgentRefundGetRequest() *TaobaoAlitripIeAgentRefundGetAPIRequest {
	return &TaobaoAlitripIeAgentRefundGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.refund.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ApplyId Setter
// 退票申请单id
func (r *TaobaoAlitripIeAgentRefundGetAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// Get ApplyId Getter
func (r TaobaoAlitripIeAgentRefundGetAPIRequest) GetApplyId() int64 {
	return r._applyId
}

// Set is AgentId Setter
// 代理商id
func (r *TaobaoAlitripIeAgentRefundGetAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// Get AgentId Getter
func (r TaobaoAlitripIeAgentRefundGetAPIRequest) GetAgentId() int64 {
	return r._agentId
}
