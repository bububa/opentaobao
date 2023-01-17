package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziAclAppGetpermisspkgsAPIRequest 分页获取应用的权限套餐 API请求
// alibaba.mozi.acl.app.getpermisspkgs
//
// 分页查询应用下的权限套餐列表
type AlibabaMoziAclAppGetpermisspkgsAPIRequest struct {
	model.Params
	// 获取应用的权限套餐请求对象
	_getAppPermissionPackagesRequest *GetAppPermissionPackageRequest
}

// NewAlibabaMoziAclAppGetpermisspkgsRequest 初始化AlibabaMoziAclAppGetpermisspkgsAPIRequest对象
func NewAlibabaMoziAclAppGetpermisspkgsRequest() *AlibabaMoziAclAppGetpermisspkgsAPIRequest {
	return &AlibabaMoziAclAppGetpermisspkgsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclAppGetpermisspkgsAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.app.getpermisspkgs"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclAppGetpermisspkgsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziAclAppGetpermisspkgsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGetAppPermissionPackagesRequest is GetAppPermissionPackagesRequest Setter
// 获取应用的权限套餐请求对象
func (r *AlibabaMoziAclAppGetpermisspkgsAPIRequest) SetGetAppPermissionPackagesRequest(_getAppPermissionPackagesRequest *GetAppPermissionPackageRequest) error {
	r._getAppPermissionPackagesRequest = _getAppPermissionPackagesRequest
	r.Set("get_app_permission_packages_request", _getAppPermissionPackagesRequest)
	return nil
}

// GetGetAppPermissionPackagesRequest GetAppPermissionPackagesRequest Getter
func (r AlibabaMoziAclAppGetpermisspkgsAPIRequest) GetGetAppPermissionPackagesRequest() *GetAppPermissionPackageRequest {
	return r._getAppPermissionPackagesRequest
}
