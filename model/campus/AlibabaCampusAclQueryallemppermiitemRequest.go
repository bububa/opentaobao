package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询员工全部权限(包括角色下面的权限) APIRequest
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

func NewAlibabaCampusAclQueryallemppermiitemRequest() *AlibabaCampusAclQueryallemppermiitemRequest{
    return &AlibabaCampusAclQueryallemppermiitemRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclQueryallemppermiitemRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.queryallemppermiitem"
}

func (r AlibabaCampusAclQueryallemppermiitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclQueryallemppermiitemRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

func (r AlibabaCampusAclQueryallemppermiitemRequest) GetCompanyId() int64 {
    return r.companyId
}

func (r *AlibabaCampusAclQueryallemppermiitemRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

func (r AlibabaCampusAclQueryallemppermiitemRequest) GetSystemId() string {
    return r.systemId
}

func (r *AlibabaCampusAclQueryallemppermiitemRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaCampusAclQueryallemppermiitemRequest) GetCampusId() int64 {
    return r.campusId
}

func (r *AlibabaCampusAclQueryallemppermiitemRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

func (r AlibabaCampusAclQueryallemppermiitemRequest) GetAccountId() string {
    return r.accountId
}

func (r *AlibabaCampusAclQueryallemppermiitemRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaCampusAclQueryallemppermiitemRequest) GetPage() int64 {
    return r.page
}

func (r *AlibabaCampusAclQueryallemppermiitemRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaCampusAclQueryallemppermiitemRequest) GetPageSize() int64 {
    return r.pageSize
}

