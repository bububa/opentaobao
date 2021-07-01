package ieagency

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ieagency"
)

/* 
【国际机票】查询辅营订单列表 
alitrip.tripvp.agent.order.search

【国际机票】查询辅营订单列表
*/
func AlitripTripvpAgentOrderSearch(clt *core.SDKClient, req *ieagency.AlitripTripvpAgentOrderSearchAPIRequest, session string) (*ieagency.AlitripTripvpAgentOrderSearchAPIResponse, error) {
    var resp ieagency.AlitripTripvpAgentOrderSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
