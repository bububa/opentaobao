package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpopapplymodifyAPIRequest 【商旅】修改出差审批单（行程） API请求
// alitrip.btrip.corpop.apply.modify
//
// 【商旅】修改出差审批单（行程）
type AlitripbtripcorpopapplymodifyAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenApiApplyRq
}

// NewAlitripbtripcorpopapplymodifyRequest 初始化AlitripbtripcorpopapplymodifyAPIRequest对象
func NewAlitripbtripcorpopapplymodifyRequest() *AlitripbtripcorpopapplymodifyAPIRequest {
	return &AlitripbtripcorpopapplymodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcorpopapplymodifyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.apply.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcorpopapplymodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcorpopapplymodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripcorpopapplymodifyAPIRequest) SetRq(_rq *OpenApiApplyRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcorpopapplymodifyAPIRequest) GetRq() *OpenApiApplyRq {
	return r._rq
}
