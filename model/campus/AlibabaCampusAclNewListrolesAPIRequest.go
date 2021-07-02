package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewListrolesAPIRequest 查询全部角色 API请求
// alibaba.campus.acl.new.listroles
//
// 查询全部角色
type AlibabaCampusAclNewListrolesAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 入参
	_rolequeryparam *RoleQueryParam
}

// NewAlibabaCampusAclNewListrolesRequest 初始化AlibabaCampusAclNewListrolesAPIRequest对象
func NewAlibabaCampusAclNewListrolesRequest() *AlibabaCampusAclNewListrolesAPIRequest {
	return &AlibabaCampusAclNewListrolesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewListrolesAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.listroles"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewListrolesAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewListrolesAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabaCampusAclNewListrolesAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// SetRolequeryparam is Rolequeryparam Setter
// 入参
func (r *AlibabaCampusAclNewListrolesAPIRequest) SetRolequeryparam(_rolequeryparam *RoleQueryParam) error {
	r._rolequeryparam = _rolequeryparam
	r.Set("rolequeryparam", _rolequeryparam)
	return nil
}

// GetRolequeryparam Rolequeryparam Getter
func (r AlibabaCampusAclNewListrolesAPIRequest) GetRolequeryparam() *RoleQueryParam {
	return r._rolequeryparam
}
