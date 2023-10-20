package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmmenulistAPIRequest 获取特价菜单 API请求
// alibaba.alsc.crm.menu.list
//
// 获取特价菜单
type AlibabaalsccrmmenulistAPIRequest struct {
	model.Params
	// 获取特价菜单请求参数
	_menuOpenReq *MenuOpenReq
}

// NewAlibabaalsccrmmenulistRequest 初始化AlibabaalsccrmmenulistAPIRequest对象
func NewAlibabaalsccrmmenulistRequest() *AlibabaalsccrmmenulistAPIRequest {
	return &AlibabaalsccrmmenulistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmmenulistAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.menu.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmmenulistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmmenulistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMenuOpenReq is MenuOpenReq Setter
// 获取特价菜单请求参数
func (r *AlibabaalsccrmmenulistAPIRequest) SetMenuOpenReq(_menuOpenReq *MenuOpenReq) error {
	r._menuOpenReq = _menuOpenReq
	r.Set("menu_open_req", _menuOpenReq)
	return nil
}

// GetMenuOpenReq MenuOpenReq Getter
func (r AlibabaalsccrmmenulistAPIRequest) GetMenuOpenReq() *MenuOpenReq {
	return r._menuOpenReq
}
