package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
机场数据查询 APIRequest
alitrip.btrip.supplychain.flight.city

机场数据查询
*/
type AlitripBtripSupplychainFlightCityRequest struct {
    model.Params

    // 请求对象
    rq   *OpenSuggestRq 

}

func NewAlitripBtripSupplychainFlightCityRequest() *AlitripBtripSupplychainFlightCityRequest{
    return &AlitripBtripSupplychainFlightCityRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripSupplychainFlightCityRequest) GetApiMethodName() string {
    return "alitrip.btrip.supplychain.flight.city"
}

func (r AlitripBtripSupplychainFlightCityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripSupplychainFlightCityRequest) SetRq(rq *OpenSuggestRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripSupplychainFlightCityRequest) GetRq() *OpenSuggestRq {
    return r.rq
}

