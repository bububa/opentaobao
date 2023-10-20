package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionmodifynewflightsearchAPIRequest 改签航班列表V2 API请求
// alitrip.btrip.flight.distribution.modify.newflightsearch
//
// 改签航班列表V2
type AlitripbtripflightdistributionmodifynewflightsearchAPIRequest struct {
	model.Params
	// 改签航班列表入参
	_paramBtripFlightModifySearchPriceRq *BtripFlightModifySearchPriceRq
}

// NewAlitripbtripflightdistributionmodifynewflightsearchRequest 初始化AlitripbtripflightdistributionmodifynewflightsearchAPIRequest对象
func NewAlitripbtripflightdistributionmodifynewflightsearchRequest() *AlitripbtripflightdistributionmodifynewflightsearchAPIRequest {
	return &AlitripbtripflightdistributionmodifynewflightsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionmodifynewflightsearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.modify.newflightsearch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionmodifynewflightsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionmodifynewflightsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightModifySearchPriceRq is ParamBtripFlightModifySearchPriceRq Setter
// 改签航班列表入参
func (r *AlitripbtripflightdistributionmodifynewflightsearchAPIRequest) SetParamBtripFlightModifySearchPriceRq(_paramBtripFlightModifySearchPriceRq *BtripFlightModifySearchPriceRq) error {
	r._paramBtripFlightModifySearchPriceRq = _paramBtripFlightModifySearchPriceRq
	r.Set("param_btrip_flight_modify_search_price_rq", _paramBtripFlightModifySearchPriceRq)
	return nil
}

// GetParamBtripFlightModifySearchPriceRq ParamBtripFlightModifySearchPriceRq Getter
func (r AlitripbtripflightdistributionmodifynewflightsearchAPIRequest) GetParamBtripFlightModifySearchPriceRq() *BtripFlightModifySearchPriceRq {
	return r._paramBtripFlightModifySearchPriceRq
}
