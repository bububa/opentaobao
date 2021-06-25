package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据天猫id查询门店服务授权 APIRequest
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

func NewAlibabaMallitemcenterEntitledserviceSupplierQueryRequest() *AlibabaMallitemcenterEntitledserviceSupplierQueryRequest{
    return &AlibabaMallitemcenterEntitledserviceSupplierQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMallitemcenterEntitledserviceSupplierQueryRequest) GetApiMethodName() string {
    return "alibaba.mallitemcenter.entitledservice.supplier.query"
}

func (r AlibabaMallitemcenterEntitledserviceSupplierQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMallitemcenterEntitledserviceSupplierQueryRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaMallitemcenterEntitledserviceSupplierQueryRequest) GetId() int64 {
    return r.id
}

func (r *AlibabaMallitemcenterEntitledserviceSupplierQueryRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r AlibabaMallitemcenterEntitledserviceSupplierQueryRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *AlibabaMallitemcenterEntitledserviceSupplierQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaMallitemcenterEntitledserviceSupplierQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

