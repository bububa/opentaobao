package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamoziaclpermissionpagerolepermissionAPIRequest 分页查询角色下包含的权限列表 API请求
// alibaba.mozi.acl.permission.page.rolepermission
//
// 根据传入的角色name，分页查询该角色包含的权限列表
type AlibabamoziaclpermissionpagerolepermissionAPIRequest struct {
	model.Params
	// 分页查询角色下包含的权限列表
	_pageRolePermisions *PageRolePermissionRequest
}

// NewAlibabamoziaclpermissionpagerolepermissionRequest 初始化AlibabamoziaclpermissionpagerolepermissionAPIRequest对象
func NewAlibabamoziaclpermissionpagerolepermissionRequest() *AlibabamoziaclpermissionpagerolepermissionAPIRequest {
	return &AlibabamoziaclpermissionpagerolepermissionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamoziaclpermissionpagerolepermissionAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.permission.page.rolepermission"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamoziaclpermissionpagerolepermissionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamoziaclpermissionpagerolepermissionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageRolePermisions is PageRolePermisions Setter
// 分页查询角色下包含的权限列表
func (r *AlibabamoziaclpermissionpagerolepermissionAPIRequest) SetPageRolePermisions(_pageRolePermisions *PageRolePermissionRequest) error {
	r._pageRolePermisions = _pageRolePermisions
	r.Set("page_role_permisions", _pageRolePermisions)
	return nil
}

// GetPageRolePermisions PageRolePermisions Getter
func (r AlibabamoziaclpermissionpagerolepermissionAPIRequest) GetPageRolePermisions() *PageRolePermissionRequest {
	return r._pageRolePermisions
}
