package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询角色下包含的权限列表 API请求
alibaba.mozi.acl.permission.page.rolepermission

根据传入的角色name，分页查询该角色包含的权限列表
*/
type AlibabaMoziAclPermissionPageRolepermissionAPIRequest struct {
    model.Params
    // 分页查询角色下包含的权限列表
    _pageRolePermisions   *PageRolePermissionRequest
}

// 初始化AlibabaMoziAclPermissionPageRolepermissionAPIRequest对象
func NewAlibabaMoziAclPermissionPageRolepermissionRequest() *AlibabaMoziAclPermissionPageRolepermissionAPIRequest{
    return &AlibabaMoziAclPermissionPageRolepermissionAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclPermissionPageRolepermissionAPIRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.permission.page.rolepermission"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclPermissionPageRolepermissionAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageRolePermisions Setter
// 分页查询角色下包含的权限列表
func (r *AlibabaMoziAclPermissionPageRolepermissionAPIRequest) SetPageRolePermisions(_pageRolePermisions *PageRolePermissionRequest) error {
    r._pageRolePermisions = _pageRolePermisions
    r.Set("page_role_permisions", _pageRolePermisions)
    return nil
}

// PageRolePermisions Getter
func (r AlibabaMoziAclPermissionPageRolepermissionAPIRequest) GetPageRolePermisions() *PageRolePermissionRequest {
    return r._pageRolePermisions
}
