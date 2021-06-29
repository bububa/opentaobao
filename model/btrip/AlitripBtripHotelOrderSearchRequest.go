package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索酒店订单列表 API请求
alitrip.btrip.hotel.order.search

企业获取商旅酒店订单数据
*/
type AlitripBtripHotelOrderSearchRequest struct {
    model.Params
    // 请求
    _rq   *OpenSearchRq
}

// 初始化AlitripBtripHotelOrderSearchRequest对象
func NewAlitripBtripHotelOrderSearchRequest() *AlitripBtripHotelOrderSearchRequest{
    return &AlitripBtripHotelOrderSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelOrderSearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.hotel.order.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelOrderSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求
func (r *AlitripBtripHotelOrderSearchRequest) SetRq(_rq *OpenSearchRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripHotelOrderSearchRequest) GetRq() *OpenSearchRq {
    return r._rq
}
