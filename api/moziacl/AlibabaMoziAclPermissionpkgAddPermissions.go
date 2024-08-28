package moziacl

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// AlibabaMoziAclPermissionpkgAddPermissions 权限套餐添加权限
// alibaba.mozi.acl.permissionpkg.add.permissions
//
// 此接口的功能为：将一批应用下的权限添加到该应用下的权限套餐中
func AlibabaMoziAclPermissionpkgAddPermissions(ctx context.Context, clt *core.SDKClient, req *moziacl.AlibabaMoziAclPermissionpkgAddPermissionsAPIRequest, resp *moziacl.AlibabaMoziAclPermissionpkgAddPermissionsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
