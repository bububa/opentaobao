package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripAgentFlightSellTicketingIssueAPIRequest
销售出票 API请求
alitrip.agent.flight.sell.ticketing.issue

销售出票 */
type AlitripAgentFlightSellTicketingIssueAPIRequest struct {
	model.Params
	// 入参对象
	_param *TicketingIssueRequestDto
}

// New
