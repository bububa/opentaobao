package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCityCarApplyQueryAPIRequest 三方市内用车申请单查询 API请求
// alitrip.btrip.city.car.apply.query
//
// 三方市内用车申请单查询
type AlitripBtripCityCarApplyQueryAPIRequest struct {
	model.Params
	// 入参对象
	_rq *CityCarApplyQueryRq
}

// NewAlitripBtripCityCarApplyQueryRequest 初始化AlitripBtripCityCarApplyQueryAPIRequest对象
func NewAlitripBtripCityCarApplyQueryRequest() *AlitripBtripCityCarApplyQueryAPIRequest {
	return &AlitripBtripCityCarApplyQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCityCarApplyQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.city.car.apply.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCityCarApplyQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCityCarApplyQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripBtripCityCarApplyQueryAPIRequest) SetRq(_rq *CityCarApplyQueryRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCityCarApplyQueryAPIRequest) GetRq() *CityCarApplyQueryRq {
	return r._rq
}
