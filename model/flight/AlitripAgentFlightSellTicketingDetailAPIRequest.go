package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripAgentFlightSellTicketingDetailAPIRequest
销售出票详情 API请求
alitrip.agent.flight.sell.ticketing.detail

销售出票详情 */
type AlitripAgentFlightSellTicketingDetailAPIRequest struct {
	model.Params
	// 国内国际标识
	_domesticIntl int64
	// 飞猪订单号
	_orderId string
}

// New
