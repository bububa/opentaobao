package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取产品编辑schema规则的接口 API请求
alibaba.gpu.update.schema.get

获取产品编辑schema规则的接口
*/
type AlibabaGpuUpdateSchemaGetRequest struct {
    model.Params
    // 产品ID
    productId   int64
    // 当前用户所在渠道如0代表天猫，8代表淘宝
    providerId   int64
}

// 初始化AlibabaGpuUpdateSchemaGetRequest对象
func NewAlibabaGpuUpdateSchemaGetRequest() *AlibabaGpuUpdateSchemaGetRequest{
    return &AlibabaGpuUpdateSchemaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaGpuUpdateSchemaGetRequest) GetApiMethodName() string {
    return "alibaba.gpu.update.schema.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaGpuUpdateSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品ID
func (r *AlibabaGpuUpdateSchemaGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r AlibabaGpuUpdateSchemaGetRequest) GetProductId() int64 {
    return r.productId
}
// ProviderId Setter
// 当前用户所在渠道如0代表天猫，8代表淘宝
func (r *AlibabaGpuUpdateSchemaGetRequest) SetProviderId(providerId int64) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

// ProviderId Getter
func (r AlibabaGpuUpdateSchemaGetRequest) GetProviderId() int64 {
    return r.providerId
}
