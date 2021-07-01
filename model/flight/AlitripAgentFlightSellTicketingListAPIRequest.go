package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripAgentFlightSellTicketingListAPIRequest
销售出票列表 API请求
alitrip.agent.flight.sell.ticketing.list

销售出票列表 */
type AlitripAgentFlightSellTicketingListAPIRequest struct {
	model.Params
	// 入参对象
	_param *TicketingListRequestDto
}

// New
