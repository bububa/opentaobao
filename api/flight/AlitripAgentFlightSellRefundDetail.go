package flight

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/flight"
)

/* 
销售退票单详情 
alitrip.agent.flight.sell.refund.detail

销售退票单详情
*/
func AlitripAgentFlightSellRefundDetail(clt *core.SDKClient, req *flight.AlitripAgentFlightSellRefundDetailRequest, session string) (*flight.AlitripAgentFlightSellRefundDetailAPIResponse, error) {
    var resp flight.AlitripAgentFlightSellRefundDetailAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
