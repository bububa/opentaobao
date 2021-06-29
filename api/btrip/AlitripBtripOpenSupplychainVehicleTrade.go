package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
商旅用车交易流水接口 
alitrip.btrip.open.supplychain.vehicle.trade

商旅用车交易流水接口
*/
func AlitripBtripOpenSupplychainVehicleTrade(clt *core.SDKClient, req *btrip.AlitripBtripOpenSupplychainVehicleTradeRequest, session string) (*btrip.AlitripBtripOpenSupplychainVehicleTradeAPIResponse, error) {
    var resp btrip.AlitripBtripOpenSupplychainVehicleTradeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
