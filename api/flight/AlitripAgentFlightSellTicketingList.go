package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightSellTicketingList 销售出票列表
// alitrip.agent.flight.sell.ticketing.list
//
// 销售出票列表
func AlitripAgentFlightSellTicketingList(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentFlightSellTicketingListAPIRequest, resp *flight.AlitripAgentFlightSellTicketingListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
