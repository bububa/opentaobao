package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewGetappmenutreeAPIRequest 查询应用下的菜单树 API请求
// alibaba.campus.acl.new.getappmenutree
//
// 查询应用下的菜单树
type AlibabaCampusAclNewGetappmenutreeAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 是否关联查询出菜单下的权限
	_withpermission bool
}

// NewAlibabaCampusAclNewGetappmenutreeRequest 初始化AlibabaCampusAclNewGetappmenutreeAPIRequest对象
func NewAlibabaCampusAclNewGetappmenutreeRequest() *AlibabaCampusAclNewGetappmenutreeAPIRequest {
	return &AlibabaCampusAclNewGetappmenutreeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusAclNewGetappmenutreeAPIRequest) Reset() {
	r._workbenchcontext = nil
	r._withpermission = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewGetappmenutreeAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.getappmenutree"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewGetappmenutreeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAclNewGetappmenutreeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewGetappmenutreeAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabaCampusAclNewGetappmenutreeAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// SetWithpermission is Withpermission Setter
// 是否关联查询出菜单下的权限
func (r *AlibabaCampusAclNewGetappmenutreeAPIRequest) SetWithpermission(_withpermission bool) error {
	r._withpermission = _withpermission
	r.Set("withpermission", _withpermission)
	return nil
}

// GetWithpermission Withpermission Getter
func (r AlibabaCampusAclNewGetappmenutreeAPIRequest) GetWithpermission() bool {
	return r._withpermission
}

var poolAlibabaCampusAclNewGetappmenutreeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusAclNewGetappmenutreeRequest()
	},
}

// GetAlibabaCampusAclNewGetappmenutreeRequest 从 sync.Pool 获取 AlibabaCampusAclNewGetappmenutreeAPIRequest
func GetAlibabaCampusAclNewGetappmenutreeAPIRequest() *AlibabaCampusAclNewGetappmenutreeAPIRequest {
	return poolAlibabaCampusAclNewGetappmenutreeAPIRequest.Get().(*AlibabaCampusAclNewGetappmenutreeAPIRequest)
}

// ReleaseAlibabaCampusAclNewGetappmenutreeAPIRequest 将 AlibabaCampusAclNewGetappmenutreeAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusAclNewGetappmenutreeAPIRequest(v *AlibabaCampusAclNewGetappmenutreeAPIRequest) {
	v.Reset()
	poolAlibabaCampusAclNewGetappmenutreeAPIRequest.Put(v)
}
