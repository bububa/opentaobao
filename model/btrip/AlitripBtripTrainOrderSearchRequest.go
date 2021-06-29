package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取火车票订单列表 API请求
alitrip.btrip.train.order.search

第三方OA厂商获取自己的火车票数据
*/
type AlitripBtripTrainOrderSearchRequest struct {
    model.Params
    // 请求
    _rq   *OpenSearchRq
}

// 初始化AlitripBtripTrainOrderSearchRequest对象
func NewAlitripBtripTrainOrderSearchRequest() *AlitripBtripTrainOrderSearchRequest{
    return &AlitripBtripTrainOrderSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripTrainOrderSearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.train.order.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripTrainOrderSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求
func (r *AlitripBtripTrainOrderSearchRequest) SetRq(_rq *OpenSearchRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripTrainOrderSearchRequest) GetRq() *OpenSearchRq {
    return r._rq
}
