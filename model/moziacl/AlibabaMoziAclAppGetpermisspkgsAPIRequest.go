package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamoziaclappgetpermisspkgsAPIRequest 分页获取应用的权限套餐 API请求
// alibaba.mozi.acl.app.getpermisspkgs
//
// 分页查询应用下的权限套餐列表
type AlibabamoziaclappgetpermisspkgsAPIRequest struct {
	model.Params
	// 获取应用的权限套餐请求对象
	_getAppPermissionPackagesRequest *GetAppPermissionPackageRequest
}

// NewAlibabamoziaclappgetpermisspkgsRequest 初始化AlibabamoziaclappgetpermisspkgsAPIRequest对象
func NewAlibabamoziaclappgetpermisspkgsRequest() *AlibabamoziaclappgetpermisspkgsAPIRequest {
	return &AlibabamoziaclappgetpermisspkgsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamoziaclappgetpermisspkgsAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.app.getpermisspkgs"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamoziaclappgetpermisspkgsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamoziaclappgetpermisspkgsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGetAppPermissionPackagesRequest is GetAppPermissionPackagesRequest Setter
// 获取应用的权限套餐请求对象
func (r *AlibabamoziaclappgetpermisspkgsAPIRequest) SetGetAppPermissionPackagesRequest(_getAppPermissionPackagesRequest *GetAppPermissionPackageRequest) error {
	r._getAppPermissionPackagesRequest = _getAppPermissionPackagesRequest
	r.Set("get_app_permission_packages_request", _getAppPermissionPackagesRequest)
	return nil
}

// GetGetAppPermissionPackagesRequest GetAppPermissionPackagesRequest Getter
func (r AlibabamoziaclappgetpermisspkgsAPIRequest) GetGetAppPermissionPackagesRequest() *GetAppPermissionPackageRequest {
	return r._getAppPermissionPackagesRequest
}
