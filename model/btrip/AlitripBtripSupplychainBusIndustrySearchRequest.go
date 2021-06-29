package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票行业搜索接口 API请求
alitrip.btrip.supplychain.bus.industry.search

汽车票行业搜索接口
*/
type AlitripBtripSupplychainBusIndustrySearchRequest struct {
    model.Params
    // 入参
    _rq   *BusSearchRq
}

// 初始化AlitripBtripSupplychainBusIndustrySearchRequest对象
func NewAlitripBtripSupplychainBusIndustrySearchRequest() *AlitripBtripSupplychainBusIndustrySearchRequest{
    return &AlitripBtripSupplychainBusIndustrySearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainBusIndustrySearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.supplychain.bus.industry.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainBusIndustrySearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参
func (r *AlitripBtripSupplychainBusIndustrySearchRequest) SetRq(_rq *BusSearchRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripSupplychainBusIndustrySearchRequest) GetRq() *BusSearchRq {
    return r._rq
}
