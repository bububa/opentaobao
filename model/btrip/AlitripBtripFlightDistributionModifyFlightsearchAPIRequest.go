package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionmodifyflightsearchAPIRequest 改签航班列表 API请求
// alitrip.btrip.flight.distribution.modify.flightsearch
//
// 商旅分销改签航班列表
type AlitripbtripflightdistributionmodifyflightsearchAPIRequest struct {
	model.Params
	// 改签航班列表入参
	_paramBtripFlightModifySearchPriceRq *BtripFlightModifySearchPriceRq
}

// NewAlitripbtripflightdistributionmodifyflightsearchRequest 初始化AlitripbtripflightdistributionmodifyflightsearchAPIRequest对象
func NewAlitripbtripflightdistributionmodifyflightsearchRequest() *AlitripbtripflightdistributionmodifyflightsearchAPIRequest {
	return &AlitripbtripflightdistributionmodifyflightsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionmodifyflightsearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.modify.flightsearch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionmodifyflightsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionmodifyflightsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightModifySearchPriceRq is ParamBtripFlightModifySearchPriceRq Setter
// 改签航班列表入参
func (r *AlitripbtripflightdistributionmodifyflightsearchAPIRequest) SetParamBtripFlightModifySearchPriceRq(_paramBtripFlightModifySearchPriceRq *BtripFlightModifySearchPriceRq) error {
	r._paramBtripFlightModifySearchPriceRq = _paramBtripFlightModifySearchPriceRq
	r.Set("param_btrip_flight_modify_search_price_rq", _paramBtripFlightModifySearchPriceRq)
	return nil
}

// GetParamBtripFlightModifySearchPriceRq ParamBtripFlightModifySearchPriceRq Getter
func (r AlitripbtripflightdistributionmodifyflightsearchAPIRequest) GetParamBtripFlightModifySearchPriceRq() *BtripFlightModifySearchPriceRq {
	return r._paramBtripFlightModifySearchPriceRq
}
