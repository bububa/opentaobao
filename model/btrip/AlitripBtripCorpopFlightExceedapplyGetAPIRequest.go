package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpopflightexceedapplygetAPIRequest 商旅机票第三方超标审批单搜索接口 API请求
// alitrip.btrip.corpop.flight.exceedapply.get
//
// 商旅机票第三方超标审批单搜索接口
type AlitripbtripcorpopflightexceedapplygetAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenIsvSearchRq
}

// NewAlitripbtripcorpopflightexceedapplygetRequest 初始化AlitripbtripcorpopflightexceedapplygetAPIRequest对象
func NewAlitripbtripcorpopflightexceedapplygetRequest() *AlitripbtripcorpopflightexceedapplygetAPIRequest {
	return &AlitripbtripcorpopflightexceedapplygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcorpopflightexceedapplygetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.flight.exceedapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcorpopflightexceedapplygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcorpopflightexceedapplygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripbtripcorpopflightexceedapplygetAPIRequest) SetRq(_rq *OpenIsvSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcorpopflightexceedapplygetAPIRequest) GetRq() *OpenIsvSearchRq {
	return r._rq
}
