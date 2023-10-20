package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripsupplychainflightsearchAPIRequest 【商旅】机票订单查询 API请求
// alitrip.btrip.supplychain.flight.search
//
// 【商旅】机票订单查询
type AlitripbtripsupplychainflightsearchAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenApiSearchRq
}

// NewAlitripbtripsupplychainflightsearchRequest 初始化AlitripbtripsupplychainflightsearchAPIRequest对象
func NewAlitripbtripsupplychainflightsearchRequest() *AlitripbtripsupplychainflightsearchAPIRequest {
	return &AlitripbtripsupplychainflightsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripsupplychainflightsearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.flight.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripsupplychainflightsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripsupplychainflightsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripsupplychainflightsearchAPIRequest) SetRq(_rq *OpenApiSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripsupplychainflightsearchAPIRequest) GetRq() *OpenApiSearchRq {
	return r._rq
}
