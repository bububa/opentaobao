package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页获取应用的权限套餐 API请求
alibaba.mozi.acl.app.getpermisspkgs

分页查询应用下的权限套餐列表
*/
type AlibabaMoziAclAppGetpermisspkgsRequest struct {
    model.Params
    // 获取应用的权限套餐请求对象
    getAppPermissionPackagesRequest   *GetAppPermissionPackageRequest
}

// 初始化AlibabaMoziAclAppGetpermisspkgsRequest对象
func NewAlibabaMoziAclAppGetpermisspkgsRequest() *AlibabaMoziAclAppGetpermisspkgsRequest{
    return &AlibabaMoziAclAppGetpermisspkgsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclAppGetpermisspkgsRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.app.getpermisspkgs"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclAppGetpermisspkgsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GetAppPermissionPackagesRequest Setter
// 获取应用的权限套餐请求对象
func (r *AlibabaMoziAclAppGetpermisspkgsRequest) SetGetAppPermissionPackagesRequest(getAppPermissionPackagesRequest *GetAppPermissionPackageRequest) error {
    r.getAppPermissionPackagesRequest = getAppPermissionPackagesRequest
    r.Set("get_app_permission_packages_request", getAppPermissionPackagesRequest)
    return nil
}

// GetAppPermissionPackagesRequest Getter
func (r AlibabaMoziAclAppGetpermisspkgsRequest) GetGetAppPermissionPackagesRequest() *GetAppPermissionPackageRequest {
    return r.getAppPermissionPackagesRequest
}
