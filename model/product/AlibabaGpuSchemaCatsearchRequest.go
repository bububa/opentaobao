package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按类目查询spu接口 API请求
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

// 初始化AlibabaGpuSchemaCatsearchRequest对象
func NewAlibabaGpuSchemaCatsearchRequest() *AlibabaGpuSchemaCatsearchRequest{
    return &AlibabaGpuSchemaCatsearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaGpuSchemaCatsearchRequest) GetApiMethodName() string {
    return "alibaba.gpu.schema.catsearch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaGpuSchemaCatsearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LeafCatId Setter
// 叶子类目ID
func (r *AlibabaGpuSchemaCatsearchRequest) SetLeafCatId(leafCatId int64) error {
    r.leafCatId = leafCatId
    r.Set("leaf_cat_id", leafCatId)
    return nil
}

// LeafCatId Getter
func (r AlibabaGpuSchemaCatsearchRequest) GetLeafCatId() int64 {
    return r.leafCatId
}
// CurrentPage Setter
// 当前页
func (r *AlibabaGpuSchemaCatsearchRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaGpuSchemaCatsearchRequest) GetCurrentPage() int64 {
    return r.currentPage
}
// PageSize Setter
// 每页大小
func (r *AlibabaGpuSchemaCatsearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaGpuSchemaCatsearchRequest) GetPageSize() int64 {
    return r.pageSize
}
// ProviderId Setter
// 渠道Id，如0代表天猫，8代表淘宝
func (r *AlibabaGpuSchemaCatsearchRequest) SetProviderId(providerId int64) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

// ProviderId Getter
func (r AlibabaGpuSchemaCatsearchRequest) GetProviderId() int64 {
    return r.providerId
}
