package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionRefundNewapplyAPIRequest 商旅机票分销-退票申请 API请求
// alitrip.btrip.flight.distribution.refund.newapply
//
// 商旅机票分销-退票申请
type AlitripBtripFlightDistributionRefundNewapplyAPIRequest struct {
	model.Params
	// 退票申请入参
	_paramBtripFlightRefundApplyRq *BtripFlightRefundApplyRq
}

// NewAlitripBtripFlightDistributionRefundNewapplyRequest 初始化AlitripBtripFlightDistributionRefundNewapplyAPIRequest对象
func NewAlitripBtripFlightDistributionRefundNewapplyRequest() *AlitripBtripFlightDistributionRefundNewapplyAPIRequest {
	return &AlitripBtripFlightDistributionRefundNewapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionRefundNewapplyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.refund.newapply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionRefundNewapplyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamBtripFlightRefundApplyRq is ParamBtripFlightRefundApplyRq Setter
// 退票申请入参
func (r *AlitripBtripFlightDistributionRefundNewapplyAPIRequest) SetParamBtripFlightRefundApplyRq(_paramBtripFlightRefundApplyRq *BtripFlightRefundApplyRq) error {
	r._paramBtripFlightRefundApplyRq = _paramBtripFlightRefundApplyRq
	r.Set("param_btrip_flight_refund_apply_rq", _paramBtripFlightRefundApplyRq)
	return nil
}

// GetParamBtripFlightRefundApplyRq ParamBtripFlightRefundApplyRq Getter
func (r AlitripBtripFlightDistributionRefundNewapplyAPIRequest) GetParamBtripFlightRefundApplyRq() *BtripFlightRefundApplyRq {
	return r._paramBtripFlightRefundApplyRq
}
