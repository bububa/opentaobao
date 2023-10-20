package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewUnfreezeroleAPIRequest 解冻角色 API请求
// alibaba.campus.acl.new.unfreezerole
//
// 解冻角色
type AlibabaCampusAclNewUnfreezeroleAPIRequest struct {
	model.Params
	// 系统参数
	_workbenchcontext *WorkBenchContext
	// 角色主键id
	_roleId int64
}

// NewAlibabaCampusAclNewUnfreezeroleRequest 初始化AlibabaCampusAclNewUnfreezeroleAPIRequest对象
func NewAlibabaCampusAclNewUnfreezeroleRequest() *AlibabaCampusAclNewUnfreezeroleAPIRequest {
	return &AlibabaCampusAclNewUnfreezeroleAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusAclNewUnfreezeroleAPIRequest) Reset() {
	r._workbenchcontext = nil
	r._roleId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewUnfreezeroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.unfreezerole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewUnfreezeroleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAclNewUnfreezeroleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统参数
func (r *AlibabaCampusAclNewUnfreezeroleAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabaCampusAclNewUnfreezeroleAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// SetRoleId is RoleId Setter
// 角色主键id
func (r *AlibabaCampusAclNewUnfreezeroleAPIRequest) SetRoleId(_roleId int64) error {
	r._roleId = _roleId
	r.Set("role_id", _roleId)
	return nil
}

// GetRoleId RoleId Getter
func (r AlibabaCampusAclNewUnfreezeroleAPIRequest) GetRoleId() int64 {
	return r._roleId
}

var poolAlibabaCampusAclNewUnfreezeroleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusAclNewUnfreezeroleRequest()
	},
}

// GetAlibabaCampusAclNewUnfreezeroleRequest 从 sync.Pool 获取 AlibabaCampusAclNewUnfreezeroleAPIRequest
func GetAlibabaCampusAclNewUnfreezeroleAPIRequest() *AlibabaCampusAclNewUnfreezeroleAPIRequest {
	return poolAlibabaCampusAclNewUnfreezeroleAPIRequest.Get().(*AlibabaCampusAclNewUnfreezeroleAPIRequest)
}

// ReleaseAlibabaCampusAclNewUnfreezeroleAPIRequest 将 AlibabaCampusAclNewUnfreezeroleAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusAclNewUnfreezeroleAPIRequest(v *AlibabaCampusAclNewUnfreezeroleAPIRequest) {
	v.Reset()
	poolAlibabaCampusAclNewUnfreezeroleAPIRequest.Put(v)
}
