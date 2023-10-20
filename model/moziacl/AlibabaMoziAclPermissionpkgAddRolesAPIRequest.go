package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamoziaclpermissionpkgaddrolesAPIRequest 将角色添加到权限套餐中 API请求
// alibaba.mozi.acl.permissionpkg.add.roles
//
// 此接口是将应用下的一批角色添加到该应用的某个权限套餐中
type AlibabamoziaclpermissionpkgaddrolesAPIRequest struct {
	model.Params
	// 请求入参对象
	_parameters *UpdateRolesToPermissionPackageRequest
}

// NewAlibabamoziaclpermissionpkgaddrolesRequest 初始化AlibabamoziaclpermissionpkgaddrolesAPIRequest对象
func NewAlibabamoziaclpermissionpkgaddrolesRequest() *AlibabamoziaclpermissionpkgaddrolesAPIRequest {
	return &AlibabamoziaclpermissionpkgaddrolesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamoziaclpermissionpkgaddrolesAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.permissionpkg.add.roles"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamoziaclpermissionpkgaddrolesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamoziaclpermissionpkgaddrolesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParameters is Parameters Setter
// 请求入参对象
func (r *AlibabamoziaclpermissionpkgaddrolesAPIRequest) SetParameters(_parameters *UpdateRolesToPermissionPackageRequest) error {
	r._parameters = _parameters
	r.Set("parameters", _parameters)
	return nil
}

// GetParameters Parameters Getter
func (r AlibabamoziaclpermissionpkgaddrolesAPIRequest) GetParameters() *UpdateRolesToPermissionPackageRequest {
	return r._parameters
}
