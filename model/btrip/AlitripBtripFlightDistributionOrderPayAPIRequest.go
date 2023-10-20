package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionorderpayAPIRequest 商旅机票分销-订单支付 API请求
// alitrip.btrip.flight.distribution.order.pay
//
// 商旅机票分销订单支付
type AlitripbtripflightdistributionorderpayAPIRequest struct {
	model.Params
	// 支付参数
	_paramBtripFlightPayOrderRq *BtripFlightPayOrderRq
}

// NewAlitripbtripflightdistributionorderpayRequest 初始化AlitripbtripflightdistributionorderpayAPIRequest对象
func NewAlitripbtripflightdistributionorderpayRequest() *AlitripbtripflightdistributionorderpayAPIRequest {
	return &AlitripbtripflightdistributionorderpayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionorderpayAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.order.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionorderpayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionorderpayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightPayOrderRq is ParamBtripFlightPayOrderRq Setter
// 支付参数
func (r *AlitripbtripflightdistributionorderpayAPIRequest) SetParamBtripFlightPayOrderRq(_paramBtripFlightPayOrderRq *BtripFlightPayOrderRq) error {
	r._paramBtripFlightPayOrderRq = _paramBtripFlightPayOrderRq
	r.Set("param_btrip_flight_pay_order_rq", _paramBtripFlightPayOrderRq)
	return nil
}

// GetParamBtripFlightPayOrderRq ParamBtripFlightPayOrderRq Getter
func (r AlitripbtripflightdistributionorderpayAPIRequest) GetParamBtripFlightPayOrderRq() *BtripFlightPayOrderRq {
	return r._paramBtripFlightPayOrderRq
}
