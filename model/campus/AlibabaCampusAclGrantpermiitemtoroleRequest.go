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
    _companyId   int64
    // 系统id
    _systemId   string
    // 园区ID
    _campusId   int64
    // 系统自动生成
    _role   *RoleReq
    // 系统自动生成
    _priv   []PermissionReq
    // 操作人id
    _userId   string
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
func (r *AlibabaCampusAclGrantpermiitemtoroleRequest) SetCompanyId(_companyId int64) error {
    r._companyId = _companyId
    r.Set("company_id", _companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetCompanyId() int64 {
    return r._companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclGrantpermiitemtoroleRequest) SetSystemId(_systemId string) error {
    r._systemId = _systemId
    r.Set("system_id", _systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetSystemId() string {
    return r._systemId
}
// CampusId Setter
// 园区ID
func (r *AlibabaCampusAclGrantpermiitemtoroleRequest) SetCampusId(_campusId int64) error {
    r._campusId = _campusId
    r.Set("campus_id", _campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetCampusId() int64 {
    return r._campusId
}
// Role Setter
// 系统自动生成
func (r *AlibabaCampusAclGrantpermiitemtoroleRequest) SetRole(_role *RoleReq) error {
    r._role = _role
    r.Set("role", _role)
    return nil
}

// Role Getter
func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetRole() *RoleReq {
    return r._role
}
// Priv Setter
// 系统自动生成
func (r *AlibabaCampusAclGrantpermiitemtoroleRequest) SetPriv(_priv []PermissionReq) error {
    r._priv = _priv
    r.Set("priv", _priv)
    return nil
}

// Priv Getter
func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetPriv() []PermissionReq {
    return r._priv
}
// UserId Setter
// 操作人id
func (r *AlibabaCampusAclGrantpermiitemtoroleRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetUserId() string {
    return r._userId
}