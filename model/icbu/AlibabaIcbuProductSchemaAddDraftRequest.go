package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）ICBU商品发布草稿 API请求
alibaba.icbu.product.schema.add.draft

提供发布ICBU商品草稿的入口
*/
type AlibabaIcbuProductSchemaAddDraftRequest struct {
    model.Params
    // 发布入参
    _paramProductTopPublishRequest   *ProductTopPublishRequest
}

// 初始化AlibabaIcbuProductSchemaAddDraftRequest对象
func NewAlibabaIcbuProductSchemaAddDraftRequest() *AlibabaIcbuProductSchemaAddDraftRequest{
    return &AlibabaIcbuProductSchemaAddDraftRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductSchemaAddDraftRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.schema.add.draft"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductSchemaAddDraftRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamProductTopPublishRequest Setter
// 发布入参
func (r *AlibabaIcbuProductSchemaAddDraftRequest) SetParamProductTopPublishRequest(_paramProductTopPublishRequest *ProductTopPublishRequest) error {
    r._paramProductTopPublishRequest = _paramProductTopPublishRequest
    r.Set("param_product_top_publish_request", _paramProductTopPublishRequest)
    return nil
}

// ParamProductTopPublishRequest Getter
func (r AlibabaIcbuProductSchemaAddDraftRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
    return r._paramProductTopPublishRequest
}
