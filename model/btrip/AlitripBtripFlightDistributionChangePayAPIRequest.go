package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionchangepayAPIRequest 商旅机票改签支付 API请求
// alitrip.btrip.flight.distribution.change.pay
//
// 改签订单支付
type AlitripbtripflightdistributionchangepayAPIRequest struct {
	model.Params
	// 改签支付入参
	_paramBtripFlightModifyPayRq *BtripFlightModifyPayRq
}

// NewAlitripbtripflightdistributionchangepayRequest 初始化AlitripbtripflightdistributionchangepayAPIRequest对象
func NewAlitripbtripflightdistributionchangepayRequest() *AlitripbtripflightdistributionchangepayAPIRequest {
	return &AlitripbtripflightdistributionchangepayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionchangepayAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionchangepayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionchangepayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightModifyPayRq is ParamBtripFlightModifyPayRq Setter
// 改签支付入参
func (r *AlitripbtripflightdistributionchangepayAPIRequest) SetParamBtripFlightModifyPayRq(_paramBtripFlightModifyPayRq *BtripFlightModifyPayRq) error {
	r._paramBtripFlightModifyPayRq = _paramBtripFlightModifyPayRq
	r.Set("param_btrip_flight_modify_pay_rq", _paramBtripFlightModifyPayRq)
	return nil
}

// GetParamBtripFlightModifyPayRq ParamBtripFlightModifyPayRq Getter
func (r AlitripbtripflightdistributionchangepayAPIRequest) GetParamBtripFlightModifyPayRq() *BtripFlightModifyPayRq {
	return r._paramBtripFlightModifyPayRq
}
