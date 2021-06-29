package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方市内用车申请单查询 APIRequest
alitrip.btrip.city.car.apply.query

三方市内用车申请单查询
*/
type AlitripBtripCityCarApplyQueryRequest struct {
    model.Params

    // 入参对象
    rq   *CityCarApplyQueryRq 

}

func NewAlitripBtripCityCarApplyQueryRequest() *AlitripBtripCityCarApplyQueryRequest{
    return &AlitripBtripCityCarApplyQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripCityCarApplyQueryRequest) GetApiMethodName() string {
    return "alitrip.btrip.city.car.apply.query"
}

func (r AlitripBtripCityCarApplyQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripCityCarApplyQueryRequest) SetRq(rq *CityCarApplyQueryRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripCityCarApplyQueryRequest) GetRq() *CityCarApplyQueryRq {
    return r.rq
}

