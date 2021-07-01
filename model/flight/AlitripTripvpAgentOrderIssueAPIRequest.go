package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTripvpAgentOrderIssueAPIRequest
廉航辅营正向订单出货接口 API请求
alitrip.tripvp.agent.order.issue

廉航辅营正向订单出货接口 */
type AlitripTripvpAgentOrderIssueAPIRequest struct {
	model.Params
	// 代理商ID
	_agentId int64
	// 出货通知信息
	_issueProOrderVo *IssueProOrderVo
}

// NewAlitripTripvpAgentOrderIssueRequest 初始化AlitripTripvpAgentOrderIssueAPIRequest对象
func NewAlitripTripvpAgentOrderIssueRequest() *AlitripTripvpAgentOrderIssueAPIRequest {
	return &AlitripTripvpAgentOrderIssueAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTripvpAgentOrderIssueAPIRequest) GetApiMethodName() string {
	return "alitrip.tripvp.agent.order.issue"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTripvpAgentOrderIssueAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AgentId Setter
// 代理商ID
func (r *AlitripTripvpAgentOrderIssueAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// Get AgentId Getter
func (r AlitripTripvpAgentOrderIssueAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// Set is IssueProOrderVo Setter
// 出货通知信息
func (r *AlitripTripvpAgentOrderIssueAPIRequest) SetIssueProOrderVo(_issueProOrderVo *IssueProOrderVo) error {
	r._issueProOrderVo = _issueProOrderVo
	r.Set("issue_pro_order_vo", _issueProOrderVo)
	return nil
}

// Get IssueProOrderVo Getter
func (r AlitripTripvpAgentOrderIssueAPIRequest) GetIssueProOrderVo() *IssueProOrderVo {
	return r._issueProOrderVo
}
