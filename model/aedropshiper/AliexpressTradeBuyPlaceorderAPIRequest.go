package aedropshiper

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresstradebuyplaceorderAPIRequest AE下单API API请求
// aliexpress.trade.buy.placeorder
//
// 150欧欧盟税改
type AliexpresstradebuyplaceorderAPIRequest struct {
	model.Params
	// 下单具体参数
	_paramPlaceOrderRequest4OpenApiDTO *PlaceOrderRequest4openApiDto
}

// NewAliexpresstradebuyplaceorderRequest 初始化AliexpresstradebuyplaceorderAPIRequest对象
func NewAliexpresstradebuyplaceorderRequest() *AliexpresstradebuyplaceorderAPIRequest {
	return &AliexpresstradebuyplaceorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresstradebuyplaceorderAPIRequest) GetApiMethodName() string {
	return "aliexpress.trade.buy.placeorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresstradebuyplaceorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresstradebuyplaceorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPlaceOrderRequest4OpenApiDTO is ParamPlaceOrderRequest4OpenApiDTO Setter
// 下单具体参数
func (r *AliexpresstradebuyplaceorderAPIRequest) SetParamPlaceOrderRequest4OpenApiDTO(_paramPlaceOrderRequest4OpenApiDTO *PlaceOrderRequest4openApiDto) error {
	r._paramPlaceOrderRequest4OpenApiDTO = _paramPlaceOrderRequest4OpenApiDTO
	r.Set("param_place_order_request4_open_api_d_t_o", _paramPlaceOrderRequest4OpenApiDTO)
	return nil
}

// GetParamPlaceOrderRequest4OpenApiDTO ParamPlaceOrderRequest4OpenApiDTO Getter
func (r AliexpresstradebuyplaceorderAPIRequest) GetParamPlaceOrderRequest4OpenApiDTO() *PlaceOrderRequest4openApiDto {
	return r._paramPlaceOrderRequest4OpenApiDTO
}
