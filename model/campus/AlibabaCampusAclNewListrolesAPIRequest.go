package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclnewlistrolesAPIRequest 查询全部角色 API请求
// alibaba.campus.acl.new.listroles
//
// 查询全部角色
type AlibabacampusaclnewlistrolesAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 入参
	_rolequeryparam *RoleQueryParam
}

// NewAlibabacampusaclnewlistrolesRequest 初始化AlibabacampusaclnewlistrolesAPIRequest对象
func NewAlibabacampusaclnewlistrolesRequest() *AlibabacampusaclnewlistrolesAPIRequest {
	return &AlibabacampusaclnewlistrolesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclnewlistrolesAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.listroles"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclnewlistrolesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclnewlistrolesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabacampusaclnewlistrolesAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabacampusaclnewlistrolesAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// SetRolequeryparam is Rolequeryparam Setter
// 入参
func (r *AlibabacampusaclnewlistrolesAPIRequest) SetRolequeryparam(_rolequeryparam *RoleQueryParam) error {
	r._rolequeryparam = _rolequeryparam
	r.Set("rolequeryparam", _rolequeryparam)
	return nil
}

// GetRolequeryparam Rolequeryparam Getter
func (r AlibabacampusaclnewlistrolesAPIRequest) GetRolequeryparam() *RoleQueryParam {
	return r._rolequeryparam
}
