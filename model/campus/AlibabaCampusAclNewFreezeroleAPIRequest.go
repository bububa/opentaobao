package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclnewfreezeroleAPIRequest 冻结角色 API请求
// alibaba.campus.acl.new.freezerole
//
// 冻结角色
type AlibabacampusaclnewfreezeroleAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 角色主键id
	_roleId int64
}

// NewAlibabacampusaclnewfreezeroleRequest 初始化AlibabacampusaclnewfreezeroleAPIRequest对象
func NewAlibabacampusaclnewfreezeroleRequest() *AlibabacampusaclnewfreezeroleAPIRequest {
	return &AlibabacampusaclnewfreezeroleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclnewfreezeroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.freezerole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclnewfreezeroleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclnewfreezeroleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabacampusaclnewfreezeroleAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabacampusaclnewfreezeroleAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// SetRoleId is RoleId Setter
// 角色主键id
func (r *AlibabacampusaclnewfreezeroleAPIRequest) SetRoleId(_roleId int64) error {
	r._roleId = _roleId
	r.Set("role_id", _roleId)
	return nil
}

// GetRoleId RoleId Getter
func (r AlibabacampusaclnewfreezeroleAPIRequest) GetRoleId() int64 {
	return r._roleId
}
