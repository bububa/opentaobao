package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripIeAgentTicketIssueAPIRequest
【国际机票】手工出票 API请求
taobao.alitrip.ie.agent.ticket.issue

代理商手工出票，并回填票号 */
type TaobaoAlitripIeAgentTicketIssueAPIRequest struct {
	model.Params
	// 代理商id
	_agentId int64
	// 出票信息
	_issueTicketVO *IeIssueTicketVO
}

// New
