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
type AlibabaGpuSchemaUpdateAPIRequest struct {
    model.Params
    // 产品ID
    _productId   int64
    // 更新产品提交的schema数据
    _schemaXmlFields   string
    // 当前用户所在渠道如0代表天猫，8代表淘宝
    _providerId   int64
}

// 初始化AlibabaGpuSchemaUpdateAPIRequest对象
func NewAlibabaGpuSchemaUpdateRequest() *AlibabaGpuSchemaUpdateAPIRequest{
    return &AlibabaGpuSchemaUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaGpuSchemaUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.gpu.schema.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaGpuSchemaUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品ID
func (r *AlibabaGpuSchemaUpdateAPIRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r AlibabaGpuSchemaUpdateAPIRequest) GetProductId() int64 {
    return r._productId
}
// SchemaXmlFields Setter
// 更新产品提交的schema数据
func (r *AlibabaGpuSchemaUpdateAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
    r._schemaXmlFields = _schemaXmlFields
    r.Set("schema_xml_fields", _schemaXmlFields)
    return nil
}

// SchemaXmlFields Getter
func (r AlibabaGpuSchemaUpdateAPIRequest) GetSchemaXmlFields() string {
    return r._schemaXmlFields
}
// ProviderId Setter
// 当前用户所在渠道如0代表天猫，8代表淘宝
func (r *AlibabaGpuSchemaUpdateAPIRequest) SetProviderId(_providerId int64) error {
    r._providerId = _providerId
    r.Set("provider_id", _providerId)
    return nil
}

// ProviderId Getter
func (r AlibabaGpuSchemaUpdateAPIRequest) GetProviderId() int64 {
    return r._providerId
}
