package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据角色Id查询权限 API请求
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

// 初始化AlibabaCampusAclGetpermissionbyroleidRequest对象
func NewAlibabaCampusAclGetpermissionbyroleidRequest() *AlibabaCampusAclGetpermissionbyroleidRequest{
    return &AlibabaCampusAclGetpermissionbyroleidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclGetpermissionbyroleidRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.getpermissionbyroleid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclGetpermissionbyroleidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclGetpermissionbyroleidRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclGetpermissionbyroleidRequest) GetSystemId() string {
    return r.systemId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclGetpermissionbyroleidRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclGetpermissionbyroleidRequest) GetCampusId() int64 {
    return r.campusId
}
// RoleId Setter
// 角色id
func (r *AlibabaCampusAclGetpermissionbyroleidRequest) SetRoleId(roleId string) error {
    r.roleId = roleId
    r.Set("role_id", roleId)
    return nil
}

// RoleId Getter
func (r AlibabaCampusAclGetpermissionbyroleidRequest) GetRoleId() string {
    return r.roleId
}
// CompanyId Setter
// 公司id
func (r *AlibabaCampusAclGetpermissionbyroleidRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclGetpermissionbyroleidRequest) GetCompanyId() int64 {
    return r.companyId
}
