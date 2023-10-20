package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpopapplyapproveAPIRequest 【商旅】更新审批单状态 API请求
// alitrip.btrip.corpop.apply.approve
//
// 【商旅】更新审批单状态
type AlitripbtripcorpopapplyapproveAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenApiUpdateApplyRq
}

// NewAlitripbtripcorpopapplyapproveRequest 初始化AlitripbtripcorpopapplyapproveAPIRequest对象
func NewAlitripbtripcorpopapplyapproveRequest() *AlitripbtripcorpopapplyapproveAPIRequest {
	return &AlitripbtripcorpopapplyapproveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcorpopapplyapproveAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.apply.approve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcorpopapplyapproveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcorpopapplyapproveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripcorpopapplyapproveAPIRequest) SetRq(_rq *OpenApiUpdateApplyRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcorpopapplyapproveAPIRequest) GetRq() *OpenApiUpdateApplyRq {
	return r._rq
}
