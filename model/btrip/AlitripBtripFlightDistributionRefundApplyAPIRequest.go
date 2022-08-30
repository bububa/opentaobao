package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionRefundApplyAPIRequest 商旅机票分销-退票申请 API请求
// alitrip.btrip.flight.distribution.refund.apply
//
// 商旅机票分销-退票申请
type AlitripBtripFlightDistributionRefundApplyAPIRequest struct {
	model.Params
	// 退票申请入参
	_paramBtripFlightRefundApplyRq *BtripFlightRefundApplyRq
}

// NewAlitripBtripFlightDistributionRefundApplyRequest 初始化AlitripBtripFlightDistributionRefundApplyAPIRequest对象
func NewAlitripBtripFlightDistributionRefundApplyRequest() *AlitripBtripFlightDistributionRefundApplyAPIRequest {
	return &AlitripBtripFlightDistributionRefundApplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionRefundApplyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.refund.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionRefundApplyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamBtripFlightRefundApplyRq is ParamBtripFlightRefundApplyRq Setter
// 退票申请入参
func (r *AlitripBtripFlightDistributionRefundApplyAPIRequest) SetParamBtripFlightRefundApplyRq(_paramBtripFlightRefundApplyRq *BtripFlightRefundApplyRq) error {
	r._paramBtripFlightRefundApplyRq = _paramBtripFlightRefundApplyRq
	r.Set("param_btrip_flight_refund_apply_rq", _paramBtripFlightRefundApplyRq)
	return nil
}

// GetParamBtripFlightRefundApplyRq ParamBtripFlightRefundApplyRq Getter
func (r AlitripBtripFlightDistributionRefundApplyAPIRequest) GetParamBtripFlightRefundApplyRq() *BtripFlightRefundApplyRq {
	return r._paramBtripFlightRefundApplyRq
}
