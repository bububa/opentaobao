package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
火车站数据查询 API请求
alitrip.btrip.supplychain.train.city

火车站数据查询
*/
type AlitripBtripSupplychainTrainCityRequest struct {
    model.Params
    // 入参对象
    rq   *OpenSuggestRq
}

// 初始化AlitripBtripSupplychainTrainCityRequest对象
func NewAlitripBtripSupplychainTrainCityRequest() *AlitripBtripSupplychainTrainCityRequest{
    return &AlitripBtripSupplychainTrainCityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainTrainCityRequest) GetApiMethodName() string {
    return "alitrip.btrip.supplychain.train.city"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainTrainCityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripSupplychainTrainCityRequest) SetRq(rq *OpenSuggestRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripSupplychainTrainCityRequest) GetRq() *OpenSuggestRq {
    return r.rq
}
