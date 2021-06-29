package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据角色id查询权限 API请求
alibaba.campus.acl.new.getrolewithmenutreenodes

根据角色id查询权限
*/
type AlibabaCampusAclNewGetrolewithmenutreenodesRequest struct {
    model.Params
    // 角色id
    roleId   int64
    // 是否查询全部类型权限
    allPermission   bool
    // 系统参数
    workbenchcontext   *WorkBenchContext
}

// 初始化AlibabaCampusAclNewGetrolewithmenutreenodesRequest对象
func NewAlibabaCampusAclNewGetrolewithmenutreenodesRequest() *AlibabaCampusAclNewGetrolewithmenutreenodesRequest{
    return &AlibabaCampusAclNewGetrolewithmenutreenodesRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewGetrolewithmenutreenodesRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.getrolewithmenutreenodes"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewGetrolewithmenutreenodesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RoleId Setter
// 角色id
func (r *AlibabaCampusAclNewGetrolewithmenutreenodesRequest) SetRoleId(roleId int64) error {
    r.roleId = roleId
    r.Set("role_id", roleId)
    return nil
}

// RoleId Getter
func (r AlibabaCampusAclNewGetrolewithmenutreenodesRequest) GetRoleId() int64 {
    return r.roleId
}
// AllPermission Setter
// 是否查询全部类型权限
func (r *AlibabaCampusAclNewGetrolewithmenutreenodesRequest) SetAllPermission(allPermission bool) error {
    r.allPermission = allPermission
    r.Set("all_permission", allPermission)
    return nil
}

// AllPermission Getter
func (r AlibabaCampusAclNewGetrolewithmenutreenodesRequest) GetAllPermission() bool {
    return r.allPermission
}
// Workbenchcontext Setter
// 系统参数
func (r *AlibabaCampusAclNewGetrolewithmenutreenodesRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

// Workbenchcontext Getter
func (r AlibabaCampusAclNewGetrolewithmenutreenodesRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}
