package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewPageuserroleAPIRequest 分页查询管理员 API请求
// alibaba.campus.acl.new.pageuserrole
//
// 新增用户和角色的关系
type AlibabaCampusAclNewPageuserroleAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 入参
	_usersRoleQueryParam *UsersRoleQueryParam
}

// NewAlibabaCampusAclNewPageuserroleRequest 初始化AlibabaCampusAclNewPageuserroleAPIRequest对象
func NewAlibabaCampusAclNewPageuserroleRequest() *AlibabaCampusAclNewPageuserroleAPIRequest {
	return &AlibabaCampusAclNewPageuserroleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewPageuserroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.pageuserrole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewPageuserroleAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewPageuserroleAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// Get Workbenchcontext Getter
func (r AlibabaCampusAclNewPageuserroleAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// Set is UsersRoleQueryParam Setter
// 入参
func (r *AlibabaCampusAclNewPageuserroleAPIRequest) SetUsersRoleQueryParam(_usersRoleQueryParam *UsersRoleQueryParam) error {
	r._usersRoleQueryParam = _usersRoleQueryParam
	r.Set("users_role_query_param", _usersRoleQueryParam)
	return nil
}

// Get UsersRoleQueryParam Getter
func (r AlibabaCampusAclNewPageuserroleAPIRequest) GetUsersRoleQueryParam() *UsersRoleQueryParam {
	return r._usersRoleQueryParam
}
