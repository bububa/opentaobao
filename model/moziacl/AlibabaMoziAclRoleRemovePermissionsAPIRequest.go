package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamoziaclroleremovepermissionsAPIRequest 角色移除功能权限 API请求
// alibaba.mozi.acl.role.remove.permissions
//
// 从角色中移除一批功能权限
type AlibabamoziaclroleremovepermissionsAPIRequest struct {
	model.Params
	// 角色移除功能权限请求对象
	_removePermissionsFromRole *RemovePermissionsFromRoleRequest
}

// NewAlibabamoziaclroleremovepermissionsRequest 初始化AlibabamoziaclroleremovepermissionsAPIRequest对象
func NewAlibabamoziaclroleremovepermissionsRequest() *AlibabamoziaclroleremovepermissionsAPIRequest {
	return &AlibabamoziaclroleremovepermissionsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamoziaclroleremovepermissionsAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.role.remove.permissions"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamoziaclroleremovepermissionsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamoziaclroleremovepermissionsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRemovePermissionsFromRole is RemovePermissionsFromRole Setter
// 角色移除功能权限请求对象
func (r *AlibabamoziaclroleremovepermissionsAPIRequest) SetRemovePermissionsFromRole(_removePermissionsFromRole *RemovePermissionsFromRoleRequest) error {
	r._removePermissionsFromRole = _removePermissionsFromRole
	r.Set("remove_permissions_from_role", _removePermissionsFromRole)
	return nil
}

// GetRemovePermissionsFromRole RemovePermissionsFromRole Getter
func (r AlibabamoziaclroleremovepermissionsAPIRequest) GetRemovePermissionsFromRole() *RemovePermissionsFromRoleRequest {
	return r._removePermissionsFromRole
}
