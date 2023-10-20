package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripsupplychaintraindetailsearchvtwoAPIRequest 【商旅】火车票订单详情查询 API请求
// alitrip.btrip.supplychain.train.detail.search.vtwo
//
// 【商旅】火车票订单详情查询
type AlitripbtripsupplychaintraindetailsearchvtwoAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenApiSearchDetailRq
}

// NewAlitripbtripsupplychaintraindetailsearchvtwoRequest 初始化AlitripbtripsupplychaintraindetailsearchvtwoAPIRequest对象
func NewAlitripbtripsupplychaintraindetailsearchvtwoRequest() *AlitripbtripsupplychaintraindetailsearchvtwoAPIRequest {
	return &AlitripbtripsupplychaintraindetailsearchvtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripsupplychaintraindetailsearchvtwoAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.train.detail.search.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripsupplychaintraindetailsearchvtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripsupplychaintraindetailsearchvtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripsupplychaintraindetailsearchvtwoAPIRequest) SetRq(_rq *OpenApiSearchDetailRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripsupplychaintraindetailsearchvtwoAPIRequest) GetRq() *OpenApiSearchDetailRq {
	return r._rq
}
