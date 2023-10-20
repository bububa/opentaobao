package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionchangenewcancelAPIRequest 商旅机票改签取消 API请求
// alitrip.btrip.flight.distribution.change.newcancel
//
// 商旅机票改签取消
type AlitripbtripflightdistributionchangenewcancelAPIRequest struct {
	model.Params
	// 改签取消输入参数
	_paramBtripFlightModifyCancelRq *BtripFlightModifyCancelRq
}

// NewAlitripbtripflightdistributionchangenewcancelRequest 初始化AlitripbtripflightdistributionchangenewcancelAPIRequest对象
func NewAlitripbtripflightdistributionchangenewcancelRequest() *AlitripbtripflightdistributionchangenewcancelAPIRequest {
	return &AlitripbtripflightdistributionchangenewcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionchangenewcancelAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.newcancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionchangenewcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionchangenewcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightModifyCancelRq is ParamBtripFlightModifyCancelRq Setter
// 改签取消输入参数
func (r *AlitripbtripflightdistributionchangenewcancelAPIRequest) SetParamBtripFlightModifyCancelRq(_paramBtripFlightModifyCancelRq *BtripFlightModifyCancelRq) error {
	r._paramBtripFlightModifyCancelRq = _paramBtripFlightModifyCancelRq
	r.Set("param_btrip_flight_modify_cancel_rq", _paramBtripFlightModifyCancelRq)
	return nil
}

// GetParamBtripFlightModifyCancelRq ParamBtripFlightModifyCancelRq Getter
func (r AlitripbtripflightdistributionchangenewcancelAPIRequest) GetParamBtripFlightModifyCancelRq() *BtripFlightModifyCancelRq {
	return r._paramBtripFlightModifyCancelRq
}
