package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionordercreateAPIRequest 商旅机票分销-创建订单 API请求
// alitrip.btrip.flight.distribution.order.create
//
// 商旅机票分销创建订单接口
type AlitripbtripflightdistributionordercreateAPIRequest struct {
	model.Params
	// 提交订单参数
	_paramBtripFlightCreateOrderRq *BtripFlightCreateOrderRq
}

// NewAlitripbtripflightdistributionordercreateRequest 初始化AlitripbtripflightdistributionordercreateAPIRequest对象
func NewAlitripbtripflightdistributionordercreateRequest() *AlitripbtripflightdistributionordercreateAPIRequest {
	return &AlitripbtripflightdistributionordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightdistributionordercreateAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightdistributionordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightdistributionordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripFlightCreateOrderRq is ParamBtripFlightCreateOrderRq Setter
// 提交订单参数
func (r *AlitripbtripflightdistributionordercreateAPIRequest) SetParamBtripFlightCreateOrderRq(_paramBtripFlightCreateOrderRq *BtripFlightCreateOrderRq) error {
	r._paramBtripFlightCreateOrderRq = _paramBtripFlightCreateOrderRq
	r.Set("param_btrip_flight_create_order_rq", _paramBtripFlightCreateOrderRq)
	return nil
}

// GetParamBtripFlightCreateOrderRq ParamBtripFlightCreateOrderRq Getter
func (r AlitripbtripflightdistributionordercreateAPIRequest) GetParamBtripFlightCreateOrderRq() *BtripFlightCreateOrderRq {
	return r._paramBtripFlightCreateOrderRq
}
