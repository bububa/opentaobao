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
type AlibabaCampusAclInsertroleAPIRequest struct {
    model.Params
    // 公司id,不填统一为SYS_000
    _companyId   int64
    // 系统id
    _systemId   string
    // 园区id
    _campusId   int64
    // 角色描述
    _roleDesc   string
    // 角色名称
    _roleName   string
    // 默认人员角色.还有device设备角色类型
    _roleType   string
    // 角色唯一ID,统一ROLE_开头,不填默认生成ROLE_UUID(32位随机数)
    _roleId   string
    // 操作人id(不填默认appCode)
    _userId   string
}

// 初始化AlibabaCampusAclInsertroleAPIRequest对象
func NewAlibabaCampusAclInsertroleRequest() *AlibabaCampusAclInsertroleAPIRequest{
    return &AlibabaCampusAclInsertroleAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclInsertroleAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.insertrole"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclInsertroleAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CompanyId Setter
// 公司id,不填统一为SYS_000
func (r *AlibabaCampusAclInsertroleAPIRequest) SetCompanyId(_companyId int64) error {
    r._companyId = _companyId
    r.Set("company_id", _companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclInsertroleAPIRequest) GetCompanyId() int64 {
    return r._companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclInsertroleAPIRequest) SetSystemId(_systemId string) error {
    r._systemId = _systemId
    r.Set("system_id", _systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclInsertroleAPIRequest) GetSystemId() string {
    return r._systemId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclInsertroleAPIRequest) SetCampusId(_campusId int64) error {
    r._campusId = _campusId
    r.Set("campus_id", _campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclInsertroleAPIRequest) GetCampusId() int64 {
    return r._campusId
}
// RoleDesc Setter
// 角色描述
func (r *AlibabaCampusAclInsertroleAPIRequest) SetRoleDesc(_roleDesc string) error {
    r._roleDesc = _roleDesc
    r.Set("role_desc", _roleDesc)
    return nil
}

// RoleDesc Getter
func (r AlibabaCampusAclInsertroleAPIRequest) GetRoleDesc() string {
    return r._roleDesc
}
// RoleName Setter
// 角色名称
func (r *AlibabaCampusAclInsertroleAPIRequest) SetRoleName(_roleName string) error {
    r._roleName = _roleName
    r.Set("role_name", _roleName)
    return nil
}

// RoleName Getter
func (r AlibabaCampusAclInsertroleAPIRequest) GetRoleName() string {
    return r._roleName
}
// RoleType Setter
// 默认人员角色.还有device设备角色类型
func (r *AlibabaCampusAclInsertroleAPIRequest) SetRoleType(_roleType string) error {
    r._roleType = _roleType
    r.Set("role_type", _roleType)
    return nil
}

// RoleType Getter
func (r AlibabaCampusAclInsertroleAPIRequest) GetRoleType() string {
    return r._roleType
}
// RoleId Setter
// 角色唯一ID,统一ROLE_开头,不填默认生成ROLE_UUID(32位随机数)
func (r *AlibabaCampusAclInsertroleAPIRequest) SetRoleId(_roleId string) error {
    r._roleId = _roleId
    r.Set("role_id", _roleId)
    return nil
}

// RoleId Getter
func (r AlibabaCampusAclInsertroleAPIRequest) GetRoleId() string {
    return r._roleId
}
// UserId Setter
// 操作人id(不填默认appCode)
func (r *AlibabaCampusAclInsertroleAPIRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaCampusAclInsertroleAPIRequest) GetUserId() string {
    return r._userId
}
