package moziacl

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziAclPermissionPageRolepermissionAPIRequest 分页查询角色下包含的权限列表 API请求
// alibaba.mozi.acl.permission.page.rolepermission
//
// 根据传入的角色name，分页查询该角色包含的权限列表
type AlibabaMoziAclPermissionPageRolepermissionAPIRequest struct {
	model.Params
	// 分页查询角色下包含的权限列表
	_pageRolePermisions *PageRolePermissionRequest
}

// NewAlibabaMoziAclPermissionPageRolepermissionRequest 初始化AlibabaMoziAclPermissionPageRolepermissionAPIRequest对象
func NewAlibabaMoziAclPermissionPageRolepermissionRequest() *AlibabaMoziAclPermissionPageRolepermissionAPIRequest {
	return &AlibabaMoziAclPermissionPageRolepermissionAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMoziAclPermissionPageRolepermissionAPIRequest) Reset() {
	r._pageRolePermisions = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclPermissionPageRolepermissionAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.permission.page.rolepermission"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclPermissionPageRolepermissionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziAclPermissionPageRolepermissionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageRolePermisions is PageRolePermisions Setter
// 分页查询角色下包含的权限列表
func (r *AlibabaMoziAclPermissionPageRolepermissionAPIRequest) SetPageRolePermisions(_pageRolePermisions *PageRolePermissionRequest) error {
	r._pageRolePermisions = _pageRolePermisions
	r.Set("page_role_permisions", _pageRolePermisions)
	return nil
}

// GetPageRolePermisions PageRolePermisions Getter
func (r AlibabaMoziAclPermissionPageRolepermissionAPIRequest) GetPageRolePermisions() *PageRolePermissionRequest {
	return r._pageRolePermisions
}

var poolAlibabaMoziAclPermissionPageRolepermissionAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMoziAclPermissionPageRolepermissionRequest()
	},
}

// GetAlibabaMoziAclPermissionPageRolepermissionRequest 从 sync.Pool 获取 AlibabaMoziAclPermissionPageRolepermissionAPIRequest
func GetAlibabaMoziAclPermissionPageRolepermissionAPIRequest() *AlibabaMoziAclPermissionPageRolepermissionAPIRequest {
	return poolAlibabaMoziAclPermissionPageRolepermissionAPIRequest.Get().(*AlibabaMoziAclPermissionPageRolepermissionAPIRequest)
}

// ReleaseAlibabaMoziAclPermissionPageRolepermissionAPIRequest 将 AlibabaMoziAclPermissionPageRolepermissionAPIRequest 放入 sync.Pool
func ReleaseAlibabaMoziAclPermissionPageRolepermissionAPIRequest(v *AlibabaMoziAclPermissionPageRolepermissionAPIRequest) {
	v.Reset()
	poolAlibabaMoziAclPermissionPageRolepermissionAPIRequest.Put(v)
}
