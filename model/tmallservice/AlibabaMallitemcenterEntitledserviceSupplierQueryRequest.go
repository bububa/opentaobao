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
type AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest struct {
    model.Params
    // 天猫id
    _id   int64
    // 第几页
    _currentPage   int64
    // 每页条数
    _pageSize   int64
}

// 初始化AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest对象
func NewAlibabaMallitemcenterEntitledserviceSupplierQueryRequest() *AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest{
    return &AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.mallitemcenter.entitledservice.supplier.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 天猫id
func (r *AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest) GetId() int64 {
    return r._id
}
// CurrentPage Setter
// 第几页
func (r *AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// PageSize Setter
// 每页条数
func (r *AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
