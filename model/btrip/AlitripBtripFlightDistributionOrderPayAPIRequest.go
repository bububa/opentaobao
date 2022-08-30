package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionOrderPayAPIRequest 商旅机票分销-订单支付 API请求
// alitrip.btrip.flight.distribution.order.pay
//
// 商旅机票分销订单支付
type AlitripBtripFlightDistributionOrderPayAPIRequest struct {
	model.Params
	// 支付参数
	_paramBtripFlightPayOrderRq *BtripFlightPayOrderRq
}

// NewAlitripBtripFlightDistributionOrderPayRequest 初始化AlitripBtripFlightDistributionOrderPayAPIRequest对象
func NewAlitripBtripFlightDistributionOrderPayRequest() *AlitripBtripFlightDistributionOrderPayAPIRequest {
	return &AlitripBtripFlightDistributionOrderPayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionOrderPayAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.order.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionOrderPayAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamBtripFlightPayOrderRq is ParamBtripFlightPayOrderRq Setter
// 支付参数
func (r *AlitripBtripFlightDistributionOrderPayAPIRequest) SetParamBtripFlightPayOrderRq(_paramBtripFlightPayOrderRq *BtripFlightPayOrderRq) error {
	r._paramBtripFlightPayOrderRq = _paramBtripFlightPayOrderRq
	r.Set("param_btrip_flight_pay_order_rq", _paramBtripFlightPayOrderRq)
	return nil
}

// GetParamBtripFlightPayOrderRq ParamBtripFlightPayOrderRq Getter
func (r AlitripBtripFlightDistributionOrderPayAPIRequest) GetParamBtripFlightPayOrderRq() *BtripFlightPayOrderRq {
	return r._paramBtripFlightPayOrderRq
}
