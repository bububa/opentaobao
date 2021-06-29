package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
权限套餐添加权限 API请求
alibaba.mozi.acl.permissionpkg.add.permissions

此接口的功能为：将一批应用下的权限添加到该应用下的权限套餐中
*/
type AlibabaMoziAclPermissionpkgAddPermissionsRequest struct {
    model.Params
    // 请求对象
    _parameters   *UpdatePermissionsToPermissionPackageRequest
}

// 初始化AlibabaMoziAclPermissionpkgAddPermissionsRequest对象
func NewAlibabaMoziAclPermissionpkgAddPermissionsRequest() *AlibabaMoziAclPermissionpkgAddPermissionsRequest{
    return &AlibabaMoziAclPermissionpkgAddPermissionsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclPermissionpkgAddPermissionsRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.permissionpkg.add.permissions"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclPermissionpkgAddPermissionsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Parameters Setter
// 请求对象
func (r *AlibabaMoziAclPermissionpkgAddPermissionsRequest) SetParameters(_parameters *UpdatePermissionsToPermissionPackageRequest) error {
    r._parameters = _parameters
    r.Set("parameters", _parameters)
    return nil
}

// Parameters Getter
func (r AlibabaMoziAclPermissionpkgAddPermissionsRequest) GetParameters() *UpdatePermissionsToPermissionPackageRequest {
    return r._parameters
}
