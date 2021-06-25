package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品更新接口 APIRequest
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

func NewAlibabaGpuSchemaUpdateRequest() *AlibabaGpuSchemaUpdateRequest{
    return &AlibabaGpuSchemaUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaGpuSchemaUpdateRequest) GetApiMethodName() string {
    return "alibaba.gpu.schema.update"
}

func (r AlibabaGpuSchemaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaGpuSchemaUpdateRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r AlibabaGpuSchemaUpdateRequest) GetProductId() int64 {
    return r.productId
}

func (r *AlibabaGpuSchemaUpdateRequest) SetSchemaXmlFields(schemaXmlFields string) error {
    r.schemaXmlFields = schemaXmlFields
    r.Set("schema_xml_fields", schemaXmlFields)
    return nil
}

func (r AlibabaGpuSchemaUpdateRequest) GetSchemaXmlFields() string {
    return r.schemaXmlFields
}

func (r *AlibabaGpuSchemaUpdateRequest) SetProviderId(providerId int64) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

func (r AlibabaGpuSchemaUpdateRequest) GetProviderId() int64 {
    return r.providerId
}

