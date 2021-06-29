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
    _userId   string
    // 角色key
    _roleKey   string
    // 系统入参
    _workbenchcontext   *WorkBenchContext
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
func (r *AlibabaCampusAclNewCheckuserroleRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaCampusAclNewCheckuserroleRequest) GetUserId() string {
    return r._userId
}
// RoleKey Setter
// 角色key
func (r *AlibabaCampusAclNewCheckuserroleRequest) SetRoleKey(_roleKey string) error {
    r._roleKey = _roleKey
    r.Set("role_key", _roleKey)
    return nil
}

// RoleKey Getter
func (r AlibabaCampusAclNewCheckuserroleRequest) GetRoleKey() string {
    return r._roleKey
}
// Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewCheckuserroleRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
    r._workbenchcontext = _workbenchcontext
    r.Set("workbenchcontext", _workbenchcontext)
    return nil
}

// Workbenchcontext Getter
func (r AlibabaCampusAclNewCheckuserroleRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r._workbenchcontext
}
