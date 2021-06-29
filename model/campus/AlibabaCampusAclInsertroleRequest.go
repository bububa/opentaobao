package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增角色 API请求
alibaba.campus.acl.insertrole

新增角色
*/
type AlibabaCampusAclInsertroleRequest struct {
    model.Params
    // 公司id,不填统一为SYS_000
    companyId   int64
    // 系统id
    systemId   string
    // 园区id
    campusId   int64
    // 角色描述
    roleDesc   string
    // 角色名称
    roleName   string
    // 默认人员角色.还有device设备角色类型
    roleType   string
    // 角色唯一ID,统一ROLE_开头,不填默认生成ROLE_UUID(32位随机数)
    roleId   string
    // 操作人id(不填默认appCode)
    userId   string
}

// 初始化AlibabaCampusAclInsertroleRequest对象
func NewAlibabaCampusAclInsertroleRequest() *AlibabaCampusAclInsertroleRequest{
    return &AlibabaCampusAclInsertroleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclInsertroleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.insertrole"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclInsertroleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CompanyId Setter
// 公司id,不填统一为SYS_000
func (r *AlibabaCampusAclInsertroleRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclInsertroleRequest) GetCompanyId() int64 {
    return r.companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclInsertroleRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclInsertroleRequest) GetSystemId() string {
    return r.systemId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclInsertroleRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclInsertroleRequest) GetCampusId() int64 {
    return r.campusId
}
// RoleDesc Setter
// 角色描述
func (r *AlibabaCampusAclInsertroleRequest) SetRoleDesc(roleDesc string) error {
    r.roleDesc = roleDesc
    r.Set("role_desc", roleDesc)
    return nil
}

// RoleDesc Getter
func (r AlibabaCampusAclInsertroleRequest) GetRoleDesc() string {
    return r.roleDesc
}
// RoleName Setter
// 角色名称
func (r *AlibabaCampusAclInsertroleRequest) SetRoleName(roleName string) error {
    r.roleName = roleName
    r.Set("role_name", roleName)
    return nil
}

// RoleName Getter
func (r AlibabaCampusAclInsertroleRequest) GetRoleName() string {
    return r.roleName
}
// RoleType Setter
// 默认人员角色.还有device设备角色类型
func (r *AlibabaCampusAclInsertroleRequest) SetRoleType(roleType string) error {
    r.roleType = roleType
    r.Set("role_type", roleType)
    return nil
}

// RoleType Getter
func (r AlibabaCampusAclInsertroleRequest) GetRoleType() string {
    return r.roleType
}
// RoleId Setter
// 角色唯一ID,统一ROLE_开头,不填默认生成ROLE_UUID(32位随机数)
func (r *AlibabaCampusAclInsertroleRequest) SetRoleId(roleId string) error {
    r.roleId = roleId
    r.Set("role_id", roleId)
    return nil
}

// RoleId Getter
func (r AlibabaCampusAclInsertroleRequest) GetRoleId() string {
    return r.roleId
}
// UserId Setter
// 操作人id(不填默认appCode)
func (r *AlibabaCampusAclInsertroleRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaCampusAclInsertroleRequest) GetUserId() string {
    return r.userId
}
