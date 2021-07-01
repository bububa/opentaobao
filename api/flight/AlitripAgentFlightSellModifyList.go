package flight

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/flight"
)

/* 
销售改签单列表 
alitrip.agent.flight.sell.modify.list

销售改签单列表
*/
func AlitripAgentFlightSellModifyList(clt *core.SDKClient, req *flight.AlitripAgentFlightSellModifyListAPIRequest, session string) (*flight.AlitripAgentFlightSellModifyListAPIResponse, error) {
    var resp flight.AlitripAgentFlightSellModifyListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
