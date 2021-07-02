package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewListuserbymenuAPIRequest 查询菜单下的人员 API请求
// alibaba.campus.acl.new.listuserbymenu
//
// 查询拥有菜单权限的用户
type AlibabaCampusAclNewListuserbymenuAPIRequest struct {
	model.Params
	// 系统入参
	_context *WorkBenchContext
	// /workbench/space/application
	_menuUrl string
}

// NewAlibabaCampusAclNewListuserbymenuRequest 初始化AlibabaCampusAclNewListuserbymenuAPIRequest对象
func NewAlibabaCampusAclNewListuserbymenuRequest() *AlibabaCampusAclNewListuserbymenuAPIRequest {
	return &AlibabaCampusAclNewListuserbymenuAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewListuserbymenuAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.listuserbymenu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewListuserbymenuAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
