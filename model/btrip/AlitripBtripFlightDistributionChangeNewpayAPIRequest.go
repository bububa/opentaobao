package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionchangenewpayAPIRequest 商旅机票改签支付V2 API请求
// alitrip.btrip.flight.distribution.change.newpay
//
// 商旅机票改签支付V2
type AlitripbtripflightdistributionchangenewpayAPIRequest struct {
	model.Params
	// 改签支付入参
	_paramBtripFlightModifyPayRq *BtripFlightModifyPayRq
}

// NewAlitripbtripflightdistributionchangenewpayRequest 初始化AlitripbtripflightdistributionchangenewpayAPIRequest对象
func NewAlitripbtripflightdistributionchangenewpayRequest() *AlitripbtripflightdistributionchangenewpayAPIRequest {
	return &AlitripbtripflightdistributionchangenewpayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionchangenewpayAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.newpay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionchangenewpayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionchangenewpayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightModifyPayRq is ParamBtripFlightModifyPayRq Setter
// 改签支付入参
func (r *AlitripbtripflightdistributionchangenewpayAPIRequest) SetParamBtripFlightModifyPayRq(_paramBtripFlightModifyPayRq *BtripFlightModifyPayRq) error {
	r._paramBtripFlightModifyPayRq = _paramBtripFlightModifyPayRq
	r.Set("param_btrip_flight_modify_pay_rq", _paramBtripFlightModifyPayRq)
	return nil
}

// GetParamBtripFlightModifyPayRq ParamBtripFlightModifyPayRq Getter
func (r AlitripbtripflightdistributionchangenewpayAPIRequest) GetParamBtripFlightModifyPayRq() *BtripFlightModifyPayRq {
	return r._paramBtripFlightModifyPayRq
}
