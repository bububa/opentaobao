package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改用户到角色关系 API请求
alibaba.campus.acl.updategrantroletouser

修改用户到角色关系
*/
type AlibabaCampusAclUpdategrantroletouserRequest struct {
    model.Params
    // 公司id
    companyId   int64
    // 系统id
    systemId   string
    // 园区id
    campusId   int64
    // 用户账号
    accountId   string
    // 角色
    role   []RoleReq
    // 操作人id(不填默认appCode)
    userId   string
}

// 初始化AlibabaCampusAclUpdategrantroletouserRequest对象
func NewAlibabaCampusAclUpdategrantroletouserRequest() *AlibabaCampusAclUpdategrantroletouserRequest{
    return &AlibabaCampusAclUpdategrantroletouserRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclUpdategrantroletouserRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.updategrantroletouser"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclUpdategrantroletouserRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CompanyId Setter
// 公司id
func (r *AlibabaCampusAclUpdategrantroletouserRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclUpdategrantroletouserRequest) GetCompanyId() int64 {
    return r.companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclUpdategrantroletouserRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclUpdategrantroletouserRequest) GetSystemId() string {
    return r.systemId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclUpdategrantroletouserRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclUpdategrantroletouserRequest) GetCampusId() int64 {
    return r.campusId
}
// AccountId Setter
// 用户账号
func (r *AlibabaCampusAclUpdategrantroletouserRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

// AccountId Getter
func (r AlibabaCampusAclUpdategrantroletouserRequest) GetAccountId() string {
    return r.accountId
}
// Role Setter
// 角色
func (r *AlibabaCampusAclUpdategrantroletouserRequest) SetRole(role []RoleReq) error {
    r.role = role
    r.Set("role", role)
    return nil
}

// Role Getter
func (r AlibabaCampusAclUpdategrantroletouserRequest) GetRole() []RoleReq {
    return r.role
}
// UserId Setter
// 操作人id(不填默认appCode)
func (r *AlibabaCampusAclUpdategrantroletouserRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaCampusAclUpdategrantroletouserRequest) GetUserId() string {
    return r.userId
}
