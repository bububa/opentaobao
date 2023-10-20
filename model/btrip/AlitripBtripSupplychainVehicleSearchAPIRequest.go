package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainVehicleSearchAPIRequest 【商旅】用车订单搜索 API请求
// alitrip.btrip.supplychain.vehicle.search
//
// 【商旅】用车订单搜索
type AlitripBtripSupplychainVehicleSearchAPIRequest struct {
	model.Params
	// 出参
	_rq *OpenApiSearchRq
}

// NewAlitripBtripSupplychainVehicleSearchRequest 初始化AlitripBtripSupplychainVehicleSearchAPIRequest对象
func NewAlitripBtripSupplychainVehicleSearchRequest() *AlitripBtripSupplychainVehicleSearchAPIRequest {
	return &AlitripBtripSupplychainVehicleSearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripSupplychainVehicleSearchAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainVehicleSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.vehicle.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainVehicleSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripSupplychainVehicleSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 出参
func (r *AlitripBtripSupplychainVehicleSearchAPIRequest) SetRq(_rq *OpenApiSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripSupplychainVehicleSearchAPIRequest) GetRq() *OpenApiSearchRq {
	return r._rq
}

var poolAlitripBtripSupplychainVehicleSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripSupplychainVehicleSearchRequest()
	},
}

// GetAlitripBtripSupplychainVehicleSearchRequest 从 sync.Pool 获取 AlitripBtripSupplychainVehicleSearchAPIRequest
func GetAlitripBtripSupplychainVehicleSearchAPIRequest() *AlitripBtripSupplychainVehicleSearchAPIRequest {
	return poolAlitripBtripSupplychainVehicleSearchAPIRequest.Get().(*AlitripBtripSupplychainVehicleSearchAPIRequest)
}

// ReleaseAlitripBtripSupplychainVehicleSearchAPIRequest 将 AlitripBtripSupplychainVehicleSearchAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripSupplychainVehicleSearchAPIRequest(v *AlitripBtripSupplychainVehicleSearchAPIRequest) {
	v.Reset()
	poolAlitripBtripSupplychainVehicleSearchAPIRequest.Put(v)
}
