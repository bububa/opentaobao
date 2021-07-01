package flight

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/flight"
)

/* 
销售退票拒绝 
alitrip.agent.flight.sell.refund.refuse

销售退票拒绝
*/
func AlitripAgentFlightSellRefundRefuse(clt *core.SDKClient, req *flight.AlitripAgentFlightSellRefundRefuseAPIRequest, session string) (*flight.AlitripAgentFlightSellRefundRefuseAPIResponse, error) {
    var resp flight.AlitripAgentFlightSellRefundRefuseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
