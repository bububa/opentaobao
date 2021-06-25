package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）商品发布新接口 APIRequest
alibaba.icbu.product.schema.add

提供发布ICBU商品的入口
*/
type AlibabaIcbuProductSchemaAddRequest struct {
    model.Params

    // 发布入参
    paramProductTopPublishRequest   *ProductTopPublishRequest 

}

func NewAlibabaIcbuProductSchemaAddRequest() *AlibabaIcbuProductSchemaAddRequest{
    return &AlibabaIcbuProductSchemaAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuProductSchemaAddRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.schema.add"
}

func (r AlibabaIcbuProductSchemaAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuProductSchemaAddRequest) SetParamProductTopPublishRequest(paramProductTopPublishRequest *ProductTopPublishRequest) error {
    r.paramProductTopPublishRequest = paramProductTopPublishRequest
    r.Set("param_product_top_publish_request", paramProductTopPublishRequest)
    return nil
}

func (r AlibabaIcbuProductSchemaAddRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
    return r.paramProductTopPublishRequest
}

