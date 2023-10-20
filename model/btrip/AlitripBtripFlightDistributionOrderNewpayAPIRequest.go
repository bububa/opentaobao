package btrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripFlightDistributionOrderNewpayAPIRequest) Reset() {
	r._paramBtripFlightPayOrderRq = nil
	r.Params.ToZero()
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

var poolAlitripBtripFlightDistributionOrderNewpayAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripFlightDistributionOrderNewpayRequest()
	},
}

// GetAlitripBtripFlightDistributionOrderNewpayRequest 从 sync.Pool 获取 AlitripBtripFlightDistributionOrderNewpayAPIRequest
func GetAlitripBtripFlightDistributionOrderNewpayAPIRequest() *AlitripBtripFlightDistributionOrderNewpayAPIRequest {
	return poolAlitripBtripFlightDistributionOrderNewpayAPIRequest.Get().(*AlitripBtripFlightDistributionOrderNewpayAPIRequest)
}

// ReleaseAlitripBtripFlightDistributionOrderNewpayAPIRequest 将 AlitripBtripFlightDistributionOrderNewpayAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripFlightDistributionOrderNewpayAPIRequest(v *AlitripBtripFlightDistributionOrderNewpayAPIRequest) {
	v.Reset()
	poolAlitripBtripFlightDistributionOrderNewpayAPIRequest.Put(v)
}
