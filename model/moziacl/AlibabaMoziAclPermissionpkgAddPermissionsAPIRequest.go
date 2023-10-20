package moziacl

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziAclPermissionpkgAddPermissionsAPIRequest 权限套餐添加权限 API请求
// alibaba.mozi.acl.permissionpkg.add.permissions
//
// 此接口的功能为：将一批应用下的权限添加到该应用下的权限套餐中
type AlibabaMoziAclPermissionpkgAddPermissionsAPIRequest struct {
	model.Params
	// 请求对象
	_parameters *UpdatePermissionsToPermissionPackageRequest
}

// NewAlibabaMoziAclPermissionpkgAddPermissionsRequest 初始化AlibabaMoziAclPermissionpkgAddPermissionsAPIRequest对象
func NewAlibabaMoziAclPermissionpkgAddPermissionsRequest() *AlibabaMoziAclPermissionpkgAddPermissionsAPIRequest {
	return &AlibabaMoziAclPermissionpkgAddPermissionsAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMoziAclPermissionpkgAddPermissionsAPIRequest) Reset() {
	r._parameters = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclPermissionpkgAddPermissionsAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.permissionpkg.add.permissions"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclPermissionpkgAddPermissionsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziAclPermissionpkgAddPermissionsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParameters is Parameters Setter
// 请求对象
func (r *AlibabaMoziAclPermissionpkgAddPermissionsAPIRequest) SetParameters(_parameters *UpdatePermissionsToPermissionPackageRequest) error {
	r._parameters = _parameters
	r.Set("parameters", _parameters)
	return nil
}

// GetParameters Parameters Getter
func (r AlibabaMoziAclPermissionpkgAddPermissionsAPIRequest) GetParameters() *UpdatePermissionsToPermissionPackageRequest {
	return r._parameters
}

var poolAlibabaMoziAclPermissionpkgAddPermissionsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMoziAclPermissionpkgAddPermissionsRequest()
	},
}

// GetAlibabaMoziAclPermissionpkgAddPermissionsRequest 从 sync.Pool 获取 AlibabaMoziAclPermissionpkgAddPermissionsAPIRequest
func GetAlibabaMoziAclPermissionpkgAddPermissionsAPIRequest() *AlibabaMoziAclPermissionpkgAddPermissionsAPIRequest {
	return poolAlibabaMoziAclPermissionpkgAddPermissionsAPIRequest.Get().(*AlibabaMoziAclPermissionpkgAddPermissionsAPIRequest)
}

// ReleaseAlibabaMoziAclPermissionpkgAddPermissionsAPIRequest 将 AlibabaMoziAclPermissionpkgAddPermissionsAPIRequest 放入 sync.Pool
func ReleaseAlibabaMoziAclPermissionpkgAddPermissionsAPIRequest(v *AlibabaMoziAclPermissionpkgAddPermissionsAPIRequest) {
	v.Reset()
	poolAlibabaMoziAclPermissionpkgAddPermissionsAPIRequest.Put(v)
}
