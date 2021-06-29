package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页获取应用的权限套餐 APIRequest
alibaba.mozi.acl.app.getpermisspkgs

分页查询应用下的权限套餐列表
*/
type AlibabaMoziAclAppGetpermisspkgsRequest struct {
    model.Params

    // 获取应用的权限套餐请求对象
    getAppPermissionPackagesRequest   *GetAppPermissionPackageRequest 

}

func NewAlibabaMoziAclAppGetpermisspkgsRequest() *AlibabaMoziAclAppGetpermisspkgsRequest{
    return &AlibabaMoziAclAppGetpermisspkgsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziAclAppGetpermisspkgsRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.app.getpermisspkgs"
}

func (r AlibabaMoziAclAppGetpermisspkgsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziAclAppGetpermisspkgsRequest) SetGetAppPermissionPackagesRequest(getAppPermissionPackagesRequest *GetAppPermissionPackageRequest) error {
    r.getAppPermissionPackagesRequest = getAppPermissionPackagesRequest
    r.Set("get_app_permission_packages_request", getAppPermissionPackagesRequest)
    return nil
}

func (r AlibabaMoziAclAppGetpermisspkgsRequest) GetGetAppPermissionPackagesRequest() *GetAppPermissionPackageRequest {
    return r.getAppPermissionPackagesRequest
}

