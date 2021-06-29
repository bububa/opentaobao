package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅用车交易流水接口 APIRequest
alitrip.btrip.open.supplychain.vehicle.trade

商旅用车交易流水接口
*/
type AlitripBtripOpenSupplychainVehicleTradeRequest struct {
    model.Params

    // 入参
    rq   *OpenApiZzdSearchRq 

}

func NewAlitripBtripOpenSupplychainVehicleTradeRequest() *AlitripBtripOpenSupplychainVehicleTradeRequest{
    return &AlitripBtripOpenSupplychainVehicleTradeRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripOpenSupplychainVehicleTradeRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.supplychain.vehicle.trade"
}

func (r AlitripBtripOpenSupplychainVehicleTradeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripOpenSupplychainVehicleTradeRequest) SetRq(rq *OpenApiZzdSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripOpenSupplychainVehicleTradeRequest) GetRq() *OpenApiZzdSearchRq {
    return r.rq
}

