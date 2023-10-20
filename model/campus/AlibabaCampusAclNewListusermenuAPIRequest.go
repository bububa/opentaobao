package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewListusermenuAPIRequest 查询用户有权限的菜单树 API请求
// alibaba.campus.acl.new.listusermenu
//
// 查询用户有权限的菜单树
type AlibabaCampusAclNewListusermenuAPIRequest struct {
	model.Params
	// 用户账号
	_userId string
	// 系统入参
	_workbenchcontext *WorkBenchContext
}

// NewAlibabaCampusAclNewListusermenuRequest 初始化AlibabaCampusAclNewListusermenuAPIRequest对象
func NewAlibabaCampusAclNewListusermenuRequest() *AlibabaCampusAclNewListusermenuAPIRequest {
	return &AlibabaCampusAclNewListusermenuAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusAclNewListusermenuAPIRequest) Reset() {
	r._userId = ""
	r._workbenchcontext = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewListusermenuAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.listusermenu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewListusermenuAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAclNewListusermenuAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 用户账号
func (r *AlibabaCampusAclNewListusermenuAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaCampusAclNewListusermenuAPIRequest) GetUserId() string {
	return r._userId
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewListusermenuAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabaCampusAclNewListusermenuAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

var poolAlibabaCampusAclNewListusermenuAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusAclNewListusermenuRequest()
	},
}

// GetAlibabaCampusAclNewListusermenuRequest 从 sync.Pool 获取 AlibabaCampusAclNewListusermenuAPIRequest
func GetAlibabaCampusAclNewListusermenuAPIRequest() *AlibabaCampusAclNewListusermenuAPIRequest {
	return poolAlibabaCampusAclNewListusermenuAPIRequest.Get().(*AlibabaCampusAclNewListusermenuAPIRequest)
}

// ReleaseAlibabaCampusAclNewListusermenuAPIRequest 将 AlibabaCampusAclNewListusermenuAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusAclNewListusermenuAPIRequest(v *AlibabaCampusAclNewListusermenuAPIRequest) {
	v.Reset()
	poolAlibabaCampusAclNewListusermenuAPIRequest.Put(v)
}
