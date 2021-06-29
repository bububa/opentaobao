package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
商旅酒店api分销-酒店最低价 
alitrip.btrip.hotel.distribution.search.low.price

商旅酒店api分销-酒店最低价
*/
func AlitripBtripHotelDistributionSearchLowPrice(clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionSearchLowPriceRequest, session string) (*btrip.AlitripBtripHotelDistributionSearchLowPriceAPIResponse, error) {
    var resp btrip.AlitripBtripHotelDistributionSearchLowPriceAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
