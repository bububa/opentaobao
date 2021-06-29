package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
权限赋予角色 API请求
alibaba.campus.acl.grantpermiitemtorole

权限赋予角色
*/
type AlibabaCampusAclGrantpermiitemtoroleRequest struct {
    model.Params
    // 公司ID,不填统一初始化SYS_000
    companyId   int64
    // 系统id
    systemId   string
    // 园区ID
    campusId   int64
    // 系统自动生成
    role   *RoleReq
    // 系统自动生成
    priv   []PermissionReq
    // 操作人id
    userId   string
}

// 初始化AlibabaCampusAclGrantpermiitemtoroleRequest对象
func NewAlibabaCampusAclGrantpermiitemtoroleRequest() *AlibabaCampusAclGrantpermiitemtoroleRequest{
    return &AlibabaCampusAclGrantpermiitemtoroleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.grantpermiitemtorole"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CompanyId Setter
// 公司ID,不填统一初始化SYS_000
func (r *AlibabaCampusAclGrantpermiitemtoroleRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetCompanyId() int64 {
    return r.companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclGrantpermiitemtoroleRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetSystemId() string {
    return r.systemId
}
// CampusId Setter
// 园区ID
func (r *AlibabaCampusAclGrantpermiitemtoroleRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetCampusId() int64 {
    return r.campusId
}
// Role Setter
// 系统自动生成
func (r *AlibabaCampusAclGrantpermiitemtoroleRequest) SetRole(role *RoleReq) error {
    r.role = role
    r.Set("role", role)
    return nil
}

// Role Getter
func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetRole() *RoleReq {
    return r.role
}
// Priv Setter
// 系统自动生成
func (r *AlibabaCampusAclGrantpermiitemtoroleRequest) SetPriv(priv []PermissionReq) error {
    r.priv = priv
    r.Set("priv", priv)
    return nil
}

// Priv Getter
func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetPriv() []PermissionReq {
    return r.priv
}
// UserId Setter
// 操作人id
func (r *AlibabaCampusAclGrantpermiitemtoroleRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetUserId() string {
    return r.userId
}
