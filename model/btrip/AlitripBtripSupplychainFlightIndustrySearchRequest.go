package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
机票行业搜索接口 API请求
alitrip.btrip.supplychain.flight.industry.search

【商旅】机票行业搜索
*/
type AlitripBtripSupplychainFlightIndustrySearchRequest struct {
    model.Params
    // 入参
    _rq   *FlightSearchRq
}

// 初始化AlitripBtripSupplychainFlightIndustrySearchRequest对象
func NewAlitripBtripSupplychainFlightIndustrySearchRequest() *AlitripBtripSupplychainFlightIndustrySearchRequest{
    return &AlitripBtripSupplychainFlightIndustrySearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainFlightIndustrySearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.supplychain.flight.industry.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainFlightIndustrySearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参
func (r *AlitripBtripSupplychainFlightIndustrySearchRequest) SetRq(_rq *FlightSearchRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripSupplychainFlightIndustrySearchRequest) GetRq() *FlightSearchRq {
    return r._rq
}
