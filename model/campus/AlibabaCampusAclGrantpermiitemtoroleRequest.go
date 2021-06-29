package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
权限赋予角色 APIRequest
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

func NewAlibabaCampusAclGrantpermiitemtoroleRequest() *AlibabaCampusAclGrantpermiitemtoroleRequest{
    return &AlibabaCampusAclGrantpermiitemtoroleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.grantpermiitemtorole"
}

func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclGrantpermiitemtoroleRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetCompanyId() int64 {
    return r.companyId
}

func (r *AlibabaCampusAclGrantpermiitemtoroleRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetSystemId() string {
    return r.systemId
}

func (r *AlibabaCampusAclGrantpermiitemtoroleRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetCampusId() int64 {
    return r.campusId
}

func (r *AlibabaCampusAclGrantpermiitemtoroleRequest) SetRole(role *RoleReq) error {
    r.role = role
    r.Set("role", role)
    return nil
}

func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetRole() *RoleReq {
    return r.role
}

func (r *AlibabaCampusAclGrantpermiitemtoroleRequest) SetPriv(priv []PermissionReq) error {
    r.priv = priv
    r.Set("priv", priv)
    return nil
}

func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetPriv() []PermissionReq {
    return r.priv
}

func (r *AlibabaCampusAclGrantpermiitemtoroleRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaCampusAclGrantpermiitemtoroleRequest) GetUserId() string {
    return r.userId
}

