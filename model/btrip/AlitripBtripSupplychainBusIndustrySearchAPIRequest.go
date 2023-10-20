package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainBusIndustrySearchAPIRequest 汽车票行业搜索接口 API请求
// alitrip.btrip.supplychain.bus.industry.search
//
// 汽车票行业搜索接口
type AlitripBtripSupplychainBusIndustrySearchAPIRequest struct {
	model.Params
	// 入参
	_rq *BusSearchRq
}

// NewAlitripBtripSupplychainBusIndustrySearchRequest 初始化AlitripBtripSupplychainBusIndustrySearchAPIRequest对象
func NewAlitripBtripSupplychainBusIndustrySearchRequest() *AlitripBtripSupplychainBusIndustrySearchAPIRequest {
	return &AlitripBtripSupplychainBusIndustrySearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripSupplychainBusIndustrySearchAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainBusIndustrySearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.bus.industry.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainBusIndustrySearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripSupplychainBusIndustrySearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripSupplychainBusIndustrySearchAPIRequest) SetRq(_rq *BusSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripSupplychainBusIndustrySearchAPIRequest) GetRq() *BusSearchRq {
	return r._rq
}

var poolAlitripBtripSupplychainBusIndustrySearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripSupplychainBusIndustrySearchRequest()
	},
}

// GetAlitripBtripSupplychainBusIndustrySearchRequest 从 sync.Pool 获取 AlitripBtripSupplychainBusIndustrySearchAPIRequest
func GetAlitripBtripSupplychainBusIndustrySearchAPIRequest() *AlitripBtripSupplychainBusIndustrySearchAPIRequest {
	return poolAlitripBtripSupplychainBusIndustrySearchAPIRequest.Get().(*AlitripBtripSupplychainBusIndustrySearchAPIRequest)
}

// ReleaseAlitripBtripSupplychainBusIndustrySearchAPIRequest 将 AlitripBtripSupplychainBusIndustrySearchAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripSupplychainBusIndustrySearchAPIRequest(v *AlitripBtripSupplychainBusIndustrySearchAPIRequest) {
	v.Reset()
	poolAlitripBtripSupplychainBusIndustrySearchAPIRequest.Put(v)
}
