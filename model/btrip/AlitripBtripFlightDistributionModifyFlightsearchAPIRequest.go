package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionModifyFlightsearchAPIRequest 改签航班列表 API请求
// alitrip.btrip.flight.distribution.modify.flightsearch
//
// 商旅分销改签航班列表
type AlitripBtripFlightDistributionModifyFlightsearchAPIRequest struct {
	model.Params
	// 改签航班列表入参
	_paramBtripFlightModifySearchPriceRq *BtripFlightModifySearchPriceRq
}

// NewAlitripBtripFlightDistributionModifyFlightsearchRequest 初始化AlitripBtripFlightDistributionModifyFlightsearchAPIRequest对象
func NewAlitripBtripFlightDistributionModifyFlightsearchRequest() *AlitripBtripFlightDistributionModifyFlightsearchAPIRequest {
	return &AlitripBtripFlightDistributionModifyFlightsearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripFlightDistributionModifyFlightsearchAPIRequest) Reset() {
	r._paramBtripFlightModifySearchPriceRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionModifyFlightsearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.modify.flightsearch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionModifyFlightsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionModifyFlightsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightModifySearchPriceRq is ParamBtripFlightModifySearchPriceRq Setter
// 改签航班列表入参
func (r *AlitripBtripFlightDistributionModifyFlightsearchAPIRequest) SetParamBtripFlightModifySearchPriceRq(_paramBtripFlightModifySearchPriceRq *BtripFlightModifySearchPriceRq) error {
	r._paramBtripFlightModifySearchPriceRq = _paramBtripFlightModifySearchPriceRq
	r.Set("param_btrip_flight_modify_search_price_rq", _paramBtripFlightModifySearchPriceRq)
	return nil
}

// GetParamBtripFlightModifySearchPriceRq ParamBtripFlightModifySearchPriceRq Getter
func (r AlitripBtripFlightDistributionModifyFlightsearchAPIRequest) GetParamBtripFlightModifySearchPriceRq() *BtripFlightModifySearchPriceRq {
	return r._paramBtripFlightModifySearchPriceRq
}

var poolAlitripBtripFlightDistributionModifyFlightsearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripFlightDistributionModifyFlightsearchRequest()
	},
}

// GetAlitripBtripFlightDistributionModifyFlightsearchRequest 从 sync.Pool 获取 AlitripBtripFlightDistributionModifyFlightsearchAPIRequest
func GetAlitripBtripFlightDistributionModifyFlightsearchAPIRequest() *AlitripBtripFlightDistributionModifyFlightsearchAPIRequest {
	return poolAlitripBtripFlightDistributionModifyFlightsearchAPIRequest.Get().(*AlitripBtripFlightDistributionModifyFlightsearchAPIRequest)
}

// ReleaseAlitripBtripFlightDistributionModifyFlightsearchAPIRequest 将 AlitripBtripFlightDistributionModifyFlightsearchAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripFlightDistributionModifyFlightsearchAPIRequest(v *AlitripBtripFlightDistributionModifyFlightsearchAPIRequest) {
	v.Reset()
	poolAlitripBtripFlightDistributionModifyFlightsearchAPIRequest.Put(v)
}
