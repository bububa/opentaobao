package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtriphoteldistributionordercreateAPIRequest 商旅酒店分销-创建订单 API请求
// alitrip.btrip.hotel.distribution.order.create
//
// 商旅酒店分销-创建订单
type AlitripbtriphoteldistributionordercreateAPIRequest struct {
	model.Params
	// 创建订单请求入参
	_paramBtripHotelCreateOrderRq *BtripHotelCreateOrderRq
}

// NewAlitripbtriphoteldistributionordercreateRequest 初始化AlitripbtriphoteldistributionordercreateAPIRequest对象
func NewAlitripbtriphoteldistributionordercreateRequest() *AlitripbtriphoteldistributionordercreateAPIRequest {
	return &AlitripbtriphoteldistributionordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtriphoteldistributionordercreateAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtriphoteldistributionordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtriphoteldistributionordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripHotelCreateOrderRq is ParamBtripHotelCreateOrderRq Setter
// 创建订单请求入参
func (r *AlitripbtriphoteldistributionordercreateAPIRequest) SetParamBtripHotelCreateOrderRq(_paramBtripHotelCreateOrderRq *BtripHotelCreateOrderRq) error {
	r._paramBtripHotelCreateOrderRq = _paramBtripHotelCreateOrderRq
	r.Set("param_btrip_hotel_create_order_rq", _paramBtripHotelCreateOrderRq)
	return nil
}

// GetParamBtripHotelCreateOrderRq ParamBtripHotelCreateOrderRq Getter
func (r AlitripbtriphoteldistributionordercreateAPIRequest) GetParamBtripHotelCreateOrderRq() *BtripHotelCreateOrderRq {
	return r._paramBtripHotelCreateOrderRq
}
