package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainTrainIndustrySearchAPIRequest 火车票行业搜索接口 API请求
// alitrip.btrip.supplychain.train.industry.search
//
// 【商旅】火车票行业搜索接口
type AlitripBtripSupplychainTrainIndustrySearchAPIRequest struct {
	model.Params
	// 入参
	_rq *TrainSearchRq
}

// NewAlitripBtripSupplychainTrainIndustrySearchRequest 初始化AlitripBtripSupplychainTrainIndustrySearchAPIRequest对象
func NewAlitripBtripSupplychainTrainIndustrySearchRequest() *AlitripBtripSupplychainTrainIndustrySearchAPIRequest {
	return &AlitripBtripSupplychainTrainIndustrySearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainTrainIndustrySearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.train.industry.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainTrainIndustrySearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripSupplychainTrainIndustrySearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripSupplychainTrainIndustrySearchAPIRequest) SetRq(_rq *TrainSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripSupplychainTrainIndustrySearchAPIRequest) GetRq() *TrainSearchRq {
	return r._rq
}
