package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionOrderNewpayAPIRequest 商旅机票分销-订单支付V2 API请求
// alitrip.btrip.flight.distribution.order.newpay
//
// 商旅机票分销-订单支付V2
type AlitripBtripFlightDistributionOrderNewpayAPIRequest struct {
	model.Params
	// 支付参数
	_paramBtripFlightPayOrderRq *BtripFlightPayOrderRq
}

// NewAlitripBtripFlightDistributionOrderNewpayRequest 初始化AlitripBtripFlightDistributionOrderNewpayAPIRequest对象
func NewAlitripBtripFlightDistributionOrderNewpayRequest() *AlitripBtripFlightDistributionOrderNewpayAPIRequest {
	return &AlitripBtripFlightDistributionOrderNewpayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionOrderNewpayAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.order.newpay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionOrderNewpayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionOrderNewpayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightPayOrderRq is ParamBtripFlightPayOrderRq Setter
// 支付参数
func (r *AlitripBtripFlightDistributionOrderNewpayAPIRequest) SetParamBtripFlightPayOrderRq(_paramBtripFlightPayOrderRq *BtripFlightPayOrderRq) error {
	r._paramBtripFlightPayOrderRq = _paramBtripFlightPayOrderRq
	r.Set("param_btrip_flight_pay_order_rq", _paramBtripFlightPayOrderRq)
	return nil
}

// GetParamBtripFlightPayOrderRq ParamBtripFlightPayOrderRq Getter
func (r AlitripBtripFlightDistributionOrderNewpayAPIRequest) GetParamBtripFlightPayOrderRq() *BtripFlightPayOrderRq {
	return r._paramBtripFlightPayOrderRq
}
