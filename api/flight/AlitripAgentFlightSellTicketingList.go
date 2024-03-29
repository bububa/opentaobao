package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightSellTicketingList 销售出票列表
// alitrip.agent.flight.sell.ticketing.list
//
// 销售出票列表
func AlitripAgentFlightSellTicketingList(clt *core.SDKClient, req *flight.AlitripAgentFlightSellTicketingListAPIRequest, resp *flight.AlitripAgentFlightSellTicketingListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
