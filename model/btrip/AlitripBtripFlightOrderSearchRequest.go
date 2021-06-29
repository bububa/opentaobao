package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取机票订单列表 API请求
alitrip.btrip.flight.order.search

第三方OA厂商获取机票订单列表
*/
type AlitripBtripFlightOrderSearchRequest struct {
    model.Params
    // 请求
    _rq   *OpenSearchRq
}

// 初始化AlitripBtripFlightOrderSearchRequest对象
func NewAlitripBtripFlightOrderSearchRequest() *AlitripBtripFlightOrderSearchRequest{
    return &AlitripBtripFlightOrderSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightOrderSearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.flight.order.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightOrderSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求
func (r *AlitripBtripFlightOrderSearchRequest) SetRq(_rq *OpenSearchRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripFlightOrderSearchRequest) GetRq() *OpenSearchRq {
    return r._rq
}
