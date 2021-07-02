package btrip

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenSupplychainHotelTradeAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.supplychain.hotel.trade"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenSupplychainHotelTradeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Rq Setter
// 入参
func (r *AlitripBtripOpenSupplychainHotelTradeAPIRequest) SetRq(_rq *OpenApiZzdSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// Get Rq Getter
func (r AlitripBtripOpenSupplychainHotelTradeAPIRequest) GetRq() *OpenApiZzdSearchRq {
	return r._rq
}
