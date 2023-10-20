package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappvirtualitemgetAPIRequest 小程序关联虚拟商品查询 API请求
// taobao.miniapp.virtual.item.get
//
// 小程序关联虚拟商品查询
type TaobaominiappvirtualitemgetAPIRequest struct {
	model.Params
	// 请求
	_itemRequest *MiniappItemRequest
}

// NewTaobaominiappvirtualitemgetRequest 初始化TaobaominiappvirtualitemgetAPIRequest对象
func NewTaobaominiappvirtualitemgetRequest() *TaobaominiappvirtualitemgetAPIRequest {
	return &TaobaominiappvirtualitemgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappvirtualitemgetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.virtual.item.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappvirtualitemgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappvirtualitemgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemRequest is ItemRequest Setter
// 请求
func (r *TaobaominiappvirtualitemgetAPIRequest) SetItemRequest(_itemRequest *MiniappItemRequest) error {
	r._itemRequest = _itemRequest
	r.Set("item_request", _itemRequest)
	return nil
}

// GetItemRequest ItemRequest Getter
func (r TaobaominiappvirtualitemgetAPIRequest) GetItemRequest() *MiniappItemRequest {
	return r._itemRequest
}
