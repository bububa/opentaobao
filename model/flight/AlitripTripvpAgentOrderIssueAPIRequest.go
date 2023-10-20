package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptripvpagentorderissueAPIRequest 廉航辅营正向订单出货接口 API请求
// alitrip.tripvp.agent.order.issue
//
// 廉航辅营正向订单出货接口
type AlitriptripvpagentorderissueAPIRequest struct {
	model.Params
	// 代理商ID
	_agentId int64
	// 出货通知信息
	_issueProOrderVo *IssueProOrderVo
}

// NewAlitriptripvpagentorderissueRequest 初始化AlitriptripvpagentorderissueAPIRequest对象
func NewAlitriptripvpagentorderissueRequest() *AlitriptripvpagentorderissueAPIRequest {
	return &AlitriptripvpagentorderissueAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptripvpagentorderissueAPIRequest) GetApiMethodName() string {
	return "alitrip.tripvp.agent.order.issue"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptripvpagentorderissueAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptripvpagentorderissueAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentId is AgentId Setter
// 代理商ID
func (r *AlitriptripvpagentorderissueAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r AlitriptripvpagentorderissueAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetIssueProOrderVo is IssueProOrderVo Setter
// 出货通知信息
func (r *AlitriptripvpagentorderissueAPIRequest) SetIssueProOrderVo(_issueProOrderVo *IssueProOrderVo) error {
	r._issueProOrderVo = _issueProOrderVo
	r.Set("issue_pro_order_vo", _issueProOrderVo)
	return nil
}

// GetIssueProOrderVo IssueProOrderVo Getter
func (r AlitriptripvpagentorderissueAPIRequest) GetIssueProOrderVo() *IssueProOrderVo {
	return r._issueProOrderVo
}
