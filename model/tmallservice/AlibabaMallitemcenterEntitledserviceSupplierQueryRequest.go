package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据天猫id查询门店服务授权 API请求
alibaba.mallitemcenter.entitledservice.supplier.query

根据天猫id查询门店服务授权
*/
type AlibabaMallitemcenterEntitledserviceSupplierQueryRequest struct {
    model.Params
    // 天猫id
    id   int64
    // 第几页
    currentPage   int64
    // 每页条数
    pageSize   int64
}

// 初始化AlibabaMallitemcenterEntitledserviceSupplierQueryRequest对象
func NewAlibabaMallitemcenterEntitledserviceSupplierQueryRequest() *AlibabaMallitemcenterEntitledserviceSupplierQueryRequest{
    return &AlibabaMallitemcenterEntitledserviceSupplierQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMallitemcenterEntitledserviceSupplierQueryRequest) GetApiMethodName() string {
    return "alibaba.mallitemcenter.entitledservice.supplier.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMallitemcenterEntitledserviceSupplierQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 天猫id
func (r *AlibabaMallitemcenterEntitledserviceSupplierQueryRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r AlibabaMallitemcenterEntitledserviceSupplierQueryRequest) GetId() int64 {
    return r.id
}
// CurrentPage Setter
// 第几页
func (r *AlibabaMallitemcenterEntitledserviceSupplierQueryRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaMallitemcenterEntitledserviceSupplierQueryRequest) GetCurrentPage() int64 {
    return r.currentPage
}
// PageSize Setter
// 每页条数
func (r *AlibabaMallitemcenterEntitledserviceSupplierQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaMallitemcenterEntitledserviceSupplierQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
