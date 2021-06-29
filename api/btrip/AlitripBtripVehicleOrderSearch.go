package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
用车订单查询接口 
alitrip.btrip.vehicle.order.search

企业获取商旅用车订单数据
*/
func AlitripBtripVehicleOrderSearch(clt *core.SDKClient, req *btrip.AlitripBtripVehicleOrderSearchRequest, session string) (*btrip.AlitripBtripVehicleOrderSearchAPIResponse, error) {
    var resp btrip.AlitripBtripVehicleOrderSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
