package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方市内用车申请单同步 API请求
alitrip.btrip.city.car.apply.add

三方市内用车申请单同步
*/
type AlitripBtripCityCarApplyAddRequest struct {
    model.Params
    // 入参对象
    _rq   *CityCarApplyAddRq
}

// 初始化AlitripBtripCityCarApplyAddRequest对象
func NewAlitripBtripCityCarApplyAddRequest() *AlitripBtripCityCarApplyAddRequest{
    return &AlitripBtripCityCarApplyAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCityCarApplyAddRequest) GetApiMethodName() string {
    return "alitrip.btrip.city.car.apply.add"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCityCarApplyAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripCityCarApplyAddRequest) SetRq(_rq *CityCarApplyAddRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCityCarApplyAddRequest) GetRq() *CityCarApplyAddRq {
    return r._rq
}
