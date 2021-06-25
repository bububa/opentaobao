package flight

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/flight"
)

/* 
销售退票确认 
alitrip.agent.flight.sell.refund.approve

销售退票确认
*/
func AlitripAgentFlightSellRefundApprove(clt *core.SDKClient, req *flight.AlitripAgentFlightSellRefundApproveRequest, session string) (*flight.AlitripAgentFlightSellRefundApproveResponse, error) {
    var resp flight.AlitripAgentFlightSellRefundApproveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
