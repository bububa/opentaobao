package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
角色添加功能权限 APIRequest
alibaba.mozi.acl.role.add.permissions

往角色中添加一批功能权限
*/
type AlibabaMoziAclRoleAddPermissionsRequest struct {
    model.Params

    // 角色添加功能权限请求对象
    addPermissionsToRole   *AddPermissionToRoleRequest 

}

func NewAlibabaMoziAclRoleAddPermissionsRequest() *AlibabaMoziAclRoleAddPermissionsRequest{
    return &AlibabaMoziAclRoleAddPermissionsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziAclRoleAddPermissionsRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.role.add.permissions"
}

func (r AlibabaMoziAclRoleAddPermissionsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziAclRoleAddPermissionsRequest) SetAddPermissionsToRole(addPermissionsToRole *AddPermissionToRoleRequest) error {
    r.addPermissionsToRole = addPermissionsToRole
    r.Set("add_permissions_to_role", addPermissionsToRole)
    return nil
}

func (r AlibabaMoziAclRoleAddPermissionsRequest) GetAddPermissionsToRole() *AddPermissionToRoleRequest {
    return r.addPermissionsToRole
}

