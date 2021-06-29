package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用车订单查询接口 APIRequest
alitrip.btrip.vehicle.order.search

企业获取商旅用车订单数据
*/
type AlitripBtripVehicleOrderSearchRequest struct {
    model.Params

    // 请求对象
    rq   *OpenSearchRq 

}

func NewAlitripBtripVehicleOrderSearchRequest() *AlitripBtripVehicleOrderSearchRequest{
    return &AlitripBtripVehicleOrderSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripVehicleOrderSearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.vehicle.order.search"
}

func (r AlitripBtripVehicleOrderSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripVehicleOrderSearchRequest) SetRq(rq *OpenSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripVehicleOrderSearchRequest) GetRq() *OpenSearchRq {
    return r.rq
}

