package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionrefundapplyAPIRequest 商旅机票分销-退票申请 API请求
// alitrip.btrip.flight.distribution.refund.apply
//
// 商旅机票分销-退票申请
type AlitripbtripflightdistributionrefundapplyAPIRequest struct {
	model.Params
	// 退票申请入参
	_paramBtripFlightRefundApplyRq *BtripFlightRefundApplyRq
}

// NewAlitripbtripflightdistributionrefundapplyRequest 初始化AlitripbtripflightdistributionrefundapplyAPIRequest对象
func NewAlitripbtripflightdistributionrefundapplyRequest() *AlitripbtripflightdistributionrefundapplyAPIRequest {
	return &AlitripbtripflightdistributionrefundapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionrefundapplyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.refund.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionrefundapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionrefundapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightRefundApplyRq is ParamBtripFlightRefundApplyRq Setter
// 退票申请入参
func (r *AlitripbtripflightdistributionrefundapplyAPIRequest) SetParamBtripFlightRefundApplyRq(_paramBtripFlightRefundApplyRq *BtripFlightRefundApplyRq) error {
	r._paramBtripFlightRefundApplyRq = _paramBtripFlightRefundApplyRq
	r.Set("param_btrip_flight_refund_apply_rq", _paramBtripFlightRefundApplyRq)
	return nil
}

// GetParamBtripFlightRefundApplyRq ParamBtripFlightRefundApplyRq Getter
func (r AlitripbtripflightdistributionrefundapplyAPIRequest) GetParamBtripFlightRefundApplyRq() *BtripFlightRefundApplyRq {
	return r._paramBtripFlightRefundApplyRq
}
