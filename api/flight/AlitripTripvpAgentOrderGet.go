package flight

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/flight"
)

/* 
廉航辅营正向订单查询详情接口 
alitrip.tripvp.agent.order.get

【国际机票】查询辅营订单详情
*/
func AlitripTripvpAgentOrderGet(clt *core.SDKClient, req *flight.AlitripTripvpAgentOrderGetAPIRequest, session string) (*flight.AlitripTripvpAgentOrderGetAPIResponse, error) {
    var resp flight.AlitripTripvpAgentOrderGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
