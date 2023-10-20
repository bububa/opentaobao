package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenSupplychainFlightTradeAPIRequest 【商旅】机票交易流水查询接口 API请求
// alitrip.btrip.open.supplychain.flight.trade
//
// 【商旅】杭州市政府机票交易流水接口查询
type AlitripBtripOpenSupplychainFlightTradeAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenApiZzdSearchRq
}

// NewAlitripBtripOpenSupplychainFlightTradeRequest 初始化AlitripBtripOpenSupplychainFlightTradeAPIRequest对象
func NewAlitripBtripOpenSupplychainFlightTradeRequest() *AlitripBtripOpenSupplychainFlightTradeAPIRequest {
	return &AlitripBtripOpenSupplychainFlightTradeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripOpenSupplychainFlightTradeAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenSupplychainFlightTradeAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.supplychain.flight.trade"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenSupplychainFlightTradeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripOpenSupplychainFlightTradeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripBtripOpenSupplychainFlightTradeAPIRequest) SetRq(_rq *OpenApiZzdSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripOpenSupplychainFlightTradeAPIRequest) GetRq() *OpenApiZzdSearchRq {
	return r._rq
}

var poolAlitripBtripOpenSupplychainFlightTradeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripOpenSupplychainFlightTradeRequest()
	},
}

// GetAlitripBtripOpenSupplychainFlightTradeRequest 从 sync.Pool 获取 AlitripBtripOpenSupplychainFlightTradeAPIRequest
func GetAlitripBtripOpenSupplychainFlightTradeAPIRequest() *AlitripBtripOpenSupplychainFlightTradeAPIRequest {
	return poolAlitripBtripOpenSupplychainFlightTradeAPIRequest.Get().(*AlitripBtripOpenSupplychainFlightTradeAPIRequest)
}

// ReleaseAlitripBtripOpenSupplychainFlightTradeAPIRequest 将 AlitripBtripOpenSupplychainFlightTradeAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripOpenSupplychainFlightTradeAPIRequest(v *AlitripBtripOpenSupplychainFlightTradeAPIRequest) {
	v.Reset()
	poolAlitripBtripOpenSupplychainFlightTradeAPIRequest.Put(v)
}
