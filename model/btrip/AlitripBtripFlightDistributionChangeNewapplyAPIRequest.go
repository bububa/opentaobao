package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionchangenewapplyAPIRequest 商旅机票改签申请V2 API请求
// alitrip.btrip.flight.distribution.change.newapply
//
// 商旅机票改签申请
type AlitripbtripflightdistributionchangenewapplyAPIRequest struct {
	model.Params
	// 改签申请入参
	_paramBtripFlightModifyApplyRq *BtripFlightModifyApplyRq
}

// NewAlitripbtripflightdistributionchangenewapplyRequest 初始化AlitripbtripflightdistributionchangenewapplyAPIRequest对象
func NewAlitripbtripflightdistributionchangenewapplyRequest() *AlitripbtripflightdistributionchangenewapplyAPIRequest {
	return &AlitripbtripflightdistributionchangenewapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionchangenewapplyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.newapply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionchangenewapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionchangenewapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightModifyApplyRq is ParamBtripFlightModifyApplyRq Setter
// 改签申请入参
func (r *AlitripbtripflightdistributionchangenewapplyAPIRequest) SetParamBtripFlightModifyApplyRq(_paramBtripFlightModifyApplyRq *BtripFlightModifyApplyRq) error {
	r._paramBtripFlightModifyApplyRq = _paramBtripFlightModifyApplyRq
	r.Set("param_btrip_flight_modify_apply_rq", _paramBtripFlightModifyApplyRq)
	return nil
}

// GetParamBtripFlightModifyApplyRq ParamBtripFlightModifyApplyRq Getter
func (r AlitripbtripflightdistributionchangenewapplyAPIRequest) GetParamBtripFlightModifyApplyRq() *BtripFlightModifyApplyRq {
	return r._paramBtripFlightModifyApplyRq
}
