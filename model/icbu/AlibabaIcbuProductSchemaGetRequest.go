package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）ICBU商品发布schema接口 APIRequest
alibaba.icbu.product.schema.get

获取ICBU商品发布的页面规则和填写字段，适用于新发商品
*/
type AlibabaIcbuProductSchemaGetRequest struct {
    model.Params

    // 商品规则渲染请求
    paramProductTopPublishRequest   *ProductTopPublishRequest 

}

func NewAlibabaIcbuProductSchemaGetRequest() *AlibabaIcbuProductSchemaGetRequest{
    return &AlibabaIcbuProductSchemaGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuProductSchemaGetRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.schema.get"
}

func (r AlibabaIcbuProductSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuProductSchemaGetRequest) SetParamProductTopPublishRequest(paramProductTopPublishRequest *ProductTopPublishRequest) error {
    r.paramProductTopPublishRequest = paramProductTopPublishRequest
    r.Set("param_product_top_publish_request", paramProductTopPublishRequest)
    return nil
}

func (r AlibabaIcbuProductSchemaGetRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
    return r.paramProductTopPublishRequest
}

