package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】机票交易流水查询接口 APIRequest
alitrip.btrip.open.supplychain.flight.trade

【商旅】杭州市政府机票交易流水接口查询
*/
type AlitripBtripOpenSupplychainFlightTradeRequest struct {
    model.Params

    // 入参对象
    rq   *OpenApiZzdSearchRq 

}

func NewAlitripBtripOpenSupplychainFlightTradeRequest() *AlitripBtripOpenSupplychainFlightTradeRequest{
    return &AlitripBtripOpenSupplychainFlightTradeRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripOpenSupplychainFlightTradeRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.supplychain.flight.trade"
}

func (r AlitripBtripOpenSupplychainFlightTradeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripOpenSupplychainFlightTradeRequest) SetRq(rq *OpenApiZzdSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripOpenSupplychainFlightTradeRequest) GetRq() *OpenApiZzdSearchRq {
    return r.rq
}

