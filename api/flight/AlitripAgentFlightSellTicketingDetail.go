package flight

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/flight"
)

/* 
销售出票详情 
alitrip.agent.flight.sell.ticketing.detail

销售出票详情
*/
func AlitripAgentFlightSellTicketingDetail(clt *core.SDKClient, req *flight.AlitripAgentFlightSellTicketingDetailAPIRequest, session string) (*flight.AlitripAgentFlightSellTicketingDetailAPIResponse, error) {
    var resp flight.AlitripAgentFlightSellTicketingDetailAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
