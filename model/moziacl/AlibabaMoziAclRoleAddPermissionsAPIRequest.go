package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziAclRoleAddPermissionsAPIRequest 角色添加功能权限 API请求
// alibaba.mozi.acl.role.add.permissions
//
// 往角色中添加一批功能权限
type AlibabaMoziAclRoleAddPermissionsAPIRequest struct {
	model.Params
	// 角色添加功能权限请求对象
	_addPermissionsToRole *AddPermissionToRoleRequest
}

// NewAlibabaMoziAclRoleAddPermissionsRequest 初始化AlibabaMoziAclRoleAddPermissionsAPIRequest对象
func NewAlibabaMoziAclRoleAddPermissionsRequest() *AlibabaMoziAclRoleAddPermissionsAPIRequest {
	return &AlibabaMoziAclRoleAddPermissionsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclRoleAddPermissionsAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.role.add.permissions"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclRoleAddPermissionsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziAclRoleAddPermissionsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAddPermissionsToRole is AddPermissionsToRole Setter
// 角色添加功能权限请求对象
func (r *AlibabaMoziAclRoleAddPermissionsAPIRequest) SetAddPermissionsToRole(_addPermissionsToRole *AddPermissionToRoleRequest) error {
	r._addPermissionsToRole = _addPermissionsToRole
	r.Set("add_permissions_to_role", _addPermissionsToRole)
	return nil
}

// GetAddPermissionsToRole AddPermissionsToRole Getter
func (r AlibabaMoziAclRoleAddPermissionsAPIRequest) GetAddPermissionsToRole() *AddPermissionToRoleRequest {
	return r._addPermissionsToRole
}
