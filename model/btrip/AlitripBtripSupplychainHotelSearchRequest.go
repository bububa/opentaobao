package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】酒店订单查询 APIRequest
alitrip.btrip.supplychain.hotel.search

【商旅】酒店订单查询
*/
type AlitripBtripSupplychainHotelSearchRequest struct {
    model.Params

    // 入参
    rq   *OpenApiSearchRq 

}

func NewAlitripBtripSupplychainHotelSearchRequest() *AlitripBtripSupplychainHotelSearchRequest{
    return &AlitripBtripSupplychainHotelSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripSupplychainHotelSearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.supplychain.hotel.search"
}

func (r AlitripBtripSupplychainHotelSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripSupplychainHotelSearchRequest) SetRq(rq *OpenApiSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripSupplychainHotelSearchRequest) GetRq() *OpenApiSearchRq {
    return r.rq
}

