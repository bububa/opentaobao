package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainHotelSearchAPIRequest 【商旅】酒店订单查询 API请求
// alitrip.btrip.supplychain.hotel.search
//
// 【商旅】酒店订单查询
type AlitripBtripSupplychainHotelSearchAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenApiSearchRq
}

// NewAlitripBtripSupplychainHotelSearchRequest 初始化AlitripBtripSupplychainHotelSearchAPIRequest对象
func NewAlitripBtripSupplychainHotelSearchRequest() *AlitripBtripSupplychainHotelSearchAPIRequest {
	return &AlitripBtripSupplychainHotelSearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripSupplychainHotelSearchAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainHotelSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.hotel.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainHotelSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripSupplychainHotelSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripSupplychainHotelSearchAPIRequest) SetRq(_rq *OpenApiSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripSupplychainHotelSearchAPIRequest) GetRq() *OpenApiSearchRq {
	return r._rq
}

var poolAlitripBtripSupplychainHotelSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripSupplychainHotelSearchRequest()
	},
}

// GetAlitripBtripSupplychainHotelSearchRequest 从 sync.Pool 获取 AlitripBtripSupplychainHotelSearchAPIRequest
func GetAlitripBtripSupplychainHotelSearchAPIRequest() *AlitripBtripSupplychainHotelSearchAPIRequest {
	return poolAlitripBtripSupplychainHotelSearchAPIRequest.Get().(*AlitripBtripSupplychainHotelSearchAPIRequest)
}

// ReleaseAlitripBtripSupplychainHotelSearchAPIRequest 将 AlitripBtripSupplychainHotelSearchAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripSupplychainHotelSearchAPIRequest(v *AlitripBtripSupplychainHotelSearchAPIRequest) {
	v.Reset()
	poolAlitripBtripSupplychainHotelSearchAPIRequest.Put(v)
}
