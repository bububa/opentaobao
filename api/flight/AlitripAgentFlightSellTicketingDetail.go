package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightSellTicketingDetail 销售出票详情
// alitrip.agent.flight.sell.ticketing.detail
//
// 销售出票详情
func AlitripAgentFlightSellTicketingDetail(clt *core.SDKClient, req *flight.AlitripAgentFlightSellTicketingDetailAPIRequest, resp *flight.AlitripAgentFlightSellTicketingDetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
