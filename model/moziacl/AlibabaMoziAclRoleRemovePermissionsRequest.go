package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
角色移除功能权限 APIRequest
alibaba.mozi.acl.role.remove.permissions

从角色中移除一批功能权限
*/
type AlibabaMoziAclRoleRemovePermissionsRequest struct {
    model.Params

    // 角色移除功能权限请求对象
    removePermissionsFromRole   *RemovePermissionsFromRoleRequest 

}

func NewAlibabaMoziAclRoleRemovePermissionsRequest() *AlibabaMoziAclRoleRemovePermissionsRequest{
    return &AlibabaMoziAclRoleRemovePermissionsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziAclRoleRemovePermissionsRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.role.remove.permissions"
}

func (r AlibabaMoziAclRoleRemovePermissionsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziAclRoleRemovePermissionsRequest) SetRemovePermissionsFromRole(removePermissionsFromRole *RemovePermissionsFromRoleRequest) error {
    r.removePermissionsFromRole = removePermissionsFromRole
    r.Set("remove_permissions_from_role", removePermissionsFromRole)
    return nil
}

func (r AlibabaMoziAclRoleRemovePermissionsRequest) GetRemovePermissionsFromRole() *RemovePermissionsFromRoleRequest {
    return r.removePermissionsFromRole
}

