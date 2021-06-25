package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取产品发布规则接口 APIRequest
alibaba.gpu.add.schema.get

获取产品发布规则接口
*/
type AlibabaGpuAddSchemaGetRequest struct {
    model.Params

    // 叶子类目ID
    leafCatId   int64 

    // 品牌ID
    brandId   int64 

    // 当前用户所在渠道如0代表天猫，8代表淘宝
    providerId   int64 

}

func NewAlibabaGpuAddSchemaGetRequest() *AlibabaGpuAddSchemaGetRequest{
    return &AlibabaGpuAddSchemaGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaGpuAddSchemaGetRequest) GetApiMethodName() string {
    return "alibaba.gpu.add.schema.get"
}

func (r AlibabaGpuAddSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaGpuAddSchemaGetRequest) SetLeafCatId(leafCatId int64) error {
    r.leafCatId = leafCatId
    r.Set("leaf_cat_id", leafCatId)
    return nil
}

func (r AlibabaGpuAddSchemaGetRequest) GetLeafCatId() int64 {
    return r.leafCatId
}

func (r *AlibabaGpuAddSchemaGetRequest) SetBrandId(brandId int64) error {
    r.brandId = brandId
    r.Set("brand_id", brandId)
    return nil
}

func (r AlibabaGpuAddSchemaGetRequest) GetBrandId() int64 {
    return r.brandId
}

func (r *AlibabaGpuAddSchemaGetRequest) SetProviderId(providerId int64) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

func (r AlibabaGpuAddSchemaGetRequest) GetProviderId() int64 {
    return r.providerId
}

