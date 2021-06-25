package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）获取商品信息 APIRequest
alibaba.icbu.product.schema.render

获取ICBU商品发布的字段填写规则和单个商品对应填写数据，适用于单个商品编辑场景，不包括草稿。
*/
type AlibabaIcbuProductSchemaRenderRequest struct {
    model.Params

    // 商品规则渲染请求
    paramProductTopPublishRequest   *ProductTopPublishRequest 

}

func NewAlibabaIcbuProductSchemaRenderRequest() *AlibabaIcbuProductSchemaRenderRequest{
    return &AlibabaIcbuProductSchemaRenderRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuProductSchemaRenderRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.schema.render"
}

func (r AlibabaIcbuProductSchemaRenderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuProductSchemaRenderRequest) SetParamProductTopPublishRequest(paramProductTopPublishRequest *ProductTopPublishRequest) error {
    r.paramProductTopPublishRequest = paramProductTopPublishRequest
    r.Set("param_product_top_publish_request", paramProductTopPublishRequest)
    return nil
}

func (r AlibabaIcbuProductSchemaRenderRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
    return r.paramProductTopPublishRequest
}

