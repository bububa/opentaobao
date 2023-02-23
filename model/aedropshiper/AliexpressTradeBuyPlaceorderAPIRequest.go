package aedropshiper

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressTradeBuyPlaceorderAPIRequest AE下单API API请求
// aliexpress.trade.buy.placeorder
//
// 150欧欧盟税改
type AliexpressTradeBuyPlaceorderAPIRequest struct {
	model.Params
	// 下单具体参数
	_paramPlaceOrderRequest4OpenApiDTO *PlaceOrderRequest4OpenApiDto
}

// NewAliexpressTradeBuyPlaceorderRequest 初始化AliexpressTradeBuyPlaceorderAPIRequest对象
func NewAliexpressTradeBuyPlaceorderRequest() *AliexpressTradeBuyPlaceorderAPIRequest {
	return &AliexpressTradeBuyPlaceorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressTradeBuyPlaceorderAPIRequest) GetApiMethodName() string {
	return "aliexpress.trade.buy.placeorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressTradeBuyPlaceorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressTradeBuyPlaceorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPlaceOrderRequest4OpenApiDTO is ParamPlaceOrderRequest4OpenApiDTO Setter
// 下单具体参数
func (r *AliexpressTradeBuyPlaceorderAPIRequest) SetParamPlaceOrderRequest4OpenApiDTO(_paramPlaceOrderRequest4OpenApiDTO *PlaceOrderRequest4OpenApiDto) error {
	r._paramPlaceOrderRequest4OpenApiDTO = _paramPlaceOrderRequest4OpenApiDTO
	r.Set("param_place_order_request4_open_api_d_t_o", _paramPlaceOrderRequest4OpenApiDTO)
	return nil
}

// GetParamPlaceOrderRequest4OpenApiDTO ParamPlaceOrderRequest4OpenApiDTO Getter
func (r AliexpressTradeBuyPlaceorderAPIRequest) GetParamPlaceOrderRequest4OpenApiDTO() *PlaceOrderRequest4OpenApiDto {
	return r._paramPlaceOrderRequest4OpenApiDTO
}
