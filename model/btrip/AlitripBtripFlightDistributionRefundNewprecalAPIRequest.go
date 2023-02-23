package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionRefundNewprecalAPIRequest 商旅机票分销-退票费预计算 API请求
// alitrip.btrip.flight.distribution.refund.newprecal
//
// 商旅机票分销-退票费预计算
type AlitripBtripFlightDistributionRefundNewprecalAPIRequest struct {
	model.Params
	// 退票费预计算请求入参
	_paramBtripFlightRefundPreCalRq *BtripFlightRefundPreCalRq
}

// NewAlitripBtripFlightDistributionRefundNewprecalRequest 初始化AlitripBtripFlightDistributionRefundNewprecalAPIRequest对象
func NewAlitripBtripFlightDistributionRefundNewprecalRequest() *AlitripBtripFlightDistributionRefundNewprecalAPIRequest {
	return &AlitripBtripFlightDistributionRefundNewprecalAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionRefundNewprecalAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.refund.newprecal"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionRefundNewprecalAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionRefundNewprecalAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightRefundPreCalRq is ParamBtripFlightRefundPreCalRq Setter
// 退票费预计算请求入参
func (r *AlitripBtripFlightDistributionRefundNewprecalAPIRequest) SetParamBtripFlightRefundPreCalRq(_paramBtripFlightRefundPreCalRq *BtripFlightRefundPreCalRq) error {
	r._paramBtripFlightRefundPreCalRq = _paramBtripFlightRefundPreCalRq
	r.Set("param_btrip_flight_refund_pre_cal_rq", _paramBtripFlightRefundPreCalRq)
	return nil
}

// GetParamBtripFlightRefundPreCalRq ParamBtripFlightRefundPreCalRq Getter
func (r AlitripBtripFlightDistributionRefundNewprecalAPIRequest) GetParamBtripFlightRefundPreCalRq() *BtripFlightRefundPreCalRq {
	return r._paramBtripFlightRefundPreCalRq
}
