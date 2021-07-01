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

// New
