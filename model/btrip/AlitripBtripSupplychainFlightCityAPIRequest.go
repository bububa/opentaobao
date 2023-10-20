package btrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripSupplychainFlightCityAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
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

var poolAlitripBtripSupplychainFlightCityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripSupplychainFlightCityRequest()
	},
}

// GetAlitripBtripSupplychainFlightCityRequest 从 sync.Pool 获取 AlitripBtripSupplychainFlightCityAPIRequest
func GetAlitripBtripSupplychainFlightCityAPIRequest() *AlitripBtripSupplychainFlightCityAPIRequest {
	return poolAlitripBtripSupplychainFlightCityAPIRequest.Get().(*AlitripBtripSupplychainFlightCityAPIRequest)
}

// ReleaseAlitripBtripSupplychainFlightCityAPIRequest 将 AlitripBtripSupplychainFlightCityAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripSupplychainFlightCityAPIRequest(v *AlitripBtripSupplychainFlightCityAPIRequest) {
	v.Reset()
	poolAlitripBtripSupplychainFlightCityAPIRequest.Put(v)
}
