package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票行业搜索接口 APIRequest
alitrip.btrip.supplychain.bus.industry.search

汽车票行业搜索接口
*/
type AlitripBtripSupplychainBusIndustrySearchRequest struct {
    model.Params

    // 入参
    rq   *BusSearchRq 

}

func NewAlitripBtripSupplychainBusIndustrySearchRequest() *AlitripBtripSupplychainBusIndustrySearchRequest{
    return &AlitripBtripSupplychainBusIndustrySearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripSupplychainBusIndustrySearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.supplychain.bus.industry.search"
}

func (r AlitripBtripSupplychainBusIndustrySearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripSupplychainBusIndustrySearchRequest) SetRq(rq *BusSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripSupplychainBusIndustrySearchRequest) GetRq() *BusSearchRq {
    return r.rq
}

