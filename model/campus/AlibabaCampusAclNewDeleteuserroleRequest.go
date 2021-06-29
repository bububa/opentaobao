package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除管理员 API请求
alibaba.campus.acl.new.deleteuserrole

删除管理员
*/
type AlibabaCampusAclNewDeleteuserroleRequest struct {
    model.Params
    // 系统入参
    workbenchcontext   *WorkBenchContext
    // 用户账号
    userId   string
    // 角色id
    roleIds   []int64
}

// 初始化AlibabaCampusAclNewDeleteuserroleRequest对象
func NewAlibabaCampusAclNewDeleteuserroleRequest() *AlibabaCampusAclNewDeleteuserroleRequest{
    return &AlibabaCampusAclNewDeleteuserroleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewDeleteuserroleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.deleteuserrole"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewDeleteuserroleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewDeleteuserroleRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

// Workbenchcontext Getter
func (r AlibabaCampusAclNewDeleteuserroleRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}
// UserId Setter
// 用户账号
func (r *AlibabaCampusAclNewDeleteuserroleRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaCampusAclNewDeleteuserroleRequest) GetUserId() string {
    return r.userId
}
// RoleIds Setter
// 角色id
func (r *AlibabaCampusAclNewDeleteuserroleRequest) SetRoleIds(roleIds []int64) error {
    r.roleIds = roleIds
    r.Set("role_ids", roleIds)
    return nil
}

// RoleIds Getter
func (r AlibabaCampusAclNewDeleteuserroleRequest) GetRoleIds() []int64 {
    return r.roleIds
}
