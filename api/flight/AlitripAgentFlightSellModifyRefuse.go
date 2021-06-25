package flight

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/flight"
)

/* 
销售改签拒绝 
alitrip.agent.flight.sell.modify.refuse

销售改签拒绝
*/
func AlitripAgentFlightSellModifyRefuse(clt *core.SDKClient, req *flight.AlitripAgentFlightSellModifyRefuseRequest, session string) (*flight.AlitripAgentFlightSellModifyRefuseResponse, error) {
    var resp flight.AlitripAgentFlightSellModifyRefuseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
