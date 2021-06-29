package aedropshiper

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AE下单API APIRequest
aliexpress.trade.buy.placeorder

A006_INVALID_ACCOUNT_INFO
*/
type AliexpressTradeBuyPlaceorderRequest struct {
    model.Params

    // 下单具体参数
    paramPlaceOrderRequest4OpenApiDTO   *PlaceOrderRequest4OpenApiDto 

}

func NewAliexpressTradeBuyPlaceorderRequest() *AliexpressTradeBuyPlaceorderRequest{
    return &AliexpressTradeBuyPlaceorderRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressTradeBuyPlaceorderRequest) GetApiMethodName() string {
    return "aliexpress.trade.buy.placeorder"
}

func (r AliexpressTradeBuyPlaceorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressTradeBuyPlaceorderRequest) SetParamPlaceOrderRequest4OpenApiDTO(paramPlaceOrderRequest4OpenApiDTO *PlaceOrderRequest4OpenApiDto) error {
    r.paramPlaceOrderRequest4OpenApiDTO = paramPlaceOrderRequest4OpenApiDTO
    r.Set("param_place_order_request4_open_api_d_t_o", paramPlaceOrderRequest4OpenApiDTO)
    return nil
}

func (r AliexpressTradeBuyPlaceorderRequest) GetParamPlaceOrderRequest4OpenApiDTO() *PlaceOrderRequest4OpenApiDto {
    return r.paramPlaceOrderRequest4OpenApiDTO
}

