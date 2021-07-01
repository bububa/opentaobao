package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】机票交易流水查询接口 API请求
alitrip.btrip.open.supplychain.flight.trade

【商旅】杭州市政府机票交易流水接口查询
*/
type AlitripBtripOpenSupplychainFlightTradeAPIRequest struct {
    model.Params
    // 入参对象
    _rq   *OpenApiZzdSearchRq
}

// 初始化AlitripBtripOpenSupplychainFlightTradeAPIRequest对象
func NewAlitripBtripOpenSupplychainFlightTradeRequest() *AlitripBtripOpenSupplychainFlightTradeAPIRequest{
    return &AlitripBtripOpenSupplychainFlightTradeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenSupplychainFlightTradeAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.open.supplychain.flight.trade"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenSupplychainFlightTradeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripOpenSupplychainFlightTradeAPIRequest) SetRq(_rq *OpenApiZzdSearchRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenSupplychainFlightTradeAPIRequest) GetRq() *OpenApiZzdSearchRq {
    return r._rq
}
