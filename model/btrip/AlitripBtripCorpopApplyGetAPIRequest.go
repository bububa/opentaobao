package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpopapplygetAPIRequest 【商旅】查询审批单 API请求
// alitrip.btrip.corpop.apply.get
//
// 【商旅】查询审批单
type AlitripbtripcorpopapplygetAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenIsvSearchRq
}

// NewAlitripbtripcorpopapplygetRequest 初始化AlitripbtripcorpopapplygetAPIRequest对象
func NewAlitripbtripcorpopapplygetRequest() *AlitripbtripcorpopapplygetAPIRequest {
	return &AlitripbtripcorpopapplygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcorpopapplygetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.apply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcorpopapplygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcorpopapplygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripcorpopapplygetAPIRequest) SetRq(_rq *OpenIsvSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcorpopapplygetAPIRequest) GetRq() *OpenIsvSearchRq {
	return r._rq
}
