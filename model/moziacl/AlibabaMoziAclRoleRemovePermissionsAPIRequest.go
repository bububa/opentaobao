package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziAclRoleRemovePermissionsAPIRequest 角色移除功能权限 API请求
// alibaba.mozi.acl.role.remove.permissions
//
// 从角色中移除一批功能权限
type AlibabaMoziAclRoleRemovePermissionsAPIRequest struct {
	model.Params
	// 角色移除功能权限请求对象
	_removePermissionsFromRole *RemovePermissionsFromRoleRequest
}

// NewAlibabaMoziAclRoleRemovePermissionsRequest 初始化AlibabaMoziAclRoleRemovePermissionsAPIRequest对象
func NewAlibabaMoziAclRoleRemovePermissionsRequest() *AlibabaMoziAclRoleRemovePermissionsAPIRequest {
	return &AlibabaMoziAclRoleRemovePermissionsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclRoleRemovePermissionsAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.role.remove.permissions"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclRoleRemovePermissionsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziAclRoleRemovePermissionsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRemovePermissionsFromRole is RemovePermissionsFromRole Setter
// 角色移除功能权限请求对象
func (r *AlibabaMoziAclRoleRemovePermissionsAPIRequest) SetRemovePermissionsFromRole(_removePermissionsFromRole *RemovePermissionsFromRoleRequest) error {
	r._removePermissionsFromRole = _removePermissionsFromRole
	r.Set("remove_permissions_from_role", _removePermissionsFromRole)
	return nil
}

// GetRemovePermissionsFromRole RemovePermissionsFromRole Getter
func (r AlibabaMoziAclRoleRemovePermissionsAPIRequest) GetRemovePermissionsFromRole() *RemovePermissionsFromRoleRequest {
	return r._removePermissionsFromRole
}
