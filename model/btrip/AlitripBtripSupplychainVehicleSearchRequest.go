package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】用车订单搜索 APIRequest
alitrip.btrip.supplychain.vehicle.search

【商旅】用车订单搜索
*/
type AlitripBtripSupplychainVehicleSearchRequest struct {
    model.Params

    // 出参
    rq   *OpenApiSearchRq 

}

func NewAlitripBtripSupplychainVehicleSearchRequest() *AlitripBtripSupplychainVehicleSearchRequest{
    return &AlitripBtripSupplychainVehicleSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripSupplychainVehicleSearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.supplychain.vehicle.search"
}

func (r AlitripBtripSupplychainVehicleSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripSupplychainVehicleSearchRequest) SetRq(rq *OpenApiSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripSupplychainVehicleSearchRequest) GetRq() *OpenApiSearchRq {
    return r.rq
}

