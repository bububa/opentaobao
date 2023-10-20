package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamoziaclroleaddpermissionsAPIRequest 角色添加功能权限 API请求
// alibaba.mozi.acl.role.add.permissions
//
// 往角色中添加一批功能权限
type AlibabamoziaclroleaddpermissionsAPIRequest struct {
	model.Params
	// 角色添加功能权限请求对象
	_addPermissionsToRole *AddPermissionToRoleRequest
}

// NewAlibabamoziaclroleaddpermissionsRequest 初始化AlibabamoziaclroleaddpermissionsAPIRequest对象
func NewAlibabamoziaclroleaddpermissionsRequest() *AlibabamoziaclroleaddpermissionsAPIRequest {
	return &AlibabamoziaclroleaddpermissionsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamoziaclroleaddpermissionsAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.role.add.permissions"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamoziaclroleaddpermissionsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamoziaclroleaddpermissionsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAddPermissionsToRole is AddPermissionsToRole Setter
// 角色添加功能权限请求对象
func (r *AlibabamoziaclroleaddpermissionsAPIRequest) SetAddPermissionsToRole(_addPermissionsToRole *AddPermissionToRoleRequest) error {
	r._addPermissionsToRole = _addPermissionsToRole
	r.Set("add_permissions_to_role", _addPermissionsToRole)
	return nil
}

// GetAddPermissionsToRole AddPermissionsToRole Getter
func (r AlibabamoziaclroleaddpermissionsAPIRequest) GetAddPermissionsToRole() *AddPermissionToRoleRequest {
	return r._addPermissionsToRole
}
