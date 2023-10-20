package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionchangenewqueryAPIRequest 改签航班询价V2 API请求
// alitrip.btrip.flight.distribution.change.newquery
//
// 商旅机票改签航班询价
type AlitripbtripflightdistributionchangenewqueryAPIRequest struct {
	model.Params
	// 改签询价入参
	_paramBtripFlightModifyFlightInfoRq *BtripFlightModifyFlightInfoRq
}

// NewAlitripbtripflightdistributionchangenewqueryRequest 初始化AlitripbtripflightdistributionchangenewqueryAPIRequest对象
func NewAlitripbtripflightdistributionchangenewqueryRequest() *AlitripbtripflightdistributionchangenewqueryAPIRequest {
	return &AlitripbtripflightdistributionchangenewqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionchangenewqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.newquery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionchangenewqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionchangenewqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightModifyFlightInfoRq is ParamBtripFlightModifyFlightInfoRq Setter
// 改签询价入参
func (r *AlitripbtripflightdistributionchangenewqueryAPIRequest) SetParamBtripFlightModifyFlightInfoRq(_paramBtripFlightModifyFlightInfoRq *BtripFlightModifyFlightInfoRq) error {
	r._paramBtripFlightModifyFlightInfoRq = _paramBtripFlightModifyFlightInfoRq
	r.Set("param_btrip_flight_modify_flight_info_rq", _paramBtripFlightModifyFlightInfoRq)
	return nil
}

// GetParamBtripFlightModifyFlightInfoRq ParamBtripFlightModifyFlightInfoRq Getter
func (r AlitripbtripflightdistributionchangenewqueryAPIRequest) GetParamBtripFlightModifyFlightInfoRq() *BtripFlightModifyFlightInfoRq {
	return r._paramBtripFlightModifyFlightInfoRq
}
