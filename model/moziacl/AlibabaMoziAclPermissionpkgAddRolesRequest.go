package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
将角色添加到权限套餐中 API请求
alibaba.mozi.acl.permissionpkg.add.roles

此接口是将应用下的一批角色添加到该应用的某个权限套餐中
*/
type AlibabaMoziAclPermissionpkgAddRolesRequest struct {
    model.Params
    // 请求入参对象
    _parameters   *UpdateRolesToPermissionPackageRequest
}

// 初始化AlibabaMoziAclPermissionpkgAddRolesRequest对象
func NewAlibabaMoziAclPermissionpkgAddRolesRequest() *AlibabaMoziAclPermissionpkgAddRolesRequest{
    return &AlibabaMoziAclPermissionpkgAddRolesRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclPermissionpkgAddRolesRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.permissionpkg.add.roles"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclPermissionpkgAddRolesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Parameters Setter
// 请求入参对象
func (r *AlibabaMoziAclPermissionpkgAddRolesRequest) SetParameters(_parameters *UpdateRolesToPermissionPackageRequest) error {
    r._parameters = _parameters
    r.Set("parameters", _parameters)
    return nil
}

// Parameters Getter
func (r AlibabaMoziAclPermissionpkgAddRolesRequest) GetParameters() *UpdateRolesToPermissionPackageRequest {
    return r._parameters
}
