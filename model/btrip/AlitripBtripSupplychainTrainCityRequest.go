package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
火车站数据查询 APIRequest
alitrip.btrip.supplychain.train.city

火车站数据查询
*/
type AlitripBtripSupplychainTrainCityRequest struct {
    model.Params

    // 入参对象
    rq   *OpenSuggestRq 

}

func NewAlitripBtripSupplychainTrainCityRequest() *AlitripBtripSupplychainTrainCityRequest{
    return &AlitripBtripSupplychainTrainCityRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripSupplychainTrainCityRequest) GetApiMethodName() string {
    return "alitrip.btrip.supplychain.train.city"
}

func (r AlitripBtripSupplychainTrainCityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripSupplychainTrainCityRequest) SetRq(rq *OpenSuggestRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripSupplychainTrainCityRequest) GetRq() *OpenSuggestRq {
    return r.rq
}

