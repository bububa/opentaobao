package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcitycarapplyaddAPIRequest 三方市内用车申请单同步 API请求
// alitrip.btrip.city.car.apply.add
//
// 三方市内用车申请单同步
type AlitripbtripcitycarapplyaddAPIRequest struct {
	model.Params
	// 入参对象
	_rq *CityCarApplyAddRq
}

// NewAlitripbtripcitycarapplyaddRequest 初始化AlitripbtripcitycarapplyaddAPIRequest对象
func NewAlitripbtripcitycarapplyaddRequest() *AlitripbtripcitycarapplyaddAPIRequest {
	return &AlitripbtripcitycarapplyaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcitycarapplyaddAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.city.car.apply.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcitycarapplyaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcitycarapplyaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripbtripcitycarapplyaddAPIRequest) SetRq(_rq *CityCarApplyAddRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcitycarapplyaddAPIRequest) GetRq() *CityCarApplyAddRq {
	return r._rq
}
