package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamoziaclpermissionpkgaddpermissionsAPIRequest 权限套餐添加权限 API请求
// alibaba.mozi.acl.permissionpkg.add.permissions
//
// 此接口的功能为：将一批应用下的权限添加到该应用下的权限套餐中
type AlibabamoziaclpermissionpkgaddpermissionsAPIRequest struct {
	model.Params
	// 请求对象
	_parameters *UpdatePermissionsToPermissionPackageRequest
}

// NewAlibabamoziaclpermissionpkgaddpermissionsRequest 初始化AlibabamoziaclpermissionpkgaddpermissionsAPIRequest对象
func NewAlibabamoziaclpermissionpkgaddpermissionsRequest() *AlibabamoziaclpermissionpkgaddpermissionsAPIRequest {
	return &AlibabamoziaclpermissionpkgaddpermissionsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamoziaclpermissionpkgaddpermissionsAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.permissionpkg.add.permissions"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamoziaclpermissionpkgaddpermissionsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamoziaclpermissionpkgaddpermissionsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParameters is Parameters Setter
// 请求对象
func (r *AlibabamoziaclpermissionpkgaddpermissionsAPIRequest) SetParameters(_parameters *UpdatePermissionsToPermissionPackageRequest) error {
	r._parameters = _parameters
	r.Set("parameters", _parameters)
	return nil
}

// GetParameters Parameters Getter
func (r AlibabamoziaclpermissionpkgaddpermissionsAPIRequest) GetParameters() *UpdatePermissionsToPermissionPackageRequest {
	return r._parameters
}
