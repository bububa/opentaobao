package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtriphoteldistributionordervalidateAPIRequest 商旅酒店API分销下单前校验 API请求
// alitrip.btrip.hotel.distribution.order.validate
//
// 商旅酒店API分销下单前校验
type AlitripbtriphoteldistributionordervalidateAPIRequest struct {
	model.Params
	// 下单前校验入参
	_paramBtripHotelValidateOrderRq *BtripHotelValidateOrderRq
}

// NewAlitripbtriphoteldistributionordervalidateRequest 初始化AlitripbtriphoteldistributionordervalidateAPIRequest对象
func NewAlitripbtriphoteldistributionordervalidateRequest() *AlitripbtriphoteldistributionordervalidateAPIRequest {
	return &AlitripbtriphoteldistributionordervalidateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtriphoteldistributionordervalidateAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.order.validate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtriphoteldistributionordervalidateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtriphoteldistributionordervalidateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBtripHotelValidateOrderRq is ParamBtripHotelValidateOrderRq Setter
// 下单前校验入参
func (r *AlitripbtriphoteldistributionordervalidateAPIRequest) SetParamBtripHotelValidateOrderRq(_paramBtripHotelValidateOrderRq *BtripHotelValidateOrderRq) error {
	r._paramBtripHotelValidateOrderRq = _paramBtripHotelValidateOrderRq
	r.Set("param_btrip_hotel_validate_order_rq", _paramBtripHotelValidateOrderRq)
	return nil
}

// GetParamBtripHotelValidateOrderRq ParamBtripHotelValidateOrderRq Getter
func (r AlitripbtriphoteldistributionordervalidateAPIRequest) GetParamBtripHotelValidateOrderRq() *BtripHotelValidateOrderRq {
	return r._paramBtripHotelValidateOrderRq
}
