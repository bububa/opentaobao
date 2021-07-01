package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziAclAppGetpermisspkgsAPIRequest
分页获取应用的权限套餐 API请求
alibaba.mozi.acl.app.getpermisspkgs

分页查询应用下的权限套餐列表 */
type AlibabaMoziAclAppGetpermisspkgsAPIRequest struct {
	model.Params
	// 获取应用的权限套餐请求对象
	_getAppPermissionPackagesRequest *GetAppPermissionPackageRequest
}

// New
