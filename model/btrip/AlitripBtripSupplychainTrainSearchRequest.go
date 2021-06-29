package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】火车票订单查询 APIRequest
alitrip.btrip.supplychain.train.search

【商旅】火车票订单查询
*/
type AlitripBtripSupplychainTrainSearchRequest struct {
    model.Params

    // 入参
    rq   *OpenApiSearchRq 

}

func NewAlitripBtripSupplychainTrainSearchRequest() *AlitripBtripSupplychainTrainSearchRequest{
    return &AlitripBtripSupplychainTrainSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripSupplychainTrainSearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.supplychain.train.search"
}

func (r AlitripBtripSupplychainTrainSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripSupplychainTrainSearchRequest) SetRq(rq *OpenApiSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripSupplychainTrainSearchRequest) GetRq() *OpenApiSearchRq {
    return r.rq
}

