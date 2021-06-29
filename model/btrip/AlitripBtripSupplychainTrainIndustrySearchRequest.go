package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票行业搜索接口 API请求
alitrip.btrip.supplychain.train.industry.search

【商旅】火车票行业搜索接口
*/
type AlitripBtripSupplychainTrainIndustrySearchRequest struct {
    model.Params
    // 入参
    _rq   *TrainSearchRq
}

// 初始化AlitripBtripSupplychainTrainIndustrySearchRequest对象
func NewAlitripBtripSupplychainTrainIndustrySearchRequest() *AlitripBtripSupplychainTrainIndustrySearchRequest{
    return &AlitripBtripSupplychainTrainIndustrySearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainTrainIndustrySearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.supplychain.train.industry.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainTrainIndustrySearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参
func (r *AlitripBtripSupplychainTrainIndustrySearchRequest) SetRq(_rq *TrainSearchRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripSupplychainTrainIndustrySearchRequest) GetRq() *TrainSearchRq {
    return r._rq
}
