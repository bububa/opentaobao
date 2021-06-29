package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否有角色 API请求
alibaba.campus.acl.new.checkuserrole

校验用户是否有角色
*/
type AlibabaCampusAclNewCheckuserroleRequest struct {
    model.Params
    // 用户账号
    userId   string
    // 角色key
    roleKey   string
    // 系统入参
    workbenchcontext   *WorkBenchContext
}

// 初始化AlibabaCampusAclNewCheckuserroleRequest对象
func NewAlibabaCampusAclNewCheckuserroleRequest() *AlibabaCampusAclNewCheckuserroleRequest{
    return &AlibabaCampusAclNewCheckuserroleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewCheckuserroleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.checkuserrole"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewCheckuserroleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户账号
func (r *AlibabaCampusAclNewCheckuserroleRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaCampusAclNewCheckuserroleRequest) GetUserId() string {
    return r.userId
}
// RoleKey Setter
// 角色key
func (r *AlibabaCampusAclNewCheckuserroleRequest) SetRoleKey(roleKey string) error {
    r.roleKey = roleKey
    r.Set("role_key", roleKey)
    return nil
}

// RoleKey Getter
func (r AlibabaCampusAclNewCheckuserroleRequest) GetRoleKey() string {
    return r.roleKey
}
// Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewCheckuserroleRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

// Workbenchcontext Getter
func (r AlibabaCampusAclNewCheckuserroleRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}
