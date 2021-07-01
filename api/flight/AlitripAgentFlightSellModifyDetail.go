package flight

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/flight"
)

/* 
销售改签详情 
alitrip.agent.flight.sell.modify.detail

销售改签详情
*/
func AlitripAgentFlightSellModifyDetail(clt *core.SDKClient, req *flight.AlitripAgentFlightSellModifyDetailAPIRequest, session string) (*flight.AlitripAgentFlightSellModifyDetailAPIResponse, error) {
    var resp flight.AlitripAgentFlightSellModifyDetailAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
