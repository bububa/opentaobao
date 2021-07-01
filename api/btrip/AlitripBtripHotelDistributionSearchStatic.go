package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
商旅酒店api分销-酒店静态信息接口 
alitrip.btrip.hotel.distribution.search.static

商旅酒店api分销-酒店静态信息接口
*/
func AlitripBtripHotelDistributionSearchStatic(clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionSearchStaticAPIRequest, session string) (*btrip.AlitripBtripHotelDistributionSearchStaticAPIResponse, error) {
    var resp btrip.AlitripBtripHotelDistributionSearchStaticAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
