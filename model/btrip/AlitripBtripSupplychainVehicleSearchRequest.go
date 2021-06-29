package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】用车订单搜索 API请求
alitrip.btrip.supplychain.vehicle.search

【商旅】用车订单搜索
*/
type AlitripBtripSupplychainVehicleSearchRequest struct {
    model.Params
    // 出参
    _rq   *OpenApiSearchRq
}

// 初始化AlitripBtripSupplychainVehicleSearchRequest对象
func NewAlitripBtripSupplychainVehicleSearchRequest() *AlitripBtripSupplychainVehicleSearchRequest{
    return &AlitripBtripSupplychainVehicleSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainVehicleSearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.supplychain.vehicle.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainVehicleSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 出参
func (r *AlitripBtripSupplychainVehicleSearchRequest) SetRq(_rq *OpenApiSearchRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripSupplychainVehicleSearchRequest) GetRq() *OpenApiSearchRq {
    return r._rq
}
