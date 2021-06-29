package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品更新接口 API请求
alibaba.gpu.schema.update

产品更新接口
*/
type AlibabaGpuSchemaUpdateRequest struct {
    model.Params
    // 产品ID
    productId   int64
    // 更新产品提交的schema数据
    schemaXmlFields   string
    // 当前用户所在渠道如0代表天猫，8代表淘宝
    providerId   int64
}

// 初始化AlibabaGpuSchemaUpdateRequest对象
func NewAlibabaGpuSchemaUpdateRequest() *AlibabaGpuSchemaUpdateRequest{
    return &AlibabaGpuSchemaUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaGpuSchemaUpdateRequest) GetApiMethodName() string {
    return "alibaba.gpu.schema.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaGpuSchemaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品ID
func (r *AlibabaGpuSchemaUpdateRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r AlibabaGpuSchemaUpdateRequest) GetProductId() int64 {
    return r.productId
}
// SchemaXmlFields Setter
// 更新产品提交的schema数据
func (r *AlibabaGpuSchemaUpdateRequest) SetSchemaXmlFields(schemaXmlFields string) error {
    r.schemaXmlFields = schemaXmlFields
    r.Set("schema_xml_fields", schemaXmlFields)
    return nil
}

// SchemaXmlFields Getter
func (r AlibabaGpuSchemaUpdateRequest) GetSchemaXmlFields() string {
    return r.schemaXmlFields
}
// ProviderId Setter
// 当前用户所在渠道如0代表天猫，8代表淘宝
func (r *AlibabaGpuSchemaUpdateRequest) SetProviderId(providerId int64) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

// ProviderId Getter
func (r AlibabaGpuSchemaUpdateRequest) GetProviderId() int64 {
    return r.providerId
}
