package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewListuserrolesAPIRequest 查询并标记用户选择的角色 API请求
// alibaba.campus.acl.new.listuserroles
//
// 查询并标记用户选择的角色
type AlibabaCampusAclNewListuserrolesAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 入参
	_param *UserRoleQueryParam
}

// NewAlibabaCampusAclNewListuserrolesRequest 初始化AlibabaCampusAclNewListuserrolesAPIRequest对象
func NewAlibabaCampusAclNewListuserrolesRequest() *AlibabaCampusAclNewListuserrolesAPIRequest {
	return &AlibabaCampusAclNewListuserrolesAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusAclNewListuserrolesAPIRequest) Reset() {
	r._workbenchcontext = nil
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewListuserrolesAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.listuserroles"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewListuserrolesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAclNewListuserrolesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewListuserrolesAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabaCampusAclNewListuserrolesAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// SetParam is Param Setter
// 入参
func (r *AlibabaCampusAclNewListuserrolesAPIRequest) SetParam(_param *UserRoleQueryParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaCampusAclNewListuserrolesAPIRequest) GetParam() *UserRoleQueryParam {
	return r._param
}

var poolAlibabaCampusAclNewListuserrolesAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusAclNewListuserrolesRequest()
	},
}

// GetAlibabaCampusAclNewListuserrolesRequest 从 sync.Pool 获取 AlibabaCampusAclNewListuserrolesAPIRequest
func GetAlibabaCampusAclNewListuserrolesAPIRequest() *AlibabaCampusAclNewListuserrolesAPIRequest {
	return poolAlibabaCampusAclNewListuserrolesAPIRequest.Get().(*AlibabaCampusAclNewListuserrolesAPIRequest)
}

// ReleaseAlibabaCampusAclNewListuserrolesAPIRequest 将 AlibabaCampusAclNewListuserrolesAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusAclNewListuserrolesAPIRequest(v *AlibabaCampusAclNewListuserrolesAPIRequest) {
	v.Reset()
	poolAlibabaCampusAclNewListuserrolesAPIRequest.Put(v)
}
