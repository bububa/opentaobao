package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionchangecancelAPIRequest 商旅机票改签取消 API请求
// alitrip.btrip.flight.distribution.change.cancel
//
// 商旅机票改签取消
type AlitripbtripflightdistributionchangecancelAPIRequest struct {
	model.Params
	// 改签取消输入参数
	_paramBtripFlightModifyCancelRq *BtripFlightModifyCancelRq
}

// NewAlitripbtripflightdistributionchangecancelRequest 初始化AlitripbtripflightdistributionchangecancelAPIRequest对象
func NewAlitripbtripflightdistributionchangecancelRequest() *AlitripbtripflightdistributionchangecancelAPIRequest {
	return &AlitripbtripflightdistributionchangecancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionchangecancelAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionchangecancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionchangecancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightModifyCancelRq is ParamBtripFlightModifyCancelRq Setter
// 改签取消输入参数
func (r *AlitripbtripflightdistributionchangecancelAPIRequest) SetParamBtripFlightModifyCancelRq(_paramBtripFlightModifyCancelRq *BtripFlightModifyCancelRq) error {
	r._paramBtripFlightModifyCancelRq = _paramBtripFlightModifyCancelRq
	r.Set("param_btrip_flight_modify_cancel_rq", _paramBtripFlightModifyCancelRq)
	return nil
}

// GetParamBtripFlightModifyCancelRq ParamBtripFlightModifyCancelRq Getter
func (r AlitripbtripflightdistributionchangecancelAPIRequest) GetParamBtripFlightModifyCancelRq() *BtripFlightModifyCancelRq {
	return r._paramBtripFlightModifyCancelRq
}
