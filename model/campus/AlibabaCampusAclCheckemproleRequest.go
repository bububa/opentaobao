package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否有该角色 API请求
alibaba.campus.acl.checkemprole

校验用户是否有该权限
*/
type AlibabaCampusAclCheckemproleRequest struct {
    model.Params
    // 公司id不填默认为SYS_000
    companyId   int64
    // 系统id
    systemId   string
    // 园区id
    campusId   int64
    // 员工账号
    accountId   string
    // 角色id
    itemKey   string
}

// 初始化AlibabaCampusAclCheckemproleRequest对象
func NewAlibabaCampusAclCheckemproleRequest() *AlibabaCampusAclCheckemproleRequest{
    return &AlibabaCampusAclCheckemproleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclCheckemproleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.checkemprole"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclCheckemproleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CompanyId Setter
// 公司id不填默认为SYS_000
func (r *AlibabaCampusAclCheckemproleRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclCheckemproleRequest) GetCompanyId() int64 {
    return r.companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclCheckemproleRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclCheckemproleRequest) GetSystemId() string {
    return r.systemId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclCheckemproleRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclCheckemproleRequest) GetCampusId() int64 {
    return r.campusId
}
// AccountId Setter
// 员工账号
func (r *AlibabaCampusAclCheckemproleRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

// AccountId Getter
func (r AlibabaCampusAclCheckemproleRequest) GetAccountId() string {
    return r.accountId
}
// ItemKey Setter
// 角色id
func (r *AlibabaCampusAclCheckemproleRequest) SetItemKey(itemKey string) error {
    r.itemKey = itemKey
    r.Set("item_key", itemKey)
    return nil
}

// ItemKey Getter
func (r AlibabaCampusAclCheckemproleRequest) GetItemKey() string {
    return r.itemKey
}
