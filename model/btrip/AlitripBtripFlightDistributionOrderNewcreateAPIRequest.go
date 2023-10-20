package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionOrderNewcreateAPIRequest 商旅机票分销-创建订单V2 API请求
// alitrip.btrip.flight.distribution.order.newcreate
//
// 商旅机票分销-创建订单V2
type AlitripBtripFlightDistributionOrderNewcreateAPIRequest struct {
	model.Params
	// 提交订单参数
	_paramBtripFlightCreateOrderRq *BtripFlightCreateOrderRq
}

// NewAlitripBtripFlightDistributionOrderNewcreateRequest 初始化AlitripBtripFlightDistributionOrderNewcreateAPIRequest对象
func NewAlitripBtripFlightDistributionOrderNewcreateRequest() *AlitripBtripFlightDistributionOrderNewcreateAPIRequest {
	return &AlitripBtripFlightDistributionOrderNewcreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripFlightDistributionOrderNewcreateAPIRequest) Reset() {
	r._paramBtripFlightCreateOrderRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionOrderNewcreateAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.order.newcreate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionOrderNewcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripFlightDistributionOrderNewcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightCreateOrderRq is ParamBtripFlightCreateOrderRq Setter
// 提交订单参数
func (r *AlitripBtripFlightDistributionOrderNewcreateAPIRequest) SetParamBtripFlightCreateOrderRq(_paramBtripFlightCreateOrderRq *BtripFlightCreateOrderRq) error {
	r._paramBtripFlightCreateOrderRq = _paramBtripFlightCreateOrderRq
	r.Set("param_btrip_flight_create_order_rq", _paramBtripFlightCreateOrderRq)
	return nil
}

// GetParamBtripFlightCreateOrderRq ParamBtripFlightCreateOrderRq Getter
func (r AlitripBtripFlightDistributionOrderNewcreateAPIRequest) GetParamBtripFlightCreateOrderRq() *BtripFlightCreateOrderRq {
	return r._paramBtripFlightCreateOrderRq
}

var poolAlitripBtripFlightDistributionOrderNewcreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripFlightDistributionOrderNewcreateRequest()
	},
}

// GetAlitripBtripFlightDistributionOrderNewcreateRequest 从 sync.Pool 获取 AlitripBtripFlightDistributionOrderNewcreateAPIRequest
func GetAlitripBtripFlightDistributionOrderNewcreateAPIRequest() *AlitripBtripFlightDistributionOrderNewcreateAPIRequest {
	return poolAlitripBtripFlightDistributionOrderNewcreateAPIRequest.Get().(*AlitripBtripFlightDistributionOrderNewcreateAPIRequest)
}

// ReleaseAlitripBtripFlightDistributionOrderNewcreateAPIRequest 将 AlitripBtripFlightDistributionOrderNewcreateAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripFlightDistributionOrderNewcreateAPIRequest(v *AlitripBtripFlightDistributionOrderNewcreateAPIRequest) {
	v.Reset()
	poolAlitripBtripFlightDistributionOrderNewcreateAPIRequest.Put(v)
}
