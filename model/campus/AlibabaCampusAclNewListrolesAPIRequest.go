package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewListrolesAPIRequest 查询全部角色 API请求
// alibaba.campus.acl.new.listroles
//
// 查询全部角色
type AlibabaCampusAclNewListrolesAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 入参
	_rolequeryparam *RoleQueryParam
}

// NewAlibabaCampusAclNewListrolesRequest 初始化AlibabaCampusAclNewListrolesAPIRequest对象
func NewAlibabaCampusAclNewListrolesRequest() *AlibabaCampusAclNewListrolesAPIRequest {
	return &AlibabaCampusAclNewListrolesAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusAclNewListrolesAPIRequest) Reset() {
	r._workbenchcontext = nil
	r._rolequeryparam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewListrolesAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.listroles"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewListrolesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAclNewListrolesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewListrolesAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabaCampusAclNewListrolesAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// SetRolequeryparam is Rolequeryparam Setter
// 入参
func (r *AlibabaCampusAclNewListrolesAPIRequest) SetRolequeryparam(_rolequeryparam *RoleQueryParam) error {
	r._rolequeryparam = _rolequeryparam
	r.Set("rolequeryparam", _rolequeryparam)
	return nil
}

// GetRolequeryparam Rolequeryparam Getter
func (r AlibabaCampusAclNewListrolesAPIRequest) GetRolequeryparam() *RoleQueryParam {
	return r._rolequeryparam
}

var poolAlibabaCampusAclNewListrolesAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusAclNewListrolesRequest()
	},
}

// GetAlibabaCampusAclNewListrolesRequest 从 sync.Pool 获取 AlibabaCampusAclNewListrolesAPIRequest
func GetAlibabaCampusAclNewListrolesAPIRequest() *AlibabaCampusAclNewListrolesAPIRequest {
	return poolAlibabaCampusAclNewListrolesAPIRequest.Get().(*AlibabaCampusAclNewListrolesAPIRequest)
}

// ReleaseAlibabaCampusAclNewListrolesAPIRequest 将 AlibabaCampusAclNewListrolesAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusAclNewListrolesAPIRequest(v *AlibabaCampusAclNewListrolesAPIRequest) {
	v.Reset()
	poolAlibabaCampusAclNewListrolesAPIRequest.Put(v)
}
