package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclnewgetrolewithmenutreenodesAPIRequest 根据角色id查询权限 API请求
// alibaba.campus.acl.new.getrolewithmenutreenodes
//
// 根据角色id查询权限
type AlibabacampusaclnewgetrolewithmenutreenodesAPIRequest struct {
	model.Params
	// 系统参数
	_workbenchcontext *WorkBenchContext
	// 角色id
	_roleId int64
	// 是否查询全部类型权限
	_allPermission bool
}

// NewAlibabacampusaclnewgetrolewithmenutreenodesRequest 初始化AlibabacampusaclnewgetrolewithmenutreenodesAPIRequest对象
func NewAlibabacampusaclnewgetrolewithmenutreenodesRequest() *AlibabacampusaclnewgetrolewithmenutreenodesAPIRequest {
	return &AlibabacampusaclnewgetrolewithmenutreenodesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclnewgetrolewithmenutreenodesAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.getrolewithmenutreenodes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclnewgetrolewithmenutreenodesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclnewgetrolewithmenutreenodesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统参数
func (r *AlibabacampusaclnewgetrolewithmenutreenodesAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabacampusaclnewgetrolewithmenutreenodesAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// SetRoleId is RoleId Setter
// 角色id
func (r *AlibabacampusaclnewgetrolewithmenutreenodesAPIRequest) SetRoleId(_roleId int64) error {
	r._roleId = _roleId
	r.Set("role_id", _roleId)
	return nil
}

// GetRoleId RoleId Getter
func (r AlibabacampusaclnewgetrolewithmenutreenodesAPIRequest) GetRoleId() int64 {
	return r._roleId
}

// SetAllPermission is AllPermission Setter
// 是否查询全部类型权限
func (r *AlibabacampusaclnewgetrolewithmenutreenodesAPIRequest) SetAllPermission(_allPermission bool) error {
	r._allPermission = _allPermission
	r.Set("all_permission", _allPermission)
	return nil
}

// GetAllPermission AllPermission Getter
func (r AlibabacampusaclnewgetrolewithmenutreenodesAPIRequest) GetAllPermission() bool {
	return r._allPermission
}
