package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方市内用车申请单查询 API请求
alitrip.btrip.city.car.apply.query

三方市内用车申请单查询
*/
type AlitripBtripCityCarApplyQueryRequest struct {
    model.Params
    // 入参对象
    _rq   *CityCarApplyQueryRq
}

// 初始化AlitripBtripCityCarApplyQueryRequest对象
func NewAlitripBtripCityCarApplyQueryRequest() *AlitripBtripCityCarApplyQueryRequest{
    return &AlitripBtripCityCarApplyQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCityCarApplyQueryRequest) GetApiMethodName() string {
    return "alitrip.btrip.city.car.apply.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCityCarApplyQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripCityCarApplyQueryRequest) SetRq(_rq *CityCarApplyQueryRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCityCarApplyQueryRequest) GetRq() *CityCarApplyQueryRq {
    return r._rq
}
