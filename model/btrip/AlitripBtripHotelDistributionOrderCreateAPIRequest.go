package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionOrderCreateAPIRequest 商旅酒店分销-创建订单 API请求
// alitrip.btrip.hotel.distribution.order.create
//
// 商旅酒店分销-创建订单
type AlitripBtripHotelDistributionOrderCreateAPIRequest struct {
	model.Params
	// 创建订单请求入参
	_paramBtripHotelCreateOrderRq *BtripHotelCreateOrderRq
}

// NewAlitripBtripHotelDistributionOrderCreateRequest 初始化AlitripBtripHotelDistributionOrderCreateAPIRequest对象
func NewAlitripBtripHotelDistributionOrderCreateRequest() *AlitripBtripHotelDistributionOrderCreateAPIRequest {
	return &AlitripBtripHotelDistributionOrderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionOrderCreateAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionOrderCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamBtripHotelCreateOrderRq is ParamBtripHotelCreateOrderRq Setter
// 创建订单请求入参
func (r *AlitripBtripHotelDistributionOrderCreateAPIRequest) SetParamBtripHotelCreateOrderRq(_paramBtripHotelCreateOrderRq *BtripHotelCreateOrderRq) error {
	r._paramBtripHotelCreateOrderRq = _paramBtripHotelCreateOrderRq
	r.Set("param_btrip_hotel_create_order_rq", _paramBtripHotelCreateOrderRq)
	return nil
}

// GetParamBtripHotelCreateOrderRq ParamBtripHotelCreateOrderRq Getter
func (r AlitripBtripHotelDistributionOrderCreateAPIRequest) GetParamBtripHotelCreateOrderRq() *BtripHotelCreateOrderRq {
	return r._paramBtripHotelCreateOrderRq
}
