package icbudropshipping

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴买家buynow下单接口 APIRequest
alibaba.buynow.order.create

阿里巴巴买家下单接口
*/
type AlibabaBuynowOrderCreateRequest struct {
    model.Params

    // Order creation parameter
    paramOrderCreateRequest   *OrderCreateRequest 

}

func NewAlibabaBuynowOrderCreateRequest() *AlibabaBuynowOrderCreateRequest{
    return &AlibabaBuynowOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaBuynowOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.buynow.order.create"
}

func (r AlibabaBuynowOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaBuynowOrderCreateRequest) SetParamOrderCreateRequest(paramOrderCreateRequest *OrderCreateRequest) error {
    r.paramOrderCreateRequest = paramOrderCreateRequest
    r.Set("param_order_create_request", paramOrderCreateRequest)
    return nil
}

func (r AlibabaBuynowOrderCreateRequest) GetParamOrderCreateRequest() *OrderCreateRequest {
    return r.paramOrderCreateRequest
}

