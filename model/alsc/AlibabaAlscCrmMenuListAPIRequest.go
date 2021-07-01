package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmMenuListAPIRequest
获取特价菜单 API请求
alibaba.alsc.crm.menu.list

获取特价菜单 */
type AlibabaAlscCrmMenuListAPIRequest struct {
	model.Params
	// 获取特价菜单请求参数
	_menuOpenReq *MenuOpenReq
}

// NewAlibabaAlscCrmMenuListRequest 初始化AlibabaAlscCrmMenuListAPIRequest对象
func NewAlibabaAlscCrmMenuListRequest() *AlibabaAlscCrmMenuListAPIRequest {
	return &AlibabaAlscCrmMenuListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmMenuListAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.menu.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmMenuListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MenuOpenReq Setter
// 获取特价菜单请求参数
func (r *AlibabaAlscCrmMenuListAPIRequest) SetMenuOpenReq(_menuOpenReq *MenuOpenReq) error {
	r._menuOpenReq = _menuOpenReq
	r.Set("menu_open_req", _menuOpenReq)
	return nil
}

// Get MenuOpenReq Getter
func (r AlibabaAlscCrmMenuListAPIRequest) GetMenuOpenReq() *MenuOpenReq {
	return r._menuOpenReq
}
