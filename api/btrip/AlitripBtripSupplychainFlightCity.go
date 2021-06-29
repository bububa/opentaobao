package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
机场数据查询 
alitrip.btrip.supplychain.flight.city

机场数据查询
*/
func AlitripBtripSupplychainFlightCity(clt *core.SDKClient, req *btrip.AlitripBtripSupplychainFlightCityRequest, session string) (*btrip.AlitripBtripSupplychainFlightCityAPIResponse, error) {
    var resp btrip.AlitripBtripSupplychainFlightCityAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
