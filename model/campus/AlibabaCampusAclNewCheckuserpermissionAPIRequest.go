package campus

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusAclNewCheckuserpermissionAPIRequest) Reset() {
	r._workbenchcontext = nil
	r._checkUserPermissionParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewCheckuserpermissionAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.checkuserpermission"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewCheckuserpermissionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAclNewCheckuserpermissionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewCheckuserpermissionAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabaCampusAclNewCheckuserpermissionAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// SetCheckUserPermissionParam is CheckUserPermissionParam Setter
// 接口入参
func (r *AlibabaCampusAclNewCheckuserpermissionAPIRequest) SetCheckUserPermissionParam(_checkUserPermissionParam *CheckUserPermissionParam) error {
	r._checkUserPermissionParam = _checkUserPermissionParam
	r.Set("check_user_permission_param", _checkUserPermissionParam)
	return nil
}

// GetCheckUserPermissionParam CheckUserPermissionParam Getter
func (r AlibabaCampusAclNewCheckuserpermissionAPIRequest) GetCheckUserPermissionParam() *CheckUserPermissionParam {
	return r._checkUserPermissionParam
}

var poolAlibabaCampusAclNewCheckuserpermissionAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusAclNewCheckuserpermissionRequest()
	},
}

// GetAlibabaCampusAclNewCheckuserpermissionRequest 从 sync.Pool 获取 AlibabaCampusAclNewCheckuserpermissionAPIRequest
func GetAlibabaCampusAclNewCheckuserpermissionAPIRequest() *AlibabaCampusAclNewCheckuserpermissionAPIRequest {
	return poolAlibabaCampusAclNewCheckuserpermissionAPIRequest.Get().(*AlibabaCampusAclNewCheckuserpermissionAPIRequest)
}

// ReleaseAlibabaCampusAclNewCheckuserpermissionAPIRequest 将 AlibabaCampusAclNewCheckuserpermissionAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusAclNewCheckuserpermissionAPIRequest(v *AlibabaCampusAclNewCheckuserpermissionAPIRequest) {
	v.Reset()
	poolAlibabaCampusAclNewCheckuserpermissionAPIRequest.Put(v)
}
