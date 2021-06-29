package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】酒店交易查询流水接口 APIRequest
alitrip.btrip.open.supplychain.hotel.trade

【商旅】酒店交易查询流水接口——杭州市政府
*/
type AlitripBtripOpenSupplychainHotelTradeRequest struct {
    model.Params

    // 入参
    rq   *OpenApiZzdSearchRq 

}

func NewAlitripBtripOpenSupplychainHotelTradeRequest() *AlitripBtripOpenSupplychainHotelTradeRequest{
    return &AlitripBtripOpenSupplychainHotelTradeRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripOpenSupplychainHotelTradeRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.supplychain.hotel.trade"
}

func (r AlitripBtripOpenSupplychainHotelTradeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripOpenSupplychainHotelTradeRequest) SetRq(rq *OpenApiZzdSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripOpenSupplychainHotelTradeRequest) GetRq() *OpenApiZzdSearchRq {
    return r.rq
}

