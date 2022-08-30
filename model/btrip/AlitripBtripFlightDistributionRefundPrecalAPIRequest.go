package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionRefundPrecalAPIRequest 商旅机票分销-退票费预计算 API请求
// alitrip.btrip.flight.distribution.refund.precal
//
// 商旅机票分销-退票费预计算
type AlitripBtripFlightDistributionRefundPrecalAPIRequest struct {
	model.Params
	// 退票费预计算请求入参
	_paramBtripFlightRefundPreCalRq *BtripFlightRefundPreCalRq
}

// NewAlitripBtripFlightDistributionRefundPrecalRequest 初始化AlitripBtripFlightDistributionRefundPrecalAPIRequest对象
func NewAlitripBtripFlightDistributionRefundPrecalRequest() *AlitripBtripFlightDistributionRefundPrecalAPIRequest {
	return &AlitripBtripFlightDistributionRefundPrecalAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionRefundPrecalAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.refund.precal"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionRefundPrecalAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamBtripFlightRefundPreCalRq is ParamBtripFlightRefundPreCalRq Setter
// 退票费预计算请求入参
func (r *AlitripBtripFlightDistributionRefundPrecalAPIRequest) SetParamBtripFlightRefundPreCalRq(_paramBtripFlightRefundPreCalRq *BtripFlightRefundPreCalRq) error {
	r._paramBtripFlightRefundPreCalRq = _paramBtripFlightRefundPreCalRq
	r.Set("param_btrip_flight_refund_pre_cal_rq", _paramBtripFlightRefundPreCalRq)
	return nil
}

// GetParamBtripFlightRefundPreCalRq ParamBtripFlightRefundPreCalRq Getter
func (r AlitripBtripFlightDistributionRefundPrecalAPIRequest) GetParamBtripFlightRefundPreCalRq() *BtripFlightRefundPreCalRq {
	return r._paramBtripFlightRefundPreCalRq
}
