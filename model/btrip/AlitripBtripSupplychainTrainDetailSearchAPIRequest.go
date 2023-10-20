package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripsupplychaintraindetailsearchAPIRequest 【商旅】火车票订单详情查询 API请求
// alitrip.btrip.supplychain.train.detail.search
//
// 【商旅】火车票订单详情查询
type AlitripbtripsupplychaintraindetailsearchAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenApiSearchDetailRq
}

// NewAlitripbtripsupplychaintraindetailsearchRequest 初始化AlitripbtripsupplychaintraindetailsearchAPIRequest对象
func NewAlitripbtripsupplychaintraindetailsearchRequest() *AlitripbtripsupplychaintraindetailsearchAPIRequest {
	return &AlitripbtripsupplychaintraindetailsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripsupplychaintraindetailsearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.train.detail.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripsupplychaintraindetailsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripsupplychaintraindetailsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripsupplychaintraindetailsearchAPIRequest) SetRq(_rq *OpenApiSearchDetailRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripsupplychaintraindetailsearchAPIRequest) GetRq() *OpenApiSearchDetailRq {
	return r._rq
}
