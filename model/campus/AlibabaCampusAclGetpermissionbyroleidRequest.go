package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据角色Id查询权限 APIRequest
alibaba.campus.acl.getpermissionbyroleid

根据角色查询权限
*/
type AlibabaCampusAclGetpermissionbyroleidRequest struct {
    model.Params

    // 系统id
    systemId   string 

    // 园区id
    campusId   int64 

    // 角色id
    roleId   string 

    // 公司id
    companyId   int64 

}

func NewAlibabaCampusAclGetpermissionbyroleidRequest() *AlibabaCampusAclGetpermissionbyroleidRequest{
    return &AlibabaCampusAclGetpermissionbyroleidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclGetpermissionbyroleidRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.getpermissionbyroleid"
}

func (r AlibabaCampusAclGetpermissionbyroleidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclGetpermissionbyroleidRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

func (r AlibabaCampusAclGetpermissionbyroleidRequest) GetSystemId() string {
    return r.systemId
}

func (r *AlibabaCampusAclGetpermissionbyroleidRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaCampusAclGetpermissionbyroleidRequest) GetCampusId() int64 {
    return r.campusId
}

func (r *AlibabaCampusAclGetpermissionbyroleidRequest) SetRoleId(roleId string) error {
    r.roleId = roleId
    r.Set("role_id", roleId)
    return nil
}

func (r AlibabaCampusAclGetpermissionbyroleidRequest) GetRoleId() string {
    return r.roleId
}

func (r *AlibabaCampusAclGetpermissionbyroleidRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

func (r AlibabaCampusAclGetpermissionbyroleidRequest) GetCompanyId() int64 {
    return r.companyId
}

