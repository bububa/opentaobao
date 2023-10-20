package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmMenuListAPIRequest 获取特价菜单 API请求
// alibaba.alsc.crm.menu.list
//
// 获取特价菜单
type AlibabaAlscCrmMenuListAPIRequest struct {
	model.Params
	// 获取特价菜单请求参数
	_menuOpenReq *MenuOpenReq
}

// NewAlibabaAlscCrmMenuListRequest 初始化AlibabaAlscCrmMenuListAPIRequest对象
func NewAlibabaAlscCrmMenuListRequest() *AlibabaAlscCrmMenuListAPIRequest {
	return &AlibabaAlscCrmMenuListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmMenuListAPIRequest) Reset() {
	r._menuOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmMenuListAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.menu.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmMenuListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmMenuListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMenuOpenReq is MenuOpenReq Setter
// 获取特价菜单请求参数
func (r *AlibabaAlscCrmMenuListAPIRequest) SetMenuOpenReq(_menuOpenReq *MenuOpenReq) error {
	r._menuOpenReq = _menuOpenReq
	r.Set("menu_open_req", _menuOpenReq)
	return nil
}

// GetMenuOpenReq MenuOpenReq Getter
func (r AlibabaAlscCrmMenuListAPIRequest) GetMenuOpenReq() *MenuOpenReq {
	return r._menuOpenReq
}

var poolAlibabaAlscCrmMenuListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmMenuListRequest()
	},
}

// GetAlibabaAlscCrmMenuListRequest 从 sync.Pool 获取 AlibabaAlscCrmMenuListAPIRequest
func GetAlibabaAlscCrmMenuListAPIRequest() *AlibabaAlscCrmMenuListAPIRequest {
	return poolAlibabaAlscCrmMenuListAPIRequest.Get().(*AlibabaAlscCrmMenuListAPIRequest)
}

// ReleaseAlibabaAlscCrmMenuListAPIRequest 将 AlibabaAlscCrmMenuListAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmMenuListAPIRequest(v *AlibabaAlscCrmMenuListAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmMenuListAPIRequest.Put(v)
}
