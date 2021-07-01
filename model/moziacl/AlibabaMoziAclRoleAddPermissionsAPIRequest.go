package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziAclRoleAddPermissionsAPIRequest
角色添加功能权限 API请求
alibaba.mozi.acl.role.add.permissions

往角色中添加一批功能权限 */
type AlibabaMoziAclRoleAddPermissionsAPIRequest struct {
	model.Params
	// 角色添加功能权限请求对象
	_addPermissionsToRole *AddPermissionToRoleRequest
}

// New
