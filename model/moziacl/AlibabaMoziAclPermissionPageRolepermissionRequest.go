package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询角色下包含的权限列表 APIRequest
alibaba.mozi.acl.permission.page.rolepermission

根据传入的角色name，分页查询该角色包含的权限列表
*/
type AlibabaMoziAclPermissionPageRolepermissionRequest struct {
    model.Params

    // 分页查询角色下包含的权限列表
    pageRolePermisions   *PageRolePermissionRequest 

}

func NewAlibabaMoziAclPermissionPageRolepermissionRequest() *AlibabaMoziAclPermissionPageRolepermissionRequest{
    return &AlibabaMoziAclPermissionPageRolepermissionRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziAclPermissionPageRolepermissionRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.permission.page.rolepermission"
}

func (r AlibabaMoziAclPermissionPageRolepermissionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziAclPermissionPageRolepermissionRequest) SetPageRolePermisions(pageRolePermisions *PageRolePermissionRequest) error {
    r.pageRolePermisions = pageRolePermisions
    r.Set("page_role_permisions", pageRolePermisions)
    return nil
}

func (r AlibabaMoziAclPermissionPageRolepermissionRequest) GetPageRolePermisions() *PageRolePermissionRequest {
    return r.pageRolePermisions
}

