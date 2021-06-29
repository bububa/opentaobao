package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
机场数据查询 API请求
alitrip.btrip.supplychain.flight.city

机场数据查询
*/
type AlitripBtripSupplychainFlightCityRequest struct {
    model.Params
    // 请求对象
    rq   *OpenSuggestRq
}

// 初始化AlitripBtripSupplychainFlightCityRequest对象
func NewAlitripBtripSupplychainFlightCityRequest() *AlitripBtripSupplychainFlightCityRequest{
    return &AlitripBtripSupplychainFlightCityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainFlightCityRequest) GetApiMethodName() string {
    return "alitrip.btrip.supplychain.flight.city"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainFlightCityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripSupplychainFlightCityRequest) SetRq(rq *OpenSuggestRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripSupplychainFlightCityRequest) GetRq() *OpenSuggestRq {
    return r.rq
}
