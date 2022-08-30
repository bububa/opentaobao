package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeNewpayAPIRequest 商旅机票改签支付V2 API请求
// alitrip.btrip.flight.distribution.change.newpay
//
// 商旅机票改签支付V2
type AlitripBtripFlightDistributionChangeNewpayAPIRequest struct {
	model.Params
	// 改签支付入参
	_paramBtripFlightModifyPayRq *BtripFlightModifyPayRq
}

// NewAlitripBtripFlightDistributionChangeNewpayRequest 初始化AlitripBtripFlightDistributionChangeNewpayAPIRequest对象
func NewAlitripBtripFlightDistributionChangeNewpayRequest() *AlitripBtripFlightDistributionChangeNewpayAPIRequest {
	return &AlitripBtripFlightDistributionChangeNewpayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionChangeNewpayAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.newpay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionChangeNewpayAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamBtripFlightModifyPayRq is ParamBtripFlightModifyPayRq Setter
// 改签支付入参
func (r *AlitripBtripFlightDistributionChangeNewpayAPIRequest) SetParamBtripFlightModifyPayRq(_paramBtripFlightModifyPayRq *BtripFlightModifyPayRq) error {
	r._paramBtripFlightModifyPayRq = _paramBtripFlightModifyPayRq
	r.Set("param_btrip_flight_modify_pay_rq", _paramBtripFlightModifyPayRq)
	return nil
}

// GetParamBtripFlightModifyPayRq ParamBtripFlightModifyPayRq Getter
func (r AlitripBtripFlightDistributionChangeNewpayAPIRequest) GetParamBtripFlightModifyPayRq() *BtripFlightModifyPayRq {
	return r._paramBtripFlightModifyPayRq
}
