package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcitycarapplyqueryAPIRequest 三方市内用车申请单查询 API请求
// alitrip.btrip.city.car.apply.query
//
// 三方市内用车申请单查询
type AlitripbtripcitycarapplyqueryAPIRequest struct {
	model.Params
	// 入参对象
	_rq *CityCarApplyQueryRq
}

// NewAlitripbtripcitycarapplyqueryRequest 初始化AlitripbtripcitycarapplyqueryAPIRequest对象
func NewAlitripbtripcitycarapplyqueryRequest() *AlitripbtripcitycarapplyqueryAPIRequest {
	return &AlitripbtripcitycarapplyqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcitycarapplyqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.city.car.apply.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcitycarapplyqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcitycarapplyqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripbtripcitycarapplyqueryAPIRequest) SetRq(_rq *CityCarApplyQueryRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcitycarapplyqueryAPIRequest) GetRq() *CityCarApplyQueryRq {
	return r._rq
}
