package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainFlightCityAPIRequest 机场数据查询 API请求
// alitrip.btrip.supplychain.flight.city
//
// 机场数据查询
type AlitripBtripSupplychainFlightCityAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenSuggestRq
}

// NewAlitripBtripSupplychainFlightCityRequest 初始化AlitripBtripSupplychainFlightCityAPIRequest对象
func NewAlitripBtripSupplychainFlightCityRequest() *AlitripBtripSupplychainFlightCityAPIRequest {
	return &AlitripBtripSupplychainFlightCityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainFlightCityAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.flight.city"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainFlightCityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripSupplychainFlightCityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripSupplychainFlightCityAPIRequest) SetRq(_rq *OpenSuggestRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripSupplychainFlightCityAPIRequest) GetRq() *OpenSuggestRq {
	return r._rq
}
