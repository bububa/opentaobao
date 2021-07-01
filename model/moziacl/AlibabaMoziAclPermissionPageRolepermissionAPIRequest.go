package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziAclPermissionPageRolepermissionAPIRequest
分页查询角色下包含的权限列表 API请求
alibaba.mozi.acl.permission.page.rolepermission

根据传入的角色name，分页查询该角色包含的权限列表 */
type AlibabaMoziAclPermissionPageRolepermissionAPIRequest struct {
	model.Params
	// 分页查询角色下包含的权限列表
	_pageRolePermisions *PageRolePermissionRequest
}

// New
