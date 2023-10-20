package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainFlightDetailSearchAPIRequest 【商旅】机票订单详情查询 API请求
// alitrip.btrip.supplychain.flight.detail.search
//
// 【商旅】机票订单详情查询
type AlitripBtripSupplychainFlightDetailSearchAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenApiSearchDetailRq
}

// NewAlitripBtripSupplychainFlightDetailSearchRequest 初始化AlitripBtripSupplychainFlightDetailSearchAPIRequest对象
func NewAlitripBtripSupplychainFlightDetailSearchRequest() *AlitripBtripSupplychainFlightDetailSearchAPIRequest {
	return &AlitripBtripSupplychainFlightDetailSearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripSupplychainFlightDetailSearchAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainFlightDetailSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.flight.detail.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainFlightDetailSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripSupplychainFlightDetailSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripSupplychainFlightDetailSearchAPIRequest) SetRq(_rq *OpenApiSearchDetailRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripSupplychainFlightDetailSearchAPIRequest) GetRq() *OpenApiSearchDetailRq {
	return r._rq
}

var poolAlitripBtripSupplychainFlightDetailSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripSupplychainFlightDetailSearchRequest()
	},
}

// GetAlitripBtripSupplychainFlightDetailSearchRequest 从 sync.Pool 获取 AlitripBtripSupplychainFlightDetailSearchAPIRequest
func GetAlitripBtripSupplychainFlightDetailSearchAPIRequest() *AlitripBtripSupplychainFlightDetailSearchAPIRequest {
	return poolAlitripBtripSupplychainFlightDetailSearchAPIRequest.Get().(*AlitripBtripSupplychainFlightDetailSearchAPIRequest)
}

// ReleaseAlitripBtripSupplychainFlightDetailSearchAPIRequest 将 AlitripBtripSupplychainFlightDetailSearchAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripSupplychainFlightDetailSearchAPIRequest(v *AlitripBtripSupplychainFlightDetailSearchAPIRequest) {
	v.Reset()
	poolAlitripBtripSupplychainFlightDetailSearchAPIRequest.Put(v)
}
