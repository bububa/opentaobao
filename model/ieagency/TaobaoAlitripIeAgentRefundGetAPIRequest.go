package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripieagentrefundgetAPIRequest 获取退票申请详情 API请求
// taobao.alitrip.ie.agent.refund.get
//
// 获取退票申请详情
type TaobaoalitripieagentrefundgetAPIRequest struct {
	model.Params
	// 退票申请单id
	_applyId int64
	// 代理商id
	_agentId int64
}

// NewTaobaoalitripieagentrefundgetRequest 初始化TaobaoalitripieagentrefundgetAPIRequest对象
func NewTaobaoalitripieagentrefundgetRequest() *TaobaoalitripieagentrefundgetAPIRequest {
	return &TaobaoalitripieagentrefundgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripieagentrefundgetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.refund.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripieagentrefundgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripieagentrefundgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyId is ApplyId Setter
// 退票申请单id
func (r *TaobaoalitripieagentrefundgetAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoalitripieagentrefundgetAPIRequest) GetApplyId() int64 {
	return r._applyId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoalitripieagentrefundgetAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoalitripieagentrefundgetAPIRequest) GetAgentId() int64 {
	return r._agentId
}
