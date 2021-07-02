package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentRefundRefundmoneyAPIRequest 确认退款 API请求
// taobao.alitrip.ie.agent.refund.refundmoney
//
// 卖家进行退款操作
type TaobaoAlitripIeAgentRefundRefundmoneyAPIRequest struct {
	model.Params
	// 退票申请单id
	_applyId int64
	// 代理商id
	_agentId int64
}

// NewTaobaoAlitripIeAgentRefundRefundmoneyRequest 初始化TaobaoAlitripIeAgentRefundRefundmoneyAPIRequest对象
func NewTaobaoAlitripIeAgentRefundRefundmoneyRequest() *TaobaoAlitripIeAgentRefundRefundmoneyAPIRequest {
	return &TaobaoAlitripIeAgentRefundRefundmoneyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundRefundmoneyAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.refund.refundmoney"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundRefundmoneyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetApplyId is ApplyId Setter
// 退票申请单id
func (r *TaobaoAlitripIeAgentRefundRefundmoneyAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoAlitripIeAgentRefundRefundmoneyAPIRequest) GetApplyId() int64 {
	return r._applyId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoAlitripIeAgentRefundRefundmoneyAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoAlitripIeAgentRefundRefundmoneyAPIRequest) GetAgentId() int64 {
	return r._agentId
}
