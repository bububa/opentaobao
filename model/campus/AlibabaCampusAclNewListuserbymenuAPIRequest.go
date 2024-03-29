package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewListuserbymenuAPIRequest 查询菜单下的人员 API请求
// alibaba.campus.acl.new.listuserbymenu
//
// 查询拥有菜单权限的用户
type AlibabaCampusAclNewListuserbymenuAPIRequest struct {
	model.Params
	// /workbench/space/application
	_menuUrl string
	// 系统入参
	_context *WorkBenchContext
}

// NewAlibabaCampusAclNewListuserbymenuRequest 初始化AlibabaCampusAclNewListuserbymenuAPIRequest对象
func NewAlibabaCampusAclNewListuserbymenuRequest() *AlibabaCampusAclNewListuserbymenuAPIRequest {
	return &AlibabaCampusAclNewListuserbymenuAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusAclNewListuserbymenuAPIRequest) Reset() {
	r._menuUrl = ""
	r._context = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewListuserbymenuAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.listuserbymenu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewListuserbymenuAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAclNewListuserbymenuAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMenuUrl is MenuUrl Setter
// /workbench/space/application
func (r *AlibabaCampusAclNewListuserbymenuAPIRequest) SetMenuUrl(_menuUrl string) error {
	r._menuUrl = _menuUrl
	r.Set("menu_url", _menuUrl)
	return nil
}

// GetMenuUrl MenuUrl Getter
func (r AlibabaCampusAclNewListuserbymenuAPIRequest) GetMenuUrl() string {
	return r._menuUrl
}

// SetContext is Context Setter
// 系统入参
func (r *AlibabaCampusAclNewListuserbymenuAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabaCampusAclNewListuserbymenuAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}

var poolAlibabaCampusAclNewListuserbymenuAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusAclNewListuserbymenuRequest()
	},
}

// GetAlibabaCampusAclNewListuserbymenuRequest 从 sync.Pool 获取 AlibabaCampusAclNewListuserbymenuAPIRequest
func GetAlibabaCampusAclNewListuserbymenuAPIRequest() *AlibabaCampusAclNewListuserbymenuAPIRequest {
	return poolAlibabaCampusAclNewListuserbymenuAPIRequest.Get().(*AlibabaCampusAclNewListuserbymenuAPIRequest)
}

// ReleaseAlibabaCampusAclNewListuserbymenuAPIRequest 将 AlibabaCampusAclNewListuserbymenuAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusAclNewListuserbymenuAPIRequest(v *AlibabaCampusAclNewListuserbymenuAPIRequest) {
	v.Reset()
	poolAlibabaCampusAclNewListuserbymenuAPIRequest.Put(v)
}
