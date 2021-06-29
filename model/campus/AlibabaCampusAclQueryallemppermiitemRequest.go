package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询员工全部权限(包括角色下面的权限) API请求
alibaba.campus.acl.queryallemppermiitem

查询员工全部权限(包括角色下面的权限)
*/
type AlibabaCampusAclQueryallemppermiitemRequest struct {
    model.Params
    // 公司id不填默认SYS_000
    companyId   int64
    // 系统id
    systemId   string
    // 园区id
    campusId   int64
    // 用户账号
    accountId   string
    // 每页多少条
    page   int64
    // 每页记录数
    pageSize   int64
}

// 初始化AlibabaCampusAclQueryallemppermiitemRequest对象
func NewAlibabaCampusAclQueryallemppermiitemRequest() *AlibabaCampusAclQueryallemppermiitemRequest{
    return &AlibabaCampusAclQueryallemppermiitemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclQueryallemppermiitemRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.queryallemppermiitem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclQueryallemppermiitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CompanyId Setter
// 公司id不填默认SYS_000
func (r *AlibabaCampusAclQueryallemppermiitemRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclQueryallemppermiitemRequest) GetCompanyId() int64 {
    return r.companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclQueryallemppermiitemRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclQueryallemppermiitemRequest) GetSystemId() string {
    return r.systemId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclQueryallemppermiitemRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclQueryallemppermiitemRequest) GetCampusId() int64 {
    return r.campusId
}
// AccountId Setter
// 用户账号
func (r *AlibabaCampusAclQueryallemppermiitemRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

// AccountId Getter
func (r AlibabaCampusAclQueryallemppermiitemRequest) GetAccountId() string {
    return r.accountId
}
// Page Setter
// 每页多少条
func (r *AlibabaCampusAclQueryallemppermiitemRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r AlibabaCampusAclQueryallemppermiitemRequest) GetPage() int64 {
    return r.page
}
// PageSize Setter
// 每页记录数
func (r *AlibabaCampusAclQueryallemppermiitemRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaCampusAclQueryallemppermiitemRequest) GetPageSize() int64 {
    return r.pageSize
}
