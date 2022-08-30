package btrip

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionModifyFlightsearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.modify.flightsearch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionModifyFlightsearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
