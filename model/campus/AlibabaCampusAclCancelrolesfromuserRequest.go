package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
撤销用户授予的角色 API请求
alibaba.campus.acl.cancelrolesfromuser

撤销用户授予的角色
*/
type AlibabaCampusAclCancelrolesfromuserRequest struct {
    model.Params
    // 系统自动生成
    role   []RoleReq
    // 公司id
    companyId   int64
    // 系统id
    systemId   string
    // 园区id
    campusId   int64
    // 用户账号
    accountId   string
    // 操作人id(不填默认appCode)
    userId   string
}

// 初始化AlibabaCampusAclCancelrolesfromuserRequest对象
func NewAlibabaCampusAclCancelrolesfromuserRequest() *AlibabaCampusAclCancelrolesfromuserRequest{
    return &AlibabaCampusAclCancelrolesfromuserRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclCancelrolesfromuserRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.cancelrolesfromuser"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclCancelrolesfromuserRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Role Setter
// 系统自动生成
func (r *AlibabaCampusAclCancelrolesfromuserRequest) SetRole(role []RoleReq) error {
    r.role = role
    r.Set("role", role)
    return nil
}

// Role Getter
func (r AlibabaCampusAclCancelrolesfromuserRequest) GetRole() []RoleReq {
    return r.role
}
// CompanyId Setter
// 公司id
func (r *AlibabaCampusAclCancelrolesfromuserRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclCancelrolesfromuserRequest) GetCompanyId() int64 {
    return r.companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclCancelrolesfromuserRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclCancelrolesfromuserRequest) GetSystemId() string {
    return r.systemId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclCancelrolesfromuserRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclCancelrolesfromuserRequest) GetCampusId() int64 {
    return r.campusId
}
// AccountId Setter
// 用户账号
func (r *AlibabaCampusAclCancelrolesfromuserRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

// AccountId Getter
func (r AlibabaCampusAclCancelrolesfromuserRequest) GetAccountId() string {
    return r.accountId
}
// UserId Setter
// 操作人id(不填默认appCode)
func (r *AlibabaCampusAclCancelrolesfromuserRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaCampusAclCancelrolesfromuserRequest) GetUserId() string {
    return r.userId
}
