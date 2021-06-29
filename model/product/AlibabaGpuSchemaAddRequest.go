package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
使用schema文件发布产品 API请求
alibaba.gpu.schema.add

使用Schema文件发布一个产品
*/
type AlibabaGpuSchemaAddRequest struct {
    model.Params
    // 叶子类目ID
    _leafCatId   int64
    // 品牌ID
    _brandId   int64
    // 根据alibaba.gpu.add.schema.get获取的规则提交上来的schema
    _schemaXmlFields   string
    // 当前用户所在渠道如0代表天猫，8代表淘宝
    _providerId   int64
}

// 初始化AlibabaGpuSchemaAddRequest对象
func NewAlibabaGpuSchemaAddRequest() *AlibabaGpuSchemaAddRequest{
    return &AlibabaGpuSchemaAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaGpuSchemaAddRequest) GetApiMethodName() string {
    return "alibaba.gpu.schema.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaGpuSchemaAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LeafCatId Setter
// 叶子类目ID
func (r *AlibabaGpuSchemaAddRequest) SetLeafCatId(_leafCatId int64) error {
    r._leafCatId = _leafCatId
    r.Set("leaf_cat_id", _leafCatId)
    return nil
}

// LeafCatId Getter
func (r AlibabaGpuSchemaAddRequest) GetLeafCatId() int64 {
    return r._leafCatId
}
// BrandId Setter
// 品牌ID
func (r *AlibabaGpuSchemaAddRequest) SetBrandId(_brandId int64) error {
    r._brandId = _brandId
    r.Set("brand_id", _brandId)
    return nil
}

// BrandId Getter
func (r AlibabaGpuSchemaAddRequest) GetBrandId() int64 {
    return r._brandId
}
// SchemaXmlFields Setter
// 根据alibaba.gpu.add.schema.get获取的规则提交上来的schema
func (r *AlibabaGpuSchemaAddRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
    r._schemaXmlFields = _schemaXmlFields
    r.Set("schema_xml_fields", _schemaXmlFields)
    return nil
}

// SchemaXmlFields Getter
func (r AlibabaGpuSchemaAddRequest) GetSchemaXmlFields() string {
    return r._schemaXmlFields
}
// ProviderId Setter
// 当前用户所在渠道如0代表天猫，8代表淘宝
func (r *AlibabaGpuSchemaAddRequest) SetProviderId(_providerId int64) error {
    r._providerId = _providerId
    r.Set("provider_id", _providerId)
    return nil
}

// ProviderId Getter
func (r AlibabaGpuSchemaAddRequest) GetProviderId() int64 {
    return r._providerId
}
