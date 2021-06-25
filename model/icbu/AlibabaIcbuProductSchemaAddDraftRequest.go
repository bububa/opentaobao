package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）ICBU商品发布草稿 APIRequest
alibaba.icbu.product.schema.add.draft

提供发布ICBU商品草稿的入口
*/
type AlibabaIcbuProductSchemaAddDraftRequest struct {
    model.Params

    // 发布入参
    paramProductTopPublishRequest   *ProductTopPublishRequest 

}

func NewAlibabaIcbuProductSchemaAddDraftRequest() *AlibabaIcbuProductSchemaAddDraftRequest{
    return &AlibabaIcbuProductSchemaAddDraftRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuProductSchemaAddDraftRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.schema.add.draft"
}

func (r AlibabaIcbuProductSchemaAddDraftRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuProductSchemaAddDraftRequest) SetParamProductTopPublishRequest(paramProductTopPublishRequest *ProductTopPublishRequest) error {
    r.paramProductTopPublishRequest = paramProductTopPublishRequest
    r.Set("param_product_top_publish_request", paramProductTopPublishRequest)
    return nil
}

func (r AlibabaIcbuProductSchemaAddDraftRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
    return r.paramProductTopPublishRequest
}

