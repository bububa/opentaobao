package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripSupplychainFlightSearchAPIRequest
【商旅】机票订单查询 API请求
alitrip.btrip.supplychain.flight.search

【商旅】机票订单查询 */
type AlitripBtripSupplychainFlightSearchAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenApiSearchRq
}

// NewAlitripBtripSupplychainFlightSearchRequest 初始化AlitripBtripSupplychainFlightSearchAPIRequest对象
func NewAlitripBtripSupplychainFlightSearchRequest() *AlitripBtripSupplychainFlightSearchAPIRequest {
	return &AlitripBtripSupplychainFlightSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainFlightSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.flight.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainFlightSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Rq Setter
// 请求对象
func (r *AlitripBtripSupplychainFlightSearchAPIRequest) SetRq(_rq *OpenApiSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// Get Rq Getter
func (r AlitripBtripSupplychainFlightSearchAPIRequest) GetRq() *OpenApiSearchRq {
	return r._rq
}
