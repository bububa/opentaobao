package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）渲染草稿商品数据 API请求
alibaba.icbu.product.schema.render.draft

获取ICBU商品发布的字段填写规则和单个商品对应填写数据，适用于单个草稿商品编辑场景，
*/
type AlibabaIcbuProductSchemaRenderDraftRequest struct {
    model.Params
    // 商品规则渲染请求
    paramProductTopPublishRequest   *ProductTopPublishRequest
}

// 初始化AlibabaIcbuProductSchemaRenderDraftRequest对象
func NewAlibabaIcbuProductSchemaRenderDraftRequest() *AlibabaIcbuProductSchemaRenderDraftRequest{
    return &AlibabaIcbuProductSchemaRenderDraftRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductSchemaRenderDraftRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.schema.render.draft"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductSchemaRenderDraftRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamProductTopPublishRequest Setter
// 商品规则渲染请求
func (r *AlibabaIcbuProductSchemaRenderDraftRequest) SetParamProductTopPublishRequest(paramProductTopPublishRequest *ProductTopPublishRequest) error {
    r.paramProductTopPublishRequest = paramProductTopPublishRequest
    r.Set("param_product_top_publish_request", paramProductTopPublishRequest)
    return nil
}

// ParamProductTopPublishRequest Getter
func (r AlibabaIcbuProductSchemaRenderDraftRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
    return r.paramProductTopPublishRequest
}
