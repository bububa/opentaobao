package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclnewcheckuserpermissionAPIRequest 校验用户是否有权限 API请求
// alibaba.campus.acl.new.checkuserpermission
//
// 校验用户是否有权限
type AlibabacampusaclnewcheckuserpermissionAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 接口入参
	_checkUserPermissionParam *CheckUserPermissionParam
}

// NewAlibabacampusaclnewcheckuserpermissionRequest 初始化AlibabacampusaclnewcheckuserpermissionAPIRequest对象
func NewAlibabacampusaclnewcheckuserpermissionRequest() *AlibabacampusaclnewcheckuserpermissionAPIRequest {
	return &AlibabacampusaclnewcheckuserpermissionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclnewcheckuserpermissionAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.checkuserpermission"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclnewcheckuserpermissionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclnewcheckuserpermissionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabacampusaclnewcheckuserpermissionAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabacampusaclnewcheckuserpermissionAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// SetCheckUserPermissionParam is CheckUserPermissionParam Setter
// 接口入参
func (r *AlibabacampusaclnewcheckuserpermissionAPIRequest) SetCheckUserPermissionParam(_checkUserPermissionParam *CheckUserPermissionParam) error {
	r._checkUserPermissionParam = _checkUserPermissionParam
	r.Set("check_user_permission_param", _checkUserPermissionParam)
	return nil
}

// GetCheckUserPermissionParam CheckUserPermissionParam Getter
func (r AlibabacampusaclnewcheckuserpermissionAPIRequest) GetCheckUserPermissionParam() *CheckUserPermissionParam {
	return r._checkUserPermissionParam
}
