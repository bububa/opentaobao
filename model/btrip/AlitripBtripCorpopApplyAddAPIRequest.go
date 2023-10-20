package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpopapplyaddAPIRequest 【商旅】isv添加审批单 API请求
// alitrip.btrip.corpop.apply.add
//
// 【商旅】isv添加审批单
type AlitripbtripcorpopapplyaddAPIRequest struct {
	model.Params
	// 请求参数
	_rq *OpenApiApplyRq
}

// NewAlitripbtripcorpopapplyaddRequest 初始化AlitripbtripcorpopapplyaddAPIRequest对象
func NewAlitripbtripcorpopapplyaddRequest() *AlitripbtripcorpopapplyaddAPIRequest {
	return &AlitripbtripcorpopapplyaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcorpopapplyaddAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.apply.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcorpopapplyaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcorpopapplyaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求参数
func (r *AlitripbtripcorpopapplyaddAPIRequest) SetRq(_rq *OpenApiApplyRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcorpopapplyaddAPIRequest) GetRq() *OpenApiApplyRq {
	return r._rq
}
