package icbudropshipping

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴买家buynow下单接口 API请求
alibaba.buynow.order.create

阿里巴巴买家下单接口
*/
type AlibabaBuynowOrderCreateRequest struct {
    model.Params
    // Order creation parameter
    _paramOrderCreateRequest   *OrderCreateRequest
}

// 初始化AlibabaBuynowOrderCreateRequest对象
func NewAlibabaBuynowOrderCreateRequest() *AlibabaBuynowOrderCreateRequest{
    return &AlibabaBuynowOrderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaBuynowOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.buynow.order.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaBuynowOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamOrderCreateRequest Setter
// Order creation parameter
func (r *AlibabaBuynowOrderCreateRequest) SetParamOrderCreateRequest(_paramOrderCreateRequest *OrderCreateRequest) error {
    r._paramOrderCreateRequest = _paramOrderCreateRequest
    r.Set("param_order_create_request", _paramOrderCreateRequest)
    return nil
}

// ParamOrderCreateRequest Getter
func (r AlibabaBuynowOrderCreateRequest) GetParamOrderCreateRequest() *OrderCreateRequest {
    return r._paramOrderCreateRequest
}
