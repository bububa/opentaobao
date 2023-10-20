package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclnewlistuserrolesAPIRequest 查询并标记用户选择的角色 API请求
// alibaba.campus.acl.new.listuserroles
//
// 查询并标记用户选择的角色
type AlibabacampusaclnewlistuserrolesAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 入参
	_param *UserRoleQueryParam
}

// NewAlibabacampusaclnewlistuserrolesRequest 初始化AlibabacampusaclnewlistuserrolesAPIRequest对象
func NewAlibabacampusaclnewlistuserrolesRequest() *AlibabacampusaclnewlistuserrolesAPIRequest {
	return &AlibabacampusaclnewlistuserrolesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclnewlistuserrolesAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.listuserroles"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclnewlistuserrolesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclnewlistuserrolesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabacampusaclnewlistuserrolesAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabacampusaclnewlistuserrolesAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// SetParam is Param Setter
// 入参
func (r *AlibabacampusaclnewlistuserrolesAPIRequest) SetParam(_param *UserRoleQueryParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabacampusaclnewlistuserrolesAPIRequest) GetParam() *UserRoleQueryParam {
	return r._param
}
