package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
机票行业搜索接口 APIRequest
alitrip.btrip.supplychain.flight.industry.search

【商旅】机票行业搜索
*/
type AlitripBtripSupplychainFlightIndustrySearchRequest struct {
    model.Params

    // 入参
    rq   *FlightSearchRq 

}

func NewAlitripBtripSupplychainFlightIndustrySearchRequest() *AlitripBtripSupplychainFlightIndustrySearchRequest{
    return &AlitripBtripSupplychainFlightIndustrySearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripSupplychainFlightIndustrySearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.supplychain.flight.industry.search"
}

func (r AlitripBtripSupplychainFlightIndustrySearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripSupplychainFlightIndustrySearchRequest) SetRq(rq *FlightSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripSupplychainFlightIndustrySearchRequest) GetRq() *FlightSearchRq {
    return r.rq
}

