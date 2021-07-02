package btrip

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainVehicleSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.vehicle.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainVehicleSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Rq Setter
// 出参
func (r *AlitripBtripSupplychainVehicleSearchAPIRequest) SetRq(_rq *OpenApiSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// Get Rq Getter
func (r AlitripBtripSupplychainVehicleSearchAPIRequest) GetRq() *OpenApiSearchRq {
	return r._rq
}
