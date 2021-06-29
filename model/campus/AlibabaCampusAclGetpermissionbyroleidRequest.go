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
    _systemId   string
    // 园区id
    _campusId   int64
    // 角色id
    _roleId   string
    // 公司id
    _companyId   int64
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
func (r *AlibabaCampusAclGetpermissionbyroleidRequest) SetSystemId(_systemId string) error {
    r._systemId = _systemId
    r.Set("system_id", _systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclGetpermissionbyroleidRequest) GetSystemId() string {
    return r._systemId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclGetpermissionbyroleidRequest) SetCampusId(_campusId int64) error {
    r._campusId = _campusId
    r.Set("campus_id", _campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclGetpermissionbyroleidRequest) GetCampusId() int64 {
    return r._campusId
}
// RoleId Setter
// 角色id
func (r *AlibabaCampusAclGetpermissionbyroleidRequest) SetRoleId(_roleId string) error {
    r._roleId = _roleId
    r.Set("role_id", _roleId)
    return nil
}

// RoleId Getter
func (r AlibabaCampusAclGetpermissionbyroleidRequest) GetRoleId() string {
    return r._roleId
}
// CompanyId Setter
// 公司id
func (r *AlibabaCampusAclGetpermissionbyroleidRequest) SetCompanyId(_companyId int64) error {
    r._companyId = _companyId
    r.Set("company_id", _companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclGetpermissionbyroleidRequest) GetCompanyId() int64 {
    return r._companyId
}
