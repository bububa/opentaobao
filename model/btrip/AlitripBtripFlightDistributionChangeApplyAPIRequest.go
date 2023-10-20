package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionchangeapplyAPIRequest 商旅机票改签申请 API请求
// alitrip.btrip.flight.distribution.change.apply
//
// 商旅机票改签申请
type AlitripbtripflightdistributionchangeapplyAPIRequest struct {
	model.Params
	// 改签申请入参
	_paramBtripFlightModifyApplyRq *BtripFlightModifyApplyRq
}

// NewAlitripbtripflightdistributionchangeapplyRequest 初始化AlitripbtripflightdistributionchangeapplyAPIRequest对象
func NewAlitripbtripflightdistributionchangeapplyRequest() *AlitripbtripflightdistributionchangeapplyAPIRequest {
	return &AlitripbtripflightdistributionchangeapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionchangeapplyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionchangeapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionchangeapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightModifyApplyRq is ParamBtripFlightModifyApplyRq Setter
// 改签申请入参
func (r *AlitripbtripflightdistributionchangeapplyAPIRequest) SetParamBtripFlightModifyApplyRq(_paramBtripFlightModifyApplyRq *BtripFlightModifyApplyRq) error {
	r._paramBtripFlightModifyApplyRq = _paramBtripFlightModifyApplyRq
	r.Set("param_btrip_flight_modify_apply_rq", _paramBtripFlightModifyApplyRq)
	return nil
}

// GetParamBtripFlightModifyApplyRq ParamBtripFlightModifyApplyRq Getter
func (r AlitripbtripflightdistributionchangeapplyAPIRequest) GetParamBtripFlightModifyApplyRq() *BtripFlightModifyApplyRq {
	return r._paramBtripFlightModifyApplyRq
}
