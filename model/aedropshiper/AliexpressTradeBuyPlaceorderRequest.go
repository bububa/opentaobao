package aedropshiper

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AE下单API API请求
aliexpress.trade.buy.placeorder

A006_INVALID_ACCOUNT_INFO
*/
type AliexpressTradeBuyPlaceorderAPIRequest struct {
    model.Params
    // 下单具体参数
    _paramPlaceOrderRequest4OpenApiDTO   *PlaceOrderRequest4OpenApiDTO
}

// 初始化AliexpressTradeBuyPlaceorderAPIRequest对象
func NewAliexpressTradeBuyPlaceorderRequest() *AliexpressTradeBuyPlaceorderAPIRequest{
    return &AliexpressTradeBuyPlaceorderAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressTradeBuyPlaceorderAPIRequest) GetApiMethodName() string {
    return "aliexpress.trade.buy.placeorder"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressTradeBuyPlaceorderAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamPlaceOrderRequest4OpenApiDTO Setter
// 下单具体参数
func (r *AliexpressTradeBuyPlaceorderAPIRequest) SetParamPlaceOrderRequest4OpenApiDTO(_paramPlaceOrderRequest4OpenApiDTO *PlaceOrderRequest4OpenApiDTO) error {
    r._paramPlaceOrderRequest4OpenApiDTO = _paramPlaceOrderRequest4OpenApiDTO
    r.Set("param_place_order_request4_open_api_d_t_o", _paramPlaceOrderRequest4OpenApiDTO)
    return nil
}

// ParamPlaceOrderRequest4OpenApiDTO Getter
func (r AliexpressTradeBuyPlaceorderAPIRequest) GetParamPlaceOrderRequest4OpenApiDTO() *PlaceOrderRequest4OpenApiDTO {
    return r._paramPlaceOrderRequest4OpenApiDTO
}
