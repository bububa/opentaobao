package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpopcommonapplygetAPIRequest 商旅审批单通用查询接口 API请求
// alitrip.btrip.corpop.commonapply.get
//
// 商旅审批单通用查询接口
type AlitripbtripcorpopcommonapplygetAPIRequest struct {
	model.Params
	// 请求入参
	_rq *OpenIsvSearchRq
}

// NewAlitripbtripcorpopcommonapplygetRequest 初始化AlitripbtripcorpopcommonapplygetAPIRequest对象
func NewAlitripbtripcorpopcommonapplygetRequest() *AlitripbtripcorpopcommonapplygetAPIRequest {
	return &AlitripbtripcorpopcommonapplygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcorpopcommonapplygetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.commonapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcorpopcommonapplygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcorpopcommonapplygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求入参
func (r *AlitripbtripcorpopcommonapplygetAPIRequest) SetRq(_rq *OpenIsvSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcorpopcommonapplygetAPIRequest) GetRq() *OpenIsvSearchRq {
	return r._rq
}
