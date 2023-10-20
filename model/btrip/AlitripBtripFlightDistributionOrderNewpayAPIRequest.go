package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionordernewpayAPIRequest 商旅机票分销-订单支付V2 API请求
// alitrip.btrip.flight.distribution.order.newpay
//
// 商旅机票分销-订单支付V2
type AlitripbtripflightdistributionordernewpayAPIRequest struct {
	model.Params
	// 支付参数
	_paramBtripFlightPayOrderRq *BtripFlightPayOrderRq
}

// NewAlitripbtripflightdistributionordernewpayRequest 初始化AlitripbtripflightdistributionordernewpayAPIRequest对象
func NewAlitripbtripflightdistributionordernewpayRequest() *AlitripbtripflightdistributionordernewpayAPIRequest {
	return &AlitripbtripflightdistributionordernewpayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionordernewpayAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.order.newpay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionordernewpayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionordernewpayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightPayOrderRq is ParamBtripFlightPayOrderRq Setter
// 支付参数
func (r *AlitripbtripflightdistributionordernewpayAPIRequest) SetParamBtripFlightPayOrderRq(_paramBtripFlightPayOrderRq *BtripFlightPayOrderRq) error {
	r._paramBtripFlightPayOrderRq = _paramBtripFlightPayOrderRq
	r.Set("param_btrip_flight_pay_order_rq", _paramBtripFlightPayOrderRq)
	return nil
}

// GetParamBtripFlightPayOrderRq ParamBtripFlightPayOrderRq Getter
func (r AlitripbtripflightdistributionordernewpayAPIRequest) GetParamBtripFlightPayOrderRq() *BtripFlightPayOrderRq {
	return r._paramBtripFlightPayOrderRq
}
