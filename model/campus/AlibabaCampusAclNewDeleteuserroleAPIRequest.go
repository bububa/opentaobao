package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewDeleteuserroleAPIRequest 删除管理员 API请求
// alibaba.campus.acl.new.deleteuserrole
//
// 删除管理员
type AlibabaCampusAclNewDeleteuserroleAPIRequest struct {
	model.Params
	// 角色id
	_roleIds []string
	// 用户账号
	_userId string
	// 系统入参
	_workbenchcontext *WorkBenchContext
}

// NewAlibabaCampusAclNewDeleteuserroleRequest 初始化AlibabaCampusAclNewDeleteuserroleAPIRequest对象
func NewAlibabaCampusAclNewDeleteuserroleRequest() *AlibabaCampusAclNewDeleteuserroleAPIRequest {
	return &AlibabaCampusAclNewDeleteuserroleAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusAclNewDeleteuserroleAPIRequest) Reset() {
	r._roleIds = r._roleIds[:0]
	r._userId = ""
	r._workbenchcontext = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewDeleteuserroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.deleteuserrole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewDeleteuserroleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAclNewDeleteuserroleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRoleIds is RoleIds Setter
// 角色id
func (r *AlibabaCampusAclNewDeleteuserroleAPIRequest) SetRoleIds(_roleIds []string) error {
	r._roleIds = _roleIds
	r.Set("role_ids", _roleIds)
	return nil
}

// GetRoleIds RoleIds Getter
func (r AlibabaCampusAclNewDeleteuserroleAPIRequest) GetRoleIds() []string {
	return r._roleIds
}

// SetUserId is UserId Setter
// 用户账号
func (r *AlibabaCampusAclNewDeleteuserroleAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaCampusAclNewDeleteuserroleAPIRequest) GetUserId() string {
	return r._userId
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewDeleteuserroleAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabaCampusAclNewDeleteuserroleAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

var poolAlibabaCampusAclNewDeleteuserroleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusAclNewDeleteuserroleRequest()
	},
}

// GetAlibabaCampusAclNewDeleteuserroleRequest 从 sync.Pool 获取 AlibabaCampusAclNewDeleteuserroleAPIRequest
func GetAlibabaCampusAclNewDeleteuserroleAPIRequest() *AlibabaCampusAclNewDeleteuserroleAPIRequest {
	return poolAlibabaCampusAclNewDeleteuserroleAPIRequest.Get().(*AlibabaCampusAclNewDeleteuserroleAPIRequest)
}

// ReleaseAlibabaCampusAclNewDeleteuserroleAPIRequest 将 AlibabaCampusAclNewDeleteuserroleAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusAclNewDeleteuserroleAPIRequest(v *AlibabaCampusAclNewDeleteuserroleAPIRequest) {
	v.Reset()
	poolAlibabaCampusAclNewDeleteuserroleAPIRequest.Put(v)
}
