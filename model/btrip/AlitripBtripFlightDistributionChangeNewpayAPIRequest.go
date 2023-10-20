package btrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripFlightDistributionChangeNewpayAPIRequest) Reset() {
	r._paramBtripFlightModifyPayRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionChangeNewpayAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.change.newpay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionChangeNewpayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionChangeNewpayAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripBtripFlightDistributionChangeNewpayAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripFlightDistributionChangeNewpayRequest()
	},
}

// GetAlitripBtripFlightDistributionChangeNewpayRequest 从 sync.Pool 获取 AlitripBtripFlightDistributionChangeNewpayAPIRequest
func GetAlitripBtripFlightDistributionChangeNewpayAPIRequest() *AlitripBtripFlightDistributionChangeNewpayAPIRequest {
	return poolAlitripBtripFlightDistributionChangeNewpayAPIRequest.Get().(*AlitripBtripFlightDistributionChangeNewpayAPIRequest)
}

// ReleaseAlitripBtripFlightDistributionChangeNewpayAPIRequest 将 AlitripBtripFlightDistributionChangeNewpayAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripFlightDistributionChangeNewpayAPIRequest(v *AlitripBtripFlightDistributionChangeNewpayAPIRequest) {
	v.Reset()
	poolAlitripBtripFlightDistributionChangeNewpayAPIRequest.Put(v)
}
