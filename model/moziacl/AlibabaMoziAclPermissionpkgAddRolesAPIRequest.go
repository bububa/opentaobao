package moziacl

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziAclPermissionpkgAddRolesAPIRequest 将角色添加到权限套餐中 API请求
// alibaba.mozi.acl.permissionpkg.add.roles
//
// 此接口是将应用下的一批角色添加到该应用的某个权限套餐中
type AlibabaMoziAclPermissionpkgAddRolesAPIRequest struct {
	model.Params
	// 请求入参对象
	_parameters *UpdateRolesToPermissionPackageRequest
}

// NewAlibabaMoziAclPermissionpkgAddRolesRequest 初始化AlibabaMoziAclPermissionpkgAddRolesAPIRequest对象
func NewAlibabaMoziAclPermissionpkgAddRolesRequest() *AlibabaMoziAclPermissionpkgAddRolesAPIRequest {
	return &AlibabaMoziAclPermissionpkgAddRolesAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMoziAclPermissionpkgAddRolesAPIRequest) Reset() {
	r._parameters = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclPermissionpkgAddRolesAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.permissionpkg.add.roles"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclPermissionpkgAddRolesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziAclPermissionpkgAddRolesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParameters is Parameters Setter
// 请求入参对象
func (r *AlibabaMoziAclPermissionpkgAddRolesAPIRequest) SetParameters(_parameters *UpdateRolesToPermissionPackageRequest) error {
	r._parameters = _parameters
	r.Set("parameters", _parameters)
	return nil
}

// GetParameters Parameters Getter
func (r AlibabaMoziAclPermissionpkgAddRolesAPIRequest) GetParameters() *UpdateRolesToPermissionPackageRequest {
	return r._parameters
}

var poolAlibabaMoziAclPermissionpkgAddRolesAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMoziAclPermissionpkgAddRolesRequest()
	},
}

// GetAlibabaMoziAclPermissionpkgAddRolesRequest 从 sync.Pool 获取 AlibabaMoziAclPermissionpkgAddRolesAPIRequest
func GetAlibabaMoziAclPermissionpkgAddRolesAPIRequest() *AlibabaMoziAclPermissionpkgAddRolesAPIRequest {
	return poolAlibabaMoziAclPermissionpkgAddRolesAPIRequest.Get().(*AlibabaMoziAclPermissionpkgAddRolesAPIRequest)
}

// ReleaseAlibabaMoziAclPermissionpkgAddRolesAPIRequest 将 AlibabaMoziAclPermissionpkgAddRolesAPIRequest 放入 sync.Pool
func ReleaseAlibabaMoziAclPermissionpkgAddRolesAPIRequest(v *AlibabaMoziAclPermissionpkgAddRolesAPIRequest) {
	v.Reset()
	poolAlibabaMoziAclPermissionpkgAddRolesAPIRequest.Put(v)
}
