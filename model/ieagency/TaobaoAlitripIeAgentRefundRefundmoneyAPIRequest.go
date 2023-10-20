package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripieagentrefundrefundmoneyAPIRequest 确认退款 API请求
// taobao.alitrip.ie.agent.refund.refundmoney
//
// 卖家进行退款操作
type TaobaoalitripieagentrefundrefundmoneyAPIRequest struct {
	model.Params
	// 退票申请单id
	_applyId int64
	// 代理商id
	_agentId int64
}

// NewTaobaoalitripieagentrefundrefundmoneyRequest 初始化TaobaoalitripieagentrefundrefundmoneyAPIRequest对象
func NewTaobaoalitripieagentrefundrefundmoneyRequest() *TaobaoalitripieagentrefundrefundmoneyAPIRequest {
	return &TaobaoalitripieagentrefundrefundmoneyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripieagentrefundrefundmoneyAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.refund.refundmoney"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripieagentrefundrefundmoneyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripieagentrefundrefundmoneyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyId is ApplyId Setter
// 退票申请单id
func (r *TaobaoalitripieagentrefundrefundmoneyAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoalitripieagentrefundrefundmoneyAPIRequest) GetApplyId() int64 {
	return r._applyId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoalitripieagentrefundrefundmoneyAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoalitripieagentrefundrefundmoneyAPIRequest) GetAgentId() int64 {
	return r._agentId
}
