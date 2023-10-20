package moziacl

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMoziAclRoleRemovePermissionsAPIRequest) Reset() {
	r._removePermissionsFromRole = nil
	r.Params.ToZero()
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

var poolAlibabaMoziAclRoleRemovePermissionsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMoziAclRoleRemovePermissionsRequest()
	},
}

// GetAlibabaMoziAclRoleRemovePermissionsRequest 从 sync.Pool 获取 AlibabaMoziAclRoleRemovePermissionsAPIRequest
func GetAlibabaMoziAclRoleRemovePermissionsAPIRequest() *AlibabaMoziAclRoleRemovePermissionsAPIRequest {
	return poolAlibabaMoziAclRoleRemovePermissionsAPIRequest.Get().(*AlibabaMoziAclRoleRemovePermissionsAPIRequest)
}

// ReleaseAlibabaMoziAclRoleRemovePermissionsAPIRequest 将 AlibabaMoziAclRoleRemovePermissionsAPIRequest 放入 sync.Pool
func ReleaseAlibabaMoziAclRoleRemovePermissionsAPIRequest(v *AlibabaMoziAclRoleRemovePermissionsAPIRequest) {
	v.Reset()
	poolAlibabaMoziAclRoleRemovePermissionsAPIRequest.Put(v)
}
