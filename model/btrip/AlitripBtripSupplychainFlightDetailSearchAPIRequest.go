package btrip

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
