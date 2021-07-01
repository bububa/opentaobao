package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziAclPermissionpkgAddPermissionsAPIRequest
权限套餐添加权限 API请求
alibaba.mozi.acl.permissionpkg.add.permissions

此接口的功能为：将一批应用下的权限添加到该应用下的权限套餐中 */
type AlibabaMoziAclPermissionpkgAddPermissionsAPIRequest struct {
	model.Params
	// 请求对象
	_parameters *UpdatePermissionsToPermissionPackageRequest
}

// New
