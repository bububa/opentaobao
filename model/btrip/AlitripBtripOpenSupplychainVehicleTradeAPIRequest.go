package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅用车交易流水接口 API请求
alitrip.btrip.open.supplychain.vehicle.trade

商旅用车交易流水接口
*/
type AlitripBtripOpenSupplychainVehicleTradeAPIRequest struct {
    model.Params
    // 入参
    _rq   *OpenApiZzdSearchRq
}

// 初始化AlitripBtripOpenSupplychainVehicleTradeAPIRequest对象
func NewAlitripBtripOpenSupplychainVehicleTradeRequest() *AlitripBtripOpenSupplychainVehicleTradeAPIRequest{
    return &AlitripBtripOpenSupplychainVehicleTradeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenSupplychainVehicleTradeAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.supplychain.vehicle.trade"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenSupplychainVehicleTradeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参
func (r *AlitripBtripOpenSupplychainVehicleTradeAPIRequest) SetRq(_rq *OpenApiZzdSearchRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenSupplychainVehicleTradeAPIRequest) GetRq() *OpenApiZzdSearchRq {
    return r._rq
}
