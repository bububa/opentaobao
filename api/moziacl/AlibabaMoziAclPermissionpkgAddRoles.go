package moziacl

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// AlibabaMoziAclPermissionpkgAddRoles 将角色添加到权限套餐中
// alibaba.mozi.acl.permissionpkg.add.roles
//
// 此接口是将应用下的一批角色添加到该应用的某个权限套餐中
func AlibabaMoziAclPermissionpkgAddRoles(ctx context.Context, clt *core.SDKClient, req *moziacl.AlibabaMoziAclPermissionpkgAddRolesAPIRequest, resp *moziacl.AlibabaMoziAclPermissionpkgAddRolesAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
