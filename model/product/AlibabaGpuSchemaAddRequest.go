package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
使用schema文件发布产品 APIRequest
alibaba.gpu.schema.add

使用Schema文件发布一个产品
*/
type AlibabaGpuSchemaAddRequest struct {
    model.Params

    // 叶子类目ID
    leafCatId   int64 

    // 品牌ID
    brandId   int64 

    // 根据alibaba.gpu.add.schema.get获取的规则提交上来的schema
    schemaXmlFields   string 

    // 当前用户所在渠道如0代表天猫，8代表淘宝
    providerId   int64 

}

func NewAlibabaGpuSchemaAddRequest() *AlibabaGpuSchemaAddRequest{
    return &AlibabaGpuSchemaAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaGpuSchemaAddRequest) GetApiMethodName() string {
    return "alibaba.gpu.schema.add"
}

func (r AlibabaGpuSchemaAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaGpuSchemaAddRequest) SetLeafCatId(leafCatId int64) error {
    r.leafCatId = leafCatId
    r.Set("leaf_cat_id", leafCatId)
    return nil
}

func (r AlibabaGpuSchemaAddRequest) GetLeafCatId() int64 {
    return r.leafCatId
}

func (r *AlibabaGpuSchemaAddRequest) SetBrandId(brandId int64) error {
    r.brandId = brandId
    r.Set("brand_id", brandId)
    return nil
}

func (r AlibabaGpuSchemaAddRequest) GetBrandId() int64 {
    return r.brandId
}

func (r *AlibabaGpuSchemaAddRequest) SetSchemaXmlFields(schemaXmlFields string) error {
    r.schemaXmlFields = schemaXmlFields
    r.Set("schema_xml_fields", schemaXmlFields)
    return nil
}

func (r AlibabaGpuSchemaAddRequest) GetSchemaXmlFields() string {
    return r.schemaXmlFields
}

func (r *AlibabaGpuSchemaAddRequest) SetProviderId(providerId int64) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

func (r AlibabaGpuSchemaAddRequest) GetProviderId() int64 {
    return r.providerId
}

