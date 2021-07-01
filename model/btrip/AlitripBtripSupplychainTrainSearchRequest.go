package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】火车票订单查询 API请求
alitrip.btrip.supplychain.train.search

【商旅】火车票订单查询
*/
type AlitripBtripSupplychainTrainSearchAPIRequest struct {
    model.Params
    // 入参
    _rq   *OpenApiSearchRq
}

// 初始化AlitripBtripSupplychainTrainSearchAPIRequest对象
func NewAlitripBtripSupplychainTrainSearchRequest() *AlitripBtripSupplychainTrainSearchAPIRequest{
    return &AlitripBtripSupplychainTrainSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainTrainSearchAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.supplychain.train.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainTrainSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参
func (r *AlitripBtripSupplychainTrainSearchAPIRequest) SetRq(_rq *OpenApiSearchRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripSupplychainTrainSearchAPIRequest) GetRq() *OpenApiSearchRq {
    return r._rq
}
