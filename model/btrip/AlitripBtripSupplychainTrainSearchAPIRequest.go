package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripsupplychaintrainsearchAPIRequest 【商旅】火车票订单查询 API请求
// alitrip.btrip.supplychain.train.search
//
// 【商旅】火车票订单查询
type AlitripbtripsupplychaintrainsearchAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenApiSearchRq
}

// NewAlitripbtripsupplychaintrainsearchRequest 初始化AlitripbtripsupplychaintrainsearchAPIRequest对象
func NewAlitripbtripsupplychaintrainsearchRequest() *AlitripbtripsupplychaintrainsearchAPIRequest {
	return &AlitripbtripsupplychaintrainsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripsupplychaintrainsearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.train.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripsupplychaintrainsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripsupplychaintrainsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripbtripsupplychaintrainsearchAPIRequest) SetRq(_rq *OpenApiSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripsupplychaintrainsearchAPIRequest) GetRq() *OpenApiSearchRq {
	return r._rq
}
