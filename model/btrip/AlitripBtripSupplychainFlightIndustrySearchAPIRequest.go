package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainFlightIndustrySearchAPIRequest 机票行业搜索接口 API请求
// alitrip.btrip.supplychain.flight.industry.search
//
// 【商旅】机票行业搜索
type AlitripBtripSupplychainFlightIndustrySearchAPIRequest struct {
	model.Params
	// 入参
	_rq *FlightSearchRq
}

// NewAlitripBtripSupplychainFlightIndustrySearchRequest 初始化AlitripBtripSupplychainFlightIndustrySearchAPIRequest对象
func NewAlitripBtripSupplychainFlightIndustrySearchRequest() *AlitripBtripSupplychainFlightIndustrySearchAPIRequest {
	return &AlitripBtripSupplychainFlightIndustrySearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripSupplychainFlightIndustrySearchAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainFlightIndustrySearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.flight.industry.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainFlightIndustrySearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripSupplychainFlightIndustrySearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripSupplychainFlightIndustrySearchAPIRequest) SetRq(_rq *FlightSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripSupplychainFlightIndustrySearchAPIRequest) GetRq() *FlightSearchRq {
	return r._rq
}

var poolAlitripBtripSupplychainFlightIndustrySearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripSupplychainFlightIndustrySearchRequest()
	},
}

// GetAlitripBtripSupplychainFlightIndustrySearchRequest 从 sync.Pool 获取 AlitripBtripSupplychainFlightIndustrySearchAPIRequest
func GetAlitripBtripSupplychainFlightIndustrySearchAPIRequest() *AlitripBtripSupplychainFlightIndustrySearchAPIRequest {
	return poolAlitripBtripSupplychainFlightIndustrySearchAPIRequest.Get().(*AlitripBtripSupplychainFlightIndustrySearchAPIRequest)
}

// ReleaseAlitripBtripSupplychainFlightIndustrySearchAPIRequest 将 AlitripBtripSupplychainFlightIndustrySearchAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripSupplychainFlightIndustrySearchAPIRequest(v *AlitripBtripSupplychainFlightIndustrySearchAPIRequest) {
	v.Reset()
	poolAlitripBtripSupplychainFlightIndustrySearchAPIRequest.Put(v)
}
