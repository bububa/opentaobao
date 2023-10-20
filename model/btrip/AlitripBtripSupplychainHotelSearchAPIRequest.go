package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripsupplychainhotelsearchAPIRequest 【商旅】酒店订单查询 API请求
// alitrip.btrip.supplychain.hotel.search
//
// 【商旅】酒店订单查询
type AlitripbtripsupplychainhotelsearchAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenApiSearchRq
}

// NewAlitripbtripsupplychainhotelsearchRequest 初始化AlitripbtripsupplychainhotelsearchAPIRequest对象
func NewAlitripbtripsupplychainhotelsearchRequest() *AlitripbtripsupplychainhotelsearchAPIRequest {
	return &AlitripbtripsupplychainhotelsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripsupplychainhotelsearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.hotel.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripsupplychainhotelsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripsupplychainhotelsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripbtripsupplychainhotelsearchAPIRequest) SetRq(_rq *OpenApiSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripsupplychainhotelsearchAPIRequest) GetRq() *OpenApiSearchRq {
	return r._rq
}
