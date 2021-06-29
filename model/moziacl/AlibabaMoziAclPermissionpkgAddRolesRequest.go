package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
将角色添加到权限套餐中 APIRequest
alibaba.mozi.acl.permissionpkg.add.roles

此接口是将应用下的一批角色添加到该应用的某个权限套餐中
*/
type AlibabaMoziAclPermissionpkgAddRolesRequest struct {
    model.Params

    // 请求入参对象
    parameters   *UpdateRolesToPermissionPackageRequest 

}

func NewAlibabaMoziAclPermissionpkgAddRolesRequest() *AlibabaMoziAclPermissionpkgAddRolesRequest{
    return &AlibabaMoziAclPermissionpkgAddRolesRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziAclPermissionpkgAddRolesRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.permissionpkg.add.roles"
}

func (r AlibabaMoziAclPermissionpkgAddRolesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziAclPermissionpkgAddRolesRequest) SetParameters(parameters *UpdateRolesToPermissionPackageRequest) error {
    r.parameters = parameters
    r.Set("parameters", parameters)
    return nil
}

func (r AlibabaMoziAclPermissionpkgAddRolesRequest) GetParameters() *UpdateRolesToPermissionPackageRequest {
    return r.parameters
}

