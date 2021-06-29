package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】机票订单查询 APIRequest
alitrip.btrip.supplychain.flight.search

【商旅】机票订单查询
*/
type AlitripBtripSupplychainFlightSearchRequest struct {
    model.Params

    // 请求对象
    rq   *OpenApiSearchRq 

}

func NewAlitripBtripSupplychainFlightSearchRequest() *AlitripBtripSupplychainFlightSearchRequest{
    return &AlitripBtripSupplychainFlightSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripSupplychainFlightSearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.supplychain.flight.search"
}

func (r AlitripBtripSupplychainFlightSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripSupplychainFlightSearchRequest) SetRq(rq *OpenApiSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripSupplychainFlightSearchRequest) GetRq() *OpenApiSearchRq {
    return r.rq
}

