package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewCheckuserroleAPIRequest 校验用户是否有角色 API请求
// alibaba.campus.acl.new.checkuserrole
//
// 校验用户是否有角色
type AlibabaCampusAclNewCheckuserroleAPIRequest struct {
	model.Params
	// 用户账号
	_userId string
	// 角色key
	_roleKey string
	// 系统入参
	_workbenchcontext *WorkBenchContext
}

// NewAlibabaCampusAclNewCheckuserroleRequest 初始化AlibabaCampusAclNewCheckuserroleAPIRequest对象
func NewAlibabaCampusAclNewCheckuserroleRequest() *AlibabaCampusAclNewCheckuserroleAPIRequest {
	return &AlibabaCampusAclNewCheckuserroleAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusAclNewCheckuserroleAPIRequest) Reset() {
	r._userId = ""
	r._roleKey = ""
	r._workbenchcontext = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewCheckuserroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.checkuserrole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewCheckuserroleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAclNewCheckuserroleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 用户账号
func (r *AlibabaCampusAclNewCheckuserroleAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaCampusAclNewCheckuserroleAPIRequest) GetUserId() string {
	return r._userId
}

// SetRoleKey is RoleKey Setter
// 角色key
func (r *AlibabaCampusAclNewCheckuserroleAPIRequest) SetRoleKey(_roleKey string) error {
	r._roleKey = _roleKey
	r.Set("role_key", _roleKey)
	return nil
}

// GetRoleKey RoleKey Getter
func (r AlibabaCampusAclNewCheckuserroleAPIRequest) GetRoleKey() string {
	return r._roleKey
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewCheckuserroleAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabaCampusAclNewCheckuserroleAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

var poolAlibabaCampusAclNewCheckuserroleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusAclNewCheckuserroleRequest()
	},
}

// GetAlibabaCampusAclNewCheckuserroleRequest 从 sync.Pool 获取 AlibabaCampusAclNewCheckuserroleAPIRequest
func GetAlibabaCampusAclNewCheckuserroleAPIRequest() *AlibabaCampusAclNewCheckuserroleAPIRequest {
	return poolAlibabaCampusAclNewCheckuserroleAPIRequest.Get().(*AlibabaCampusAclNewCheckuserroleAPIRequest)
}

// ReleaseAlibabaCampusAclNewCheckuserroleAPIRequest 将 AlibabaCampusAclNewCheckuserroleAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusAclNewCheckuserroleAPIRequest(v *AlibabaCampusAclNewCheckuserroleAPIRequest) {
	v.Reset()
	poolAlibabaCampusAclNewCheckuserroleAPIRequest.Put(v)
}
