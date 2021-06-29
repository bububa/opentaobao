package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
权限套餐添加权限 APIRequest
alibaba.mozi.acl.permissionpkg.add.permissions

此接口的功能为：将一批应用下的权限添加到该应用下的权限套餐中
*/
type AlibabaMoziAclPermissionpkgAddPermissionsRequest struct {
    model.Params

    // 请求对象
    parameters   *UpdatePermissionsToPermissionPackageRequest 

}

func NewAlibabaMoziAclPermissionpkgAddPermissionsRequest() *AlibabaMoziAclPermissionpkgAddPermissionsRequest{
    return &AlibabaMoziAclPermissionpkgAddPermissionsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziAclPermissionpkgAddPermissionsRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.permissionpkg.add.permissions"
}

func (r AlibabaMoziAclPermissionpkgAddPermissionsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziAclPermissionpkgAddPermissionsRequest) SetParameters(parameters *UpdatePermissionsToPermissionPackageRequest) error {
    r.parameters = parameters
    r.Set("parameters", parameters)
    return nil
}

func (r AlibabaMoziAclPermissionpkgAddPermissionsRequest) GetParameters() *UpdatePermissionsToPermissionPackageRequest {
    return r.parameters
}

