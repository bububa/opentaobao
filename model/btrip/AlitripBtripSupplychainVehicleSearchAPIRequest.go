package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripsupplychainvehiclesearchAPIRequest 【商旅】用车订单搜索 API请求
// alitrip.btrip.supplychain.vehicle.search
//
// 【商旅】用车订单搜索
type AlitripbtripsupplychainvehiclesearchAPIRequest struct {
	model.Params
	// 出参
	_rq *OpenApiSearchRq
}

// NewAlitripbtripsupplychainvehiclesearchRequest 初始化AlitripbtripsupplychainvehiclesearchAPIRequest对象
func NewAlitripbtripsupplychainvehiclesearchRequest() *AlitripbtripsupplychainvehiclesearchAPIRequest {
	return &AlitripbtripsupplychainvehiclesearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripsupplychainvehiclesearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.vehicle.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripsupplychainvehiclesearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripsupplychainvehiclesearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 出参
func (r *AlitripbtripsupplychainvehiclesearchAPIRequest) SetRq(_rq *OpenApiSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripsupplychainvehiclesearchAPIRequest) GetRq() *OpenApiSearchRq {
	return r._rq
}
