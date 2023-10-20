package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcitycarapplyapproveAPIRequest 三方市内用车申请单审批 API请求
// alitrip.btrip.city.car.apply.approve
//
// 三方市内用车申请单审批
type AlitripbtripcitycarapplyapproveAPIRequest struct {
	model.Params
	// 入参对象
	_rq *CityCarApplyApproveRq
}

// NewAlitripbtripcitycarapplyapproveRequest 初始化AlitripbtripcitycarapplyapproveAPIRequest对象
func NewAlitripbtripcitycarapplyapproveRequest() *AlitripbtripcitycarapplyapproveAPIRequest {
	return &AlitripbtripcitycarapplyapproveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcitycarapplyapproveAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.city.car.apply.approve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcitycarapplyapproveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcitycarapplyapproveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripbtripcitycarapplyapproveAPIRequest) SetRq(_rq *CityCarApplyApproveRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcitycarapplyapproveAPIRequest) GetRq() *CityCarApplyApproveRq {
	return r._rq
}
