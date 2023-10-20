package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclnewlistuserbymenuAPIRequest 查询菜单下的人员 API请求
// alibaba.campus.acl.new.listuserbymenu
//
// 查询拥有菜单权限的用户
type AlibabacampusaclnewlistuserbymenuAPIRequest struct {
	model.Params
	// /workbench/space/application
	_menuUrl string
	// 系统入参
	_context *WorkBenchContext
}

// NewAlibabacampusaclnewlistuserbymenuRequest 初始化AlibabacampusaclnewlistuserbymenuAPIRequest对象
func NewAlibabacampusaclnewlistuserbymenuRequest() *AlibabacampusaclnewlistuserbymenuAPIRequest {
	return &AlibabacampusaclnewlistuserbymenuAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclnewlistuserbymenuAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.listuserbymenu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclnewlistuserbymenuAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclnewlistuserbymenuAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMenuUrl is MenuUrl Setter
// /workbench/space/application
func (r *AlibabacampusaclnewlistuserbymenuAPIRequest) SetMenuUrl(_menuUrl string) error {
	r._menuUrl = _menuUrl
	r.Set("menu_url", _menuUrl)
	return nil
}

// GetMenuUrl MenuUrl Getter
func (r AlibabacampusaclnewlistuserbymenuAPIRequest) GetMenuUrl() string {
	return r._menuUrl
}

// SetContext is Context Setter
// 系统入参
func (r *AlibabacampusaclnewlistuserbymenuAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabacampusaclnewlistuserbymenuAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}
