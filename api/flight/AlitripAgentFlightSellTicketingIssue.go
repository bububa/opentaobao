package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

/* AlitripAgentFlightSellTicketingIssue
销售出票
alitrip.agent.flight.sell.ticketing.issue

销售出票 */
func AlitripAgentFlightSellTicketingIssue(clt *core.SDKClient, req *flight.AlitripAgentFlightSellTicketingIssueAPIRequest, session string) (*flight.AlitripAgentFlightSellTicketingIssueAPIResponse, error) {
	var resp flight.AlitripAgentFlightSellTicketingIssueAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
