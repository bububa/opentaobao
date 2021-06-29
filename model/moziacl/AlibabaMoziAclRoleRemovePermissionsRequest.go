package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
角色移除功能权限 API请求
alibaba.mozi.acl.role.remove.permissions

从角色中移除一批功能权限
*/
type AlibabaMoziAclRoleRemovePermissionsRequest struct {
    model.Params
    // 角色移除功能权限请求对象
    removePermissionsFromRole   *RemovePermissionsFromRoleRequest
}

// 初始化AlibabaMoziAclRoleRemovePermissionsRequest对象
func NewAlibabaMoziAclRoleRemovePermissionsRequest() *AlibabaMoziAclRoleRemovePermissionsRequest{
    return &AlibabaMoziAclRoleRemovePermissionsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclRoleRemovePermissionsRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.role.remove.permissions"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclRoleRemovePermissionsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RemovePermissionsFromRole Setter
// 角色移除功能权限请求对象
func (r *AlibabaMoziAclRoleRemovePermissionsRequest) SetRemovePermissionsFromRole(removePermissionsFromRole *RemovePermissionsFromRoleRequest) error {
    r.removePermissionsFromRole = removePermissionsFromRole
    r.Set("remove_permissions_from_role", removePermissionsFromRole)
    return nil
}

// RemovePermissionsFromRole Getter
func (r AlibabaMoziAclRoleRemovePermissionsRequest) GetRemovePermissionsFromRole() *RemovePermissionsFromRoleRequest {
    return r.removePermissionsFromRole
}
