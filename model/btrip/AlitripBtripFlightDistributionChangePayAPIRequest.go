package btrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripFlightDistributionChangePayAPIRequest) Reset() {
	r._paramBtripFlightModifyPayRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionChangePayAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionChangePayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionChangePayAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripBtripFlightDistributionChangePayAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripFlightDistributionChangePayRequest()
	},
}

// GetAlitripBtripFlightDistributionChangePayRequest 从 sync.Pool 获取 AlitripBtripFlightDistributionChangePayAPIRequest
func GetAlitripBtripFlightDistributionChangePayAPIRequest() *AlitripBtripFlightDistributionChangePayAPIRequest {
	return poolAlitripBtripFlightDistributionChangePayAPIRequest.Get().(*AlitripBtripFlightDistributionChangePayAPIRequest)
}

// ReleaseAlitripBtripFlightDistributionChangePayAPIRequest 将 AlitripBtripFlightDistributionChangePayAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripFlightDistributionChangePayAPIRequest(v *AlitripBtripFlightDistributionChangePayAPIRequest) {
	v.Reset()
	poolAlitripBtripFlightDistributionChangePayAPIRequest.Put(v)
}
