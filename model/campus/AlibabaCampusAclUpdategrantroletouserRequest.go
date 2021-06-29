package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改用户到角色关系 APIRequest
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

func NewAlibabaCampusAclUpdategrantroletouserRequest() *AlibabaCampusAclUpdategrantroletouserRequest{
    return &AlibabaCampusAclUpdategrantroletouserRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclUpdategrantroletouserRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.updategrantroletouser"
}

func (r AlibabaCampusAclUpdategrantroletouserRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclUpdategrantroletouserRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

func (r AlibabaCampusAclUpdategrantroletouserRequest) GetCompanyId() int64 {
    return r.companyId
}

func (r *AlibabaCampusAclUpdategrantroletouserRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

func (r AlibabaCampusAclUpdategrantroletouserRequest) GetSystemId() string {
    return r.systemId
}

func (r *AlibabaCampusAclUpdategrantroletouserRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaCampusAclUpdategrantroletouserRequest) GetCampusId() int64 {
    return r.campusId
}

func (r *AlibabaCampusAclUpdategrantroletouserRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

func (r AlibabaCampusAclUpdategrantroletouserRequest) GetAccountId() string {
    return r.accountId
}

func (r *AlibabaCampusAclUpdategrantroletouserRequest) SetRole(role []RoleReq) error {
    r.role = role
    r.Set("role", role)
    return nil
}

func (r AlibabaCampusAclUpdategrantroletouserRequest) GetRole() []RoleReq {
    return r.role
}

func (r *AlibabaCampusAclUpdategrantroletouserRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaCampusAclUpdategrantroletouserRequest) GetUserId() string {
    return r.userId
}

