package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangePayAPIRequest 商旅机票改签支付 API请求
// alitrip.btrip.flight.distribution.change.pay
//
// 改签订单支付
type AlitripBtripFlightDistributionChangePayAPIRequest struct {
	model.Params
	// 改签支付入参
	_paramBtripFlightModifyPayRq *BtripFlightModifyPayRq
}

// NewAlitripBtripFlightDistributionChangePayRequest 初始化AlitripBtripFlightDistributionChangePayAPIRequest对象
func NewAlitripBtripFlightDistributionChangePayRequest() *AlitripBtripFlightDistributionChangePayAPIRequest {
	return &AlitripBtripFlightDistributionChangePayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionChangePayAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionChangePayAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamBtripFlightModifyPayRq is ParamBtripFlightModifyPayRq Setter
// 改签支付入参
func (r *AlitripBtripFlightDistributionChangePayAPIRequest) SetParamBtripFlightModifyPayRq(_paramBtripFlightModifyPayRq *BtripFlightModifyPayRq) error {
	r._paramBtripFlightModifyPayRq = _paramBtripFlightModifyPayRq
	r.Set("param_btrip_flight_modify_pay_rq", _paramBtripFlightModifyPayRq)
	return nil
}

// GetParamBtripFlightModifyPayRq ParamBtripFlightModifyPayRq Getter
func (r AlitripBtripFlightDistributionChangePayAPIRequest) GetParamBtripFlightModifyPayRq() *BtripFlightModifyPayRq {
	return r._paramBtripFlightModifyPayRq
}
