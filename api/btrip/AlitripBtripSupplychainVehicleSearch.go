package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
【商旅】用车订单搜索 
alitrip.btrip.supplychain.vehicle.search

【商旅】用车订单搜索
*/
func AlitripBtripSupplychainVehicleSearch(clt *core.SDKClient, req *btrip.AlitripBtripSupplychainVehicleSearchRequest, session string) (*btrip.AlitripBtripSupplychainVehicleSearchAPIResponse, error) {
    var resp btrip.AlitripBtripSupplychainVehicleSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
