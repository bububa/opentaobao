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
type AlitripBtripSupplychainHotelSearchAPIRequest struct {
    model.Params
    // 入参
    _rq   *OpenApiSearchRq
}

// 初始化AlitripBtripSupplychainHotelSearchAPIRequest对象
func NewAlitripBtripSupplychainHotelSearchRequest() *AlitripBtripSupplychainHotelSearchAPIRequest{
    return &AlitripBtripSupplychainHotelSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainHotelSearchAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.supplychain.hotel.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainHotelSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参
func (r *AlitripBtripSupplychainHotelSearchAPIRequest) SetRq(_rq *OpenApiSearchRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripSupplychainHotelSearchAPIRequest) GetRq() *OpenApiSearchRq {
    return r._rq
}
