package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】酒店订单查询 API请求
alitrip.btrip.supplychain.hotel.search

【商旅】酒店订单查询
*/
type AlitripBtripSupplychainHotelSearchRequest struct {
    model.Params
    // 入参
    _rq   *OpenApiSearchRq
}

// 初始化AlitripBtripSupplychainHotelSearchRequest对象
func NewAlitripBtripSupplychainHotelSearchRequest() *AlitripBtripSupplychainHotelSearchRequest{
    return &AlitripBtripSupplychainHotelSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainHotelSearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.supplychain.hotel.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainHotelSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参
func (r *AlitripBtripSupplychainHotelSearchRequest) SetRq(_rq *OpenApiSearchRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripSupplychainHotelSearchRequest) GetRq() *OpenApiSearchRq {
    return r._rq
}
