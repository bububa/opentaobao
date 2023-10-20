package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentTicketIssueAPIRequest 【国际机票】手工出票 API请求
// taobao.alitrip.ie.agent.ticket.issue
//
// 代理商手工出票，并回填票号
type TaobaoAlitripIeAgentTicketIssueAPIRequest struct {
	model.Params
	// 代理商id
	_agentId int64
	// 出票信息
	_issueTicketVO *IeIssueTicketVo
}

// NewTaobaoAlitripIeAgentTicketIssueRequest 初始化TaobaoAlitripIeAgentTicketIssueAPIRequest对象
func NewTaobaoAlitripIeAgentTicketIssueRequest() *TaobaoAlitripIeAgentTicketIssueAPIRequest {
	return &TaobaoAlitripIeAgentTicketIssueAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentTicketIssueAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.ticket.issue"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentTicketIssueAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripIeAgentTicketIssueAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoAlitripIeAgentTicketIssueAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoAlitripIeAgentTicketIssueAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetIssueTicketVO is IssueTicketVO Setter
// 出票信息
func (r *TaobaoAlitripIeAgentTicketIssueAPIRequest) SetIssueTicketVO(_issueTicketVO *IeIssueTicketVo) error {
	r._issueTicketVO = _issueTicketVO
	r.Set("issue_ticket_v_o", _issueTicketVO)
	return nil
}

// GetIssueTicketVO IssueTicketVO Getter
func (r TaobaoAlitripIeAgentTicketIssueAPIRequest) GetIssueTicketVO() *IeIssueTicketVo {
	return r._issueTicketVO
}
