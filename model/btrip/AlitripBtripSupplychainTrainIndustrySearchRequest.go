package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票行业搜索接口 APIRequest
alitrip.btrip.supplychain.train.industry.search

【商旅】火车票行业搜索接口
*/
type AlitripBtripSupplychainTrainIndustrySearchRequest struct {
    model.Params

    // 入参
    rq   *TrainSearchRq 

}

func NewAlitripBtripSupplychainTrainIndustrySearchRequest() *AlitripBtripSupplychainTrainIndustrySearchRequest{
    return &AlitripBtripSupplychainTrainIndustrySearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripSupplychainTrainIndustrySearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.supplychain.train.industry.search"
}

func (r AlitripBtripSupplychainTrainIndustrySearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripSupplychainTrainIndustrySearchRequest) SetRq(rq *TrainSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripSupplychainTrainIndustrySearchRequest) GetRq() *TrainSearchRq {
    return r.rq
}

