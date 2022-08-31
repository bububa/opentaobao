package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionOrderCreateAPIRequest 商旅机票分销-创建订单 API请求
// alitrip.btrip.flight.distribution.order.create
//
// 商旅机票分销创建订单接口
type AlitripBtripFlightDistributionOrderCreateAPIRequest struct {
	model.Params
	// 提交订单参数
	_paramBtripFlightCreateOrderRq *BtripFlightCreateOrderRq
}

// NewAlitripBtripFlightDistributionOrderCreateRequest 初始化AlitripBtripFlightDistributionOrderCreateAPIRequest对象
func NewAlitripBtripFlightDistributionOrderCreateRequest() *AlitripBtripFlightDistributionOrderCreateAPIRequest {
	return &AlitripBtripFlightDistributionOrderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionOrderCreateAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionOrderCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamBtripFlightCreateOrderRq is ParamBtripFlightCreateOrderRq Setter
// 提交订单参数
func (r *AlitripBtripFlightDistributionOrderCreateAPIRequest) SetParamBtripFlightCreateOrderRq(_paramBtripFlightCreateOrderRq *BtripFlightCreateOrderRq) error {
	r._paramBtripFlightCreateOrderRq = _paramBtripFlightCreateOrderRq
	r.Set("param_btrip_flight_create_order_rq", _paramBtripFlightCreateOrderRq)
	return nil
}

// GetParamBtripFlightCreateOrderRq ParamBtripFlightCreateOrderRq Getter
func (r AlitripBtripFlightDistributionOrderCreateAPIRequest) GetParamBtripFlightCreateOrderRq() *BtripFlightCreateOrderRq {
	return r._paramBtripFlightCreateOrderRq
}
