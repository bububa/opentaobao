package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpophotelexceedapplygetAPIRequest 商旅酒店第三方超标审批单搜索接口 API请求
// alitrip.btrip.corpop.hotel.exceedapply.get
//
// 商旅酒店第三方超标审批单搜索接口
type AlitripbtripcorpophotelexceedapplygetAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenIsvSearchRq
}

// NewAlitripbtripcorpophotelexceedapplygetRequest 初始化AlitripbtripcorpophotelexceedapplygetAPIRequest对象
func NewAlitripbtripcorpophotelexceedapplygetRequest() *AlitripbtripcorpophotelexceedapplygetAPIRequest {
	return &AlitripbtripcorpophotelexceedapplygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcorpophotelexceedapplygetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.hotel.exceedapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcorpophotelexceedapplygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcorpophotelexceedapplygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripbtripcorpophotelexceedapplygetAPIRequest) SetRq(_rq *OpenIsvSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcorpophotelexceedapplygetAPIRequest) GetRq() *OpenIsvSearchRq {
	return r._rq
}
