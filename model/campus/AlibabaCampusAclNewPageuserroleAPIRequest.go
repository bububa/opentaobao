package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclnewpageuserroleAPIRequest 分页查询管理员 API请求
// alibaba.campus.acl.new.pageuserrole
//
// 新增用户和角色的关系
type AlibabacampusaclnewpageuserroleAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 入参
	_usersRoleQueryParam *UsersRoleQueryParam
}

// NewAlibabacampusaclnewpageuserroleRequest 初始化AlibabacampusaclnewpageuserroleAPIRequest对象
func NewAlibabacampusaclnewpageuserroleRequest() *AlibabacampusaclnewpageuserroleAPIRequest {
	return &AlibabacampusaclnewpageuserroleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclnewpageuserroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.pageuserrole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclnewpageuserroleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclnewpageuserroleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabacampusaclnewpageuserroleAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabacampusaclnewpageuserroleAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// SetUsersRoleQueryParam is UsersRoleQueryParam Setter
// 入参
func (r *AlibabacampusaclnewpageuserroleAPIRequest) SetUsersRoleQueryParam(_usersRoleQueryParam *UsersRoleQueryParam) error {
	r._usersRoleQueryParam = _usersRoleQueryParam
	r.Set("users_role_query_param", _usersRoleQueryParam)
	return nil
}

// GetUsersRoleQueryParam UsersRoleQueryParam Getter
func (r AlibabacampusaclnewpageuserroleAPIRequest) GetUsersRoleQueryParam() *UsersRoleQueryParam {
	return r._usersRoleQueryParam
}
