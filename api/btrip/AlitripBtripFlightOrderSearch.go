package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
获取机票订单列表 
alitrip.btrip.flight.order.search

第三方OA厂商获取机票订单列表
*/
func AlitripBtripFlightOrderSearch(clt *core.SDKClient, req *btrip.AlitripBtripFlightOrderSearchRequest, session string) (*btrip.AlitripBtripFlightOrderSearchAPIResponse, error) {
    var resp btrip.AlitripBtripFlightOrderSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
