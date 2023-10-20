package btrip

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
