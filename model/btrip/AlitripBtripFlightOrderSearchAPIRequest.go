package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightOrderSearchAPIRequest 获取机票订单列表 API请求
// alitrip.btrip.flight.order.search
//
// 第三方OA厂商获取机票订单列表
type AlitripBtripFlightOrderSearchAPIRequest struct {
	model.Params
	// 请求
	_rq *OpenSearchRq
}

// NewAlitripBtripFlightOrderSearchRequest 初始化AlitripBtripFlightOrderSearchAPIRequest对象
func NewAlitripBtripFlightOrderSearchRequest() *AlitripBtripFlightOrderSearchAPIRequest {
	return &AlitripBtripFlightOrderSearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripFlightOrderSearchAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightOrderSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightOrderSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightOrderSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求
func (r *AlitripBtripFlightOrderSearchAPIRequest) SetRq(_rq *OpenSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripFlightOrderSearchAPIRequest) GetRq() *OpenSearchRq {
	return r._rq
}

var poolAlitripBtripFlightOrderSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripFlightOrderSearchRequest()
	},
}

// GetAlitripBtripFlightOrderSearchRequest 从 sync.Pool 获取 AlitripBtripFlightOrderSearchAPIRequest
func GetAlitripBtripFlightOrderSearchAPIRequest() *AlitripBtripFlightOrderSearchAPIRequest {
	return poolAlitripBtripFlightOrderSearchAPIRequest.Get().(*AlitripBtripFlightOrderSearchAPIRequest)
}

// ReleaseAlitripBtripFlightOrderSearchAPIRequest 将 AlitripBtripFlightOrderSearchAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripFlightOrderSearchAPIRequest(v *AlitripBtripFlightOrderSearchAPIRequest) {
	v.Reset()
	poolAlitripBtripFlightOrderSearchAPIRequest.Put(v)
}
