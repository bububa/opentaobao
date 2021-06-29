package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）商品发布新接口 API请求
alibaba.icbu.product.schema.add

提供发布ICBU商品的入口
*/
type AlibabaIcbuProductSchemaAddRequest struct {
    model.Params
    // 发布入参
    paramProductTopPublishRequest   *ProductTopPublishRequest
}

// 初始化AlibabaIcbuProductSchemaAddRequest对象
func NewAlibabaIcbuProductSchemaAddRequest() *AlibabaIcbuProductSchemaAddRequest{
    return &AlibabaIcbuProductSchemaAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductSchemaAddRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.schema.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductSchemaAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamProductTopPublishRequest Setter
// 发布入参
func (r *AlibabaIcbuProductSchemaAddRequest) SetParamProductTopPublishRequest(paramProductTopPublishRequest *ProductTopPublishRequest) error {
    r.paramProductTopPublishRequest = paramProductTopPublishRequest
    r.Set("param_product_top_publish_request", paramProductTopPublishRequest)
    return nil
}

// ParamProductTopPublishRequest Getter
func (r AlibabaIcbuProductSchemaAddRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
    return r.paramProductTopPublishRequest
}
