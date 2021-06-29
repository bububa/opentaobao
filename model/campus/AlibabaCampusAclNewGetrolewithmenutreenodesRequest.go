package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据角色id查询权限 APIRequest
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

func NewAlibabaCampusAclNewGetrolewithmenutreenodesRequest() *AlibabaCampusAclNewGetrolewithmenutreenodesRequest{
    return &AlibabaCampusAclNewGetrolewithmenutreenodesRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclNewGetrolewithmenutreenodesRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.getrolewithmenutreenodes"
}

func (r AlibabaCampusAclNewGetrolewithmenutreenodesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclNewGetrolewithmenutreenodesRequest) SetRoleId(roleId int64) error {
    r.roleId = roleId
    r.Set("role_id", roleId)
    return nil
}

func (r AlibabaCampusAclNewGetrolewithmenutreenodesRequest) GetRoleId() int64 {
    return r.roleId
}

func (r *AlibabaCampusAclNewGetrolewithmenutreenodesRequest) SetAllPermission(allPermission bool) error {
    r.allPermission = allPermission
    r.Set("all_permission", allPermission)
    return nil
}

func (r AlibabaCampusAclNewGetrolewithmenutreenodesRequest) GetAllPermission() bool {
    return r.allPermission
}

func (r *AlibabaCampusAclNewGetrolewithmenutreenodesRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

func (r AlibabaCampusAclNewGetrolewithmenutreenodesRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}

