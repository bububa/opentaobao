package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
【商旅】酒店订单查询 
alitrip.btrip.supplychain.hotel.search

【商旅】酒店订单查询
*/
func AlitripBtripSupplychainHotelSearch(clt *core.SDKClient, req *btrip.AlitripBtripSupplychainHotelSearchAPIRequest, session string) (*btrip.AlitripBtripSupplychainHotelSearchAPIResponse, error) {
    var resp btrip.AlitripBtripSupplychainHotelSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
