package flight

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/flight"
)

/* 
销售退票单列表 
alitrip.agent.flight.sell.refund.list

销售退票单列表
*/
func AlitripAgentFlightSellRefundList(clt *core.SDKClient, req *flight.AlitripAgentFlightSellRefundListAPIRequest, session string) (*flight.AlitripAgentFlightSellRefundListAPIResponse, error) {
    var resp flight.AlitripAgentFlightSellRefundListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
