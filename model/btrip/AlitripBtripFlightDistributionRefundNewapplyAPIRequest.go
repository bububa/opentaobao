package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionrefundnewapplyAPIRequest 商旅机票分销-退票申请 API请求
// alitrip.btrip.flight.distribution.refund.newapply
//
// 商旅机票分销-退票申请
type AlitripbtripflightdistributionrefundnewapplyAPIRequest struct {
	model.Params
	// 退票申请入参
	_paramBtripFlightRefundApplyRq *BtripFlightRefundApplyRq
}

// NewAlitripbtripflightdistributionrefundnewapplyRequest 初始化AlitripbtripflightdistributionrefundnewapplyAPIRequest对象
func NewAlitripbtripflightdistributionrefundnewapplyRequest() *AlitripbtripflightdistributionrefundnewapplyAPIRequest {
	return &AlitripbtripflightdistributionrefundnewapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionrefundnewapplyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.refund.newapply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionrefundnewapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionrefundnewapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightRefundApplyRq is ParamBtripFlightRefundApplyRq Setter
// 退票申请入参
func (r *AlitripbtripflightdistributionrefundnewapplyAPIRequest) SetParamBtripFlightRefundApplyRq(_paramBtripFlightRefundApplyRq *BtripFlightRefundApplyRq) error {
	r._paramBtripFlightRefundApplyRq = _paramBtripFlightRefundApplyRq
	r.Set("param_btrip_flight_refund_apply_rq", _paramBtripFlightRefundApplyRq)
	return nil
}

// GetParamBtripFlightRefundApplyRq ParamBtripFlightRefundApplyRq Getter
func (r AlitripbtripflightdistributionrefundnewapplyAPIRequest) GetParamBtripFlightRefundApplyRq() *BtripFlightRefundApplyRq {
	return r._paramBtripFlightRefundApplyRq
}
