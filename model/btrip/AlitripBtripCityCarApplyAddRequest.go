package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方市内用车申请单同步 APIRequest
alitrip.btrip.city.car.apply.add

三方市内用车申请单同步
*/
type AlitripBtripCityCarApplyAddRequest struct {
    model.Params

    // 入参对象
    rq   *CityCarApplyAddRq 

}

func NewAlitripBtripCityCarApplyAddRequest() *AlitripBtripCityCarApplyAddRequest{
    return &AlitripBtripCityCarApplyAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripCityCarApplyAddRequest) GetApiMethodName() string {
    return "alitrip.btrip.city.car.apply.add"
}

func (r AlitripBtripCityCarApplyAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripCityCarApplyAddRequest) SetRq(rq *CityCarApplyAddRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripCityCarApplyAddRequest) GetRq() *CityCarApplyAddRq {
    return r.rq
}

