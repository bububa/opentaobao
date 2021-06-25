package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）渲染草稿商品数据 APIRequest
alibaba.icbu.product.schema.render.draft

获取ICBU商品发布的字段填写规则和单个商品对应填写数据，适用于单个草稿商品编辑场景，
*/
type AlibabaIcbuProductSchemaRenderDraftRequest struct {
    model.Params

    // 商品规则渲染请求
    paramProductTopPublishRequest   *ProductTopPublishRequest 

}

func NewAlibabaIcbuProductSchemaRenderDraftRequest() *AlibabaIcbuProductSchemaRenderDraftRequest{
    return &AlibabaIcbuProductSchemaRenderDraftRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuProductSchemaRenderDraftRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.schema.render.draft"
}

func (r AlibabaIcbuProductSchemaRenderDraftRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuProductSchemaRenderDraftRequest) SetParamProductTopPublishRequest(paramProductTopPublishRequest *ProductTopPublishRequest) error {
    r.paramProductTopPublishRequest = paramProductTopPublishRequest
    r.Set("param_product_top_publish_request", paramProductTopPublishRequest)
    return nil
}

func (r AlibabaIcbuProductSchemaRenderDraftRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
    return r.paramProductTopPublishRequest
}

