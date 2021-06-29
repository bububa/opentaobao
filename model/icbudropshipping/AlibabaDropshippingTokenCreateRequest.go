package icbudropshipping

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际站dropshipping 选品token 创建 APIRequest
alibaba.dropshipping.token.create

国际站dropshipping 选品token 创建，用于让买家有权限访问我们指定的 商品场馆
*/
type AlibabaDropshippingTokenCreateRequest struct {
    model.Params

}

func NewAlibabaDropshippingTokenCreateRequest() *AlibabaDropshippingTokenCreateRequest{
    return &AlibabaDropshippingTokenCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDropshippingTokenCreateRequest) GetApiMethodName() string {
    return "alibaba.dropshipping.token.create"
}

func (r AlibabaDropshippingTokenCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


