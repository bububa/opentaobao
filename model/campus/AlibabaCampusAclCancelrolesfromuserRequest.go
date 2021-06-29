package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
撤销用户授予的角色 APIRequest
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

func NewAlibabaCampusAclCancelrolesfromuserRequest() *AlibabaCampusAclCancelrolesfromuserRequest{
    return &AlibabaCampusAclCancelrolesfromuserRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclCancelrolesfromuserRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.cancelrolesfromuser"
}

func (r AlibabaCampusAclCancelrolesfromuserRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclCancelrolesfromuserRequest) SetRole(role []RoleReq) error {
    r.role = role
    r.Set("role", role)
    return nil
}

func (r AlibabaCampusAclCancelrolesfromuserRequest) GetRole() []RoleReq {
    return r.role
}

func (r *AlibabaCampusAclCancelrolesfromuserRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

func (r AlibabaCampusAclCancelrolesfromuserRequest) GetCompanyId() int64 {
    return r.companyId
}

func (r *AlibabaCampusAclCancelrolesfromuserRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

func (r AlibabaCampusAclCancelrolesfromuserRequest) GetSystemId() string {
    return r.systemId
}

func (r *AlibabaCampusAclCancelrolesfromuserRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaCampusAclCancelrolesfromuserRequest) GetCampusId() int64 {
    return r.campusId
}

func (r *AlibabaCampusAclCancelrolesfromuserRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

func (r AlibabaCampusAclCancelrolesfromuserRequest) GetAccountId() string {
    return r.accountId
}

func (r *AlibabaCampusAclCancelrolesfromuserRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaCampusAclCancelrolesfromuserRequest) GetUserId() string {
    return r.userId
}

