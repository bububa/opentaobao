package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询全部角色 APIRequest
alibaba.campus.acl.queryallrole

查询全部园区
*/
type AlibabaCampusAclQueryallroleRequest struct {
    model.Params

    // 公司id不填统一SYS_000
    companyId   int64 

    // 系统id
    systemId   string 

    // 园区id
    campusId   int64 

    // 角色名称
    roleName   string 

    // 角色类型
    roleType   string 

    // 角色id
    roleId   string 

}

func NewAlibabaCampusAclQueryallroleRequest() *AlibabaCampusAclQueryallroleRequest{
    return &AlibabaCampusAclQueryallroleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclQueryallroleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.queryallrole"
}

func (r AlibabaCampusAclQueryallroleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclQueryallroleRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

func (r AlibabaCampusAclQueryallroleRequest) GetCompanyId() int64 {
    return r.companyId
}

func (r *AlibabaCampusAclQueryallroleRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

func (r AlibabaCampusAclQueryallroleRequest) GetSystemId() string {
    return r.systemId
}

func (r *AlibabaCampusAclQueryallroleRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaCampusAclQueryallroleRequest) GetCampusId() int64 {
    return r.campusId
}

func (r *AlibabaCampusAclQueryallroleRequest) SetRoleName(roleName string) error {
    r.roleName = roleName
    r.Set("role_name", roleName)
    return nil
}

func (r AlibabaCampusAclQueryallroleRequest) GetRoleName() string {
    return r.roleName
}

func (r *AlibabaCampusAclQueryallroleRequest) SetRoleType(roleType string) error {
    r.roleType = roleType
    r.Set("role_type", roleType)
    return nil
}

func (r AlibabaCampusAclQueryallroleRequest) GetRoleType() string {
    return r.roleType
}

func (r *AlibabaCampusAclQueryallroleRequest) SetRoleId(roleId string) error {
    r.roleId = roleId
    r.Set("role_id", roleId)
    return nil
}

func (r AlibabaCampusAclQueryallroleRequest) GetRoleId() string {
    return r.roleId
}

