package btrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripFlightDistributionOrderPayAPIRequest) Reset() {
	r._paramBtripFlightPayOrderRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionOrderPayAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.order.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionOrderPayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionOrderPayAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripBtripFlightDistributionOrderPayAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripFlightDistributionOrderPayRequest()
	},
}

// GetAlitripBtripFlightDistributionOrderPayRequest 从 sync.Pool 获取 AlitripBtripFlightDistributionOrderPayAPIRequest
func GetAlitripBtripFlightDistributionOrderPayAPIRequest() *AlitripBtripFlightDistributionOrderPayAPIRequest {
	return poolAlitripBtripFlightDistributionOrderPayAPIRequest.Get().(*AlitripBtripFlightDistributionOrderPayAPIRequest)
}

// ReleaseAlitripBtripFlightDistributionOrderPayAPIRequest 将 AlitripBtripFlightDistributionOrderPayAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripFlightDistributionOrderPayAPIRequest(v *AlitripBtripFlightDistributionOrderPayAPIRequest) {
	v.Reset()
	poolAlitripBtripFlightDistributionOrderPayAPIRequest.Put(v)
}
