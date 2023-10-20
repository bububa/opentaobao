package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclnewunfreezeroleAPIRequest 解冻角色 API请求
// alibaba.campus.acl.new.unfreezerole
//
// 解冻角色
type AlibabacampusaclnewunfreezeroleAPIRequest struct {
	model.Params
	// 系统参数
	_workbenchcontext *WorkBenchContext
	// 角色主键id
	_roleId int64
}

// NewAlibabacampusaclnewunfreezeroleRequest 初始化AlibabacampusaclnewunfreezeroleAPIRequest对象
func NewAlibabacampusaclnewunfreezeroleRequest() *AlibabacampusaclnewunfreezeroleAPIRequest {
	return &AlibabacampusaclnewunfreezeroleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclnewunfreezeroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.unfreezerole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclnewunfreezeroleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclnewunfreezeroleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统参数
func (r *AlibabacampusaclnewunfreezeroleAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabacampusaclnewunfreezeroleAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// SetRoleId is RoleId Setter
// 角色主键id
func (r *AlibabacampusaclnewunfreezeroleAPIRequest) SetRoleId(_roleId int64) error {
	r._roleId = _roleId
	r.Set("role_id", _roleId)
	return nil
}

// GetRoleId RoleId Getter
func (r AlibabacampusaclnewunfreezeroleAPIRequest) GetRoleId() int64 {
	return r._roleId
}
