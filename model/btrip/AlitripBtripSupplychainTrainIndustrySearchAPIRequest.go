package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripsupplychaintrainindustrysearchAPIRequest 火车票行业搜索接口 API请求
// alitrip.btrip.supplychain.train.industry.search
//
// 【商旅】火车票行业搜索接口
type AlitripbtripsupplychaintrainindustrysearchAPIRequest struct {
	model.Params
	// 入参
	_rq *TrainSearchRq
}

// NewAlitripbtripsupplychaintrainindustrysearchRequest 初始化AlitripbtripsupplychaintrainindustrysearchAPIRequest对象
func NewAlitripbtripsupplychaintrainindustrysearchRequest() *AlitripbtripsupplychaintrainindustrysearchAPIRequest {
	return &AlitripbtripsupplychaintrainindustrysearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripsupplychaintrainindustrysearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.train.industry.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripsupplychaintrainindustrysearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripsupplychaintrainindustrysearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripbtripsupplychaintrainindustrysearchAPIRequest) SetRq(_rq *TrainSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripsupplychaintrainindustrysearchAPIRequest) GetRq() *TrainSearchRq {
	return r._rq
}
