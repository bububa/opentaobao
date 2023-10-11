package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappVirtualItemGetAPIRequest 小程序关联虚拟商品查询 API请求
// taobao.miniapp.virtual.item.get
//
// 小程序关联虚拟商品查询
type TaobaoMiniappVirtualItemGetAPIRequest struct {
	model.Params
	// 请求
	_itemRequest *MiniappItemRequest
}

// NewTaobaoMiniappVirtualItemGetRequest 初始化TaobaoMiniappVirtualItemGetAPIRequest对象
func NewTaobaoMiniappVirtualItemGetRequest() *TaobaoMiniappVirtualItemGetAPIRequest {
	return &TaobaoMiniappVirtualItemGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappVirtualItemGetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.virtual.item.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappVirtualItemGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappVirtualItemGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemRequest is ItemRequest Setter
// 请求
func (r *TaobaoMiniappVirtualItemGetAPIRequest) SetItemRequest(_itemRequest *MiniappItemRequest) error {
	r._itemRequest = _itemRequest
	r.Set("item_request", _itemRequest)
	return nil
}

// GetItemRequest ItemRequest Getter
func (r TaobaoMiniappVirtualItemGetAPIRequest) GetItemRequest() *MiniappItemRequest {
	return r._itemRequest
}
