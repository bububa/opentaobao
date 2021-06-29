package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方市内用车申请单审批 APIRequest
alitrip.btrip.city.car.apply.approve

三方市内用车申请单审批
*/
type AlitripBtripCityCarApplyApproveRequest struct {
    model.Params

    // 入参对象
    rq   *CityCarApplyApproveRq 

}

func NewAlitripBtripCityCarApplyApproveRequest() *AlitripBtripCityCarApplyApproveRequest{
    return &AlitripBtripCityCarApplyApproveRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripCityCarApplyApproveRequest) GetApiMethodName() string {
    return "alitrip.btrip.city.car.apply.approve"
}

func (r AlitripBtripCityCarApplyApproveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripCityCarApplyApproveRequest) SetRq(rq *CityCarApplyApproveRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripCityCarApplyApproveRequest) GetRq() *CityCarApplyApproveRq {
    return r.rq
}

