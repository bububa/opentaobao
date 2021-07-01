package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用车订单查询接口 API请求
alitrip.btrip.vehicle.order.search

企业获取商旅用车订单数据
*/
type AlitripBtripVehicleOrderSearchAPIRequest struct {
    model.Params
    // 请求对象
    _rq   *OpenSearchRq
}

// 初始化AlitripBtripVehicleOrderSearchAPIRequest对象
func NewAlitripBtripVehicleOrderSearchRequest() *AlitripBtripVehicleOrderSearchAPIRequest{
    return &AlitripBtripVehicleOrderSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripVehicleOrderSearchAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.vehicle.order.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripVehicleOrderSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripVehicleOrderSearchAPIRequest) SetRq(_rq *OpenSearchRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripVehicleOrderSearchAPIRequest) GetRq() *OpenSearchRq {
    return r._rq
}
