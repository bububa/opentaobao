package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）获取商品信息 API请求
alibaba.icbu.product.schema.render

获取ICBU商品发布的字段填写规则和单个商品对应填写数据，适用于单个商品编辑场景，不包括草稿。
*/
type AlibabaIcbuProductSchemaRenderRequest struct {
    model.Params
    // 商品规则渲染请求
    paramProductTopPublishRequest   *ProductTopPublishRequest
}

// 初始化AlibabaIcbuProductSchemaRenderRequest对象
func NewAlibabaIcbuProductSchemaRenderRequest() *AlibabaIcbuProductSchemaRenderRequest{
    return &AlibabaIcbuProductSchemaRenderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductSchemaRenderRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.schema.render"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductSchemaRenderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamProductTopPublishRequest Setter
// 商品规则渲染请求
func (r *AlibabaIcbuProductSchemaRenderRequest) SetParamProductTopPublishRequest(paramProductTopPublishRequest *ProductTopPublishRequest) error {
    r.paramProductTopPublishRequest = paramProductTopPublishRequest
    r.Set("param_product_top_publish_request", paramProductTopPublishRequest)
    return nil
}

// ParamProductTopPublishRequest Getter
func (r AlibabaIcbuProductSchemaRenderRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
    return r.paramProductTopPublishRequest
}
