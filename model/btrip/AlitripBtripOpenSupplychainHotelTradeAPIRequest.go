package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenSupplychainHotelTradeAPIRequest 【商旅】酒店交易查询流水接口 API请求
// alitrip.btrip.open.supplychain.hotel.trade
//
// 【商旅】酒店交易查询流水接口——杭州市政府
type AlitripBtripOpenSupplychainHotelTradeAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenApiZzdSearchRq
}

// NewAlitripBtripOpenSupplychainHotelTradeRequest 初始化AlitripBtripOpenSupplychainHotelTradeAPIRequest对象
func NewAlitripBtripOpenSupplychainHotelTradeRequest() *AlitripBtripOpenSupplychainHotelTradeAPIRequest {
	return &AlitripBtripOpenSupplychainHotelTradeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripOpenSupplychainHotelTradeAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenSupplychainHotelTradeAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.supplychain.hotel.trade"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenSupplychainHotelTradeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripOpenSupplychainHotelTradeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripOpenSupplychainHotelTradeAPIRequest) SetRq(_rq *OpenApiZzdSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripOpenSupplychainHotelTradeAPIRequest) GetRq() *OpenApiZzdSearchRq {
	return r._rq
}

var poolAlitripBtripOpenSupplychainHotelTradeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripOpenSupplychainHotelTradeRequest()
	},
}

// GetAlitripBtripOpenSupplychainHotelTradeRequest 从 sync.Pool 获取 AlitripBtripOpenSupplychainHotelTradeAPIRequest
func GetAlitripBtripOpenSupplychainHotelTradeAPIRequest() *AlitripBtripOpenSupplychainHotelTradeAPIRequest {
	return poolAlitripBtripOpenSupplychainHotelTradeAPIRequest.Get().(*AlitripBtripOpenSupplychainHotelTradeAPIRequest)
}

// ReleaseAlitripBtripOpenSupplychainHotelTradeAPIRequest 将 AlitripBtripOpenSupplychainHotelTradeAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripOpenSupplychainHotelTradeAPIRequest(v *AlitripBtripOpenSupplychainHotelTradeAPIRequest) {
	v.Reset()
	poolAlitripBtripOpenSupplychainHotelTradeAPIRequest.Put(v)
}
