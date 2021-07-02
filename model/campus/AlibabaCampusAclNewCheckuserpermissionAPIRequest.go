package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewCheckuserpermissionAPIRequest 校验用户是否有权限 API请求
// alibaba.campus.acl.new.checkuserpermission
//
// 校验用户是否有权限
type AlibabaCampusAclNewCheckuserpermissionAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 接口入参
	_checkUserPermissionParam *CheckUserPermissionParam
}

// NewAlibabaCampusAclNewCheckuserpermissionRequest 初始化AlibabaCampusAclNewCheckuserpermissionAPIRequest对象
func NewAlibabaCampusAclNewCheckuserpermissionRequest() *AlibabaCampusAclNewCheckuserpermissionAPIRequest {
	return &AlibabaCampusAclNewCheckuserpermissionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewCheckuserpermissionAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.checkuserpermission"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewCheckuserpermissionAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewCheckuserpermissionAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// Get Workbenchcontext Getter
func (r AlibabaCampusAclNewCheckuserpermissionAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// Set is CheckUserPermissionParam Setter
// 接口入参
func (r *AlibabaCampusAclNewCheckuserpermissionAPIRequest) SetCheckUserPermissionParam(_checkUserPermissionParam *CheckUserPermissionParam) error {
	r._checkUserPermissionParam = _checkUserPermissionParam
	r.Set("check_user_permission_param", _checkUserPermissionParam)
	return nil
}

// Get CheckUserPermissionParam Getter
func (r AlibabaCampusAclNewCheckuserpermissionAPIRequest) GetCheckUserPermissionParam() *CheckUserPermissionParam {
	return r._checkUserPermissionParam
}
