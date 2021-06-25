package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按类目查询spu接口 APIRequest
alibaba.gpu.schema.catsearch

按类目查询spu的schema接口
*/
type AlibabaGpuSchemaCatsearchRequest struct {
    model.Params

    // 叶子类目ID
    leafCatId   int64 

    // 当前页
    currentPage   int64 

    // 每页大小
    pageSize   int64 

    // 渠道Id，如0代表天猫，8代表淘宝
    providerId   int64 

}

func NewAlibabaGpuSchemaCatsearchRequest() *AlibabaGpuSchemaCatsearchRequest{
    return &AlibabaGpuSchemaCatsearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaGpuSchemaCatsearchRequest) GetApiMethodName() string {
    return "alibaba.gpu.schema.catsearch"
}

func (r AlibabaGpuSchemaCatsearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaGpuSchemaCatsearchRequest) SetLeafCatId(leafCatId int64) error {
    r.leafCatId = leafCatId
    r.Set("leaf_cat_id", leafCatId)
    return nil
}

func (r AlibabaGpuSchemaCatsearchRequest) GetLeafCatId() int64 {
    return r.leafCatId
}

func (r *AlibabaGpuSchemaCatsearchRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r AlibabaGpuSchemaCatsearchRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *AlibabaGpuSchemaCatsearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaGpuSchemaCatsearchRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaGpuSchemaCatsearchRequest) SetProviderId(providerId int64) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

func (r AlibabaGpuSchemaCatsearchRequest) GetProviderId() int64 {
    return r.providerId
}

