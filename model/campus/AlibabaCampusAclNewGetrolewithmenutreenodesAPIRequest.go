package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest 根据角色id查询权限 API请求
// alibaba.campus.acl.new.getrolewithmenutreenodes
//
// 根据角色id查询权限
type AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest struct {
	model.Params
	// 系统参数
	_workbenchcontext *WorkBenchContext
	// 角色id
	_roleId int64
	// 是否查询全部类型权限
	_allPermission bool
}

// NewAlibabaCampusAclNewGetrolewithmenutreenodesRequest 初始化AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest对象
func NewAlibabaCampusAclNewGetrolewithmenutreenodesRequest() *AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest {
	return &AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest) Reset() {
	r._workbenchcontext = nil
	r._roleId = 0
	r._allPermission = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.getrolewithmenutreenodes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统参数
func (r *AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// SetRoleId is RoleId Setter
// 角色id
func (r *AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest) SetRoleId(_roleId int64) error {
	r._roleId = _roleId
	r.Set("role_id", _roleId)
	return nil
}

// GetRoleId RoleId Getter
func (r AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest) GetRoleId() int64 {
	return r._roleId
}

// SetAllPermission is AllPermission Setter
// 是否查询全部类型权限
func (r *AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest) SetAllPermission(_allPermission bool) error {
	r._allPermission = _allPermission
	r.Set("all_permission", _allPermission)
	return nil
}

// GetAllPermission AllPermission Getter
func (r AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest) GetAllPermission() bool {
	return r._allPermission
}

var poolAlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusAclNewGetrolewithmenutreenodesRequest()
	},
}

// GetAlibabaCampusAclNewGetrolewithmenutreenodesRequest 从 sync.Pool 获取 AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest
func GetAlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest() *AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest {
	return poolAlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest.Get().(*AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest)
}

// ReleaseAlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest 将 AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest(v *AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest) {
	v.Reset()
	poolAlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest.Put(v)
}
