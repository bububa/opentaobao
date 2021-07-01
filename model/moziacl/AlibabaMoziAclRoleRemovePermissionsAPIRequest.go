package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziAclRoleRemovePermissionsAPIRequest
角色移除功能权限 API请求
alibaba.mozi.acl.role.remove.permissions

从角色中移除一批功能权限 */
type AlibabaMoziAclRoleRemovePermissionsAPIRequest struct {
	model.Params
	// 角色移除功能权限请求对象
	_removePermissionsFromRole *RemovePermissionsFromRoleRequest
}

// New
