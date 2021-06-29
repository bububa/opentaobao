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
    _companyId   int64
    // 系统id
    _systemId   string
    // 园区id
    _campusId   int64
    // 用户账号
    _accountId   string
    // 每页多少条
    _page   int64
    // 每页记录数
    _pageSize   int64
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
func (r *AlibabaCampusAclQueryallemppermiitemRequest) SetCompanyId(_companyId int64) error {
    r._companyId = _companyId
    r.Set("company_id", _companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclQueryallemppermiitemRequest) GetCompanyId() int64 {
    return r._companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclQueryallemppermiitemRequest) SetSystemId(_systemId string) error {
    r._systemId = _systemId
    r.Set("system_id", _systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclQueryallemppermiitemRequest) GetSystemId() string {
    return r._systemId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclQueryallemppermiitemRequest) SetCampusId(_campusId int64) error {
    r._campusId = _campusId
    r.Set("campus_id", _campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclQueryallemppermiitemRequest) GetCampusId() int64 {
    return r._campusId
}
// AccountId Setter
// 用户账号
func (r *AlibabaCampusAclQueryallemppermiitemRequest) SetAccountId(_accountId string) error {
    r._accountId = _accountId
    r.Set("account_id", _accountId)
    return nil
}

// AccountId Getter
func (r AlibabaCampusAclQueryallemppermiitemRequest) GetAccountId() string {
    return r._accountId
}
// Page Setter
// 每页多少条
func (r *AlibabaCampusAclQueryallemppermiitemRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaCampusAclQueryallemppermiitemRequest) GetPage() int64 {
    return r._page
}
// PageSize Setter
// 每页记录数
func (r *AlibabaCampusAclQueryallemppermiitemRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaCampusAclQueryallemppermiitemRequest) GetPageSize() int64 {
    return r._pageSize
}
