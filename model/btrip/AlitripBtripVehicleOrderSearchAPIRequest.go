package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripVehicleOrderSearchAPIRequest 用车订单查询接口 API请求
// alitrip.btrip.vehicle.order.search
//
// 企业获取商旅用车订单数据
type AlitripBtripVehicleOrderSearchAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenSearchRq
}

// NewAlitripBtripVehicleOrderSearchRequest 初始化AlitripBtripVehicleOrderSearchAPIRequest对象
func NewAlitripBtripVehicleOrderSearchRequest() *AlitripBtripVehicleOrderSearchAPIRequest {
	return &AlitripBtripVehicleOrderSearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripVehicleOrderSearchAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripVehicleOrderSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.vehicle.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripVehicleOrderSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripVehicleOrderSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripVehicleOrderSearchAPIRequest) SetRq(_rq *OpenSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripVehicleOrderSearchAPIRequest) GetRq() *OpenSearchRq {
	return r._rq
}

var poolAlitripBtripVehicleOrderSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripVehicleOrderSearchRequest()
	},
}

// GetAlitripBtripVehicleOrderSearchRequest 从 sync.Pool 获取 AlitripBtripVehicleOrderSearchAPIRequest
func GetAlitripBtripVehicleOrderSearchAPIRequest() *AlitripBtripVehicleOrderSearchAPIRequest {
	return poolAlitripBtripVehicleOrderSearchAPIRequest.Get().(*AlitripBtripVehicleOrderSearchAPIRequest)
}

// ReleaseAlitripBtripVehicleOrderSearchAPIRequest 将 AlitripBtripVehicleOrderSearchAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripVehicleOrderSearchAPIRequest(v *AlitripBtripVehicleOrderSearchAPIRequest) {
	v.Reset()
	poolAlitripBtripVehicleOrderSearchAPIRequest.Put(v)
}
