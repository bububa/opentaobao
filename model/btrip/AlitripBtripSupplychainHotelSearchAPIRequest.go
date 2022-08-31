package btrip

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainHotelSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.hotel.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainHotelSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
