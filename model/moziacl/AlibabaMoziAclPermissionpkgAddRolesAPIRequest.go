package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziAclPermissionpkgAddRolesAPIRequest
将角色添加到权限套餐中 API请求
alibaba.mozi.acl.permissionpkg.add.roles

此接口是将应用下的一批角色添加到该应用的某个权限套餐中 */
type AlibabaMoziAclPermissionpkgAddRolesAPIRequest struct {
	model.Params
	// 请求入参对象
	_parameters *UpdateRolesToPermissionPackageRequest
}

// NewAlibabaMoziAclPermissionpkgAddRolesRequest 初始化AlibabaMoziAclPermissionpkgAddRolesAPIRequest对象
func NewAlibabaMoziAclPermissionpkgAddRolesRequest() *AlibabaMoziAclPermissionpkgAddRolesAPIRequest {
	return &AlibabaMoziAclPermissionpkgAddRolesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclPermissionpkgAddRolesAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.permissionpkg.add.roles"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclPermissionpkgAddRolesAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Parameters Setter
// 请求入参对象
func (r *AlibabaMoziAclPermissionpkgAddRolesAPIRequest) SetParameters(_parameters *UpdateRolesToPermissionPackageRequest) error {
	r._parameters = _parameters
	r.Set("parameters", _parameters)
	return nil
}

// Get Parameters Getter
func (r AlibabaMoziAclPermissionpkgAddRolesAPIRequest) GetParameters() *UpdateRolesToPermissionPackageRequest {
	return r._parameters
}
