package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】酒店交易查询流水接口 API请求
alitrip.btrip.open.supplychain.hotel.trade

【商旅】酒店交易查询流水接口——杭州市政府
*/
type AlitripBtripOpenSupplychainHotelTradeRequest struct {
    model.Params
    // 入参
    _rq   *OpenApiZzdSearchRq
}

// 初始化AlitripBtripOpenSupplychainHotelTradeRequest对象
func NewAlitripBtripOpenSupplychainHotelTradeRequest() *AlitripBtripOpenSupplychainHotelTradeRequest{
    return &AlitripBtripOpenSupplychainHotelTradeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenSupplychainHotelTradeRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.supplychain.hotel.trade"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenSupplychainHotelTradeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参
func (r *AlitripBtripOpenSupplychainHotelTradeRequest) SetRq(_rq *OpenApiZzdSearchRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenSupplychainHotelTradeRequest) GetRq() *OpenApiZzdSearchRq {
    return r._rq
}
