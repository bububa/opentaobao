package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripsupplychainflightdetailsearchAPIRequest 【商旅】机票订单详情查询 API请求
// alitrip.btrip.supplychain.flight.detail.search
//
// 【商旅】机票订单详情查询
type AlitripbtripsupplychainflightdetailsearchAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenApiSearchDetailRq
}

// NewAlitripbtripsupplychainflightdetailsearchRequest 初始化AlitripbtripsupplychainflightdetailsearchAPIRequest对象
func NewAlitripbtripsupplychainflightdetailsearchRequest() *AlitripbtripsupplychainflightdetailsearchAPIRequest {
	return &AlitripbtripsupplychainflightdetailsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripsupplychainflightdetailsearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.flight.detail.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripsupplychainflightdetailsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripsupplychainflightdetailsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripsupplychainflightdetailsearchAPIRequest) SetRq(_rq *OpenApiSearchDetailRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripsupplychainflightdetailsearchAPIRequest) GetRq() *OpenApiSearchDetailRq {
	return r._rq
}
