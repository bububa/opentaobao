package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
商旅酒店api分销-详情报价接口 
alitrip.btrip.hotel.distribution.search.detail

商旅酒店api分销-详情报价接口
*/
func AlitripBtripHotelDistributionSearchDetail(clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionSearchDetailAPIRequest, session string) (*btrip.AlitripBtripHotelDistributionSearchDetailAPIResponse, error) {
    var resp btrip.AlitripBtripHotelDistributionSearchDetailAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
