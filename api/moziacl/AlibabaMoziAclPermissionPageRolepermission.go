package moziacl

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// AlibabaMoziAclPermissionPageRolepermission 分页查询角色下包含的权限列表
// alibaba.mozi.acl.permission.page.rolepermission
//
// 根据传入的角色name，分页查询该角色包含的权限列表
func AlibabaMoziAclPermissionPageRolepermission(ctx context.Context, clt *core.SDKClient, req *moziacl.AlibabaMoziAclPermissionPageRolepermissionAPIRequest, resp *moziacl.AlibabaMoziAclPermissionPageRolepermissionAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
