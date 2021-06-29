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
type AlitripBtripOpenSupplychainVehicleTradeRequest struct {
    model.Params
    // 入参
    rq   *OpenApiZzdSearchRq
}

// 初始化AlitripBtripOpenSupplychainVehicleTradeRequest对象
func NewAlitripBtripOpenSupplychainVehicleTradeRequest() *AlitripBtripOpenSupplychainVehicleTradeRequest{
    return &AlitripBtripOpenSupplychainVehicleTradeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenSupplychainVehicleTradeRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.supplychain.vehicle.trade"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenSupplychainVehicleTradeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参
func (r *AlitripBtripOpenSupplychainVehicleTradeRequest) SetRq(rq *OpenApiZzdSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenSupplychainVehicleTradeRequest) GetRq() *OpenApiZzdSearchRq {
    return r.rq
}
