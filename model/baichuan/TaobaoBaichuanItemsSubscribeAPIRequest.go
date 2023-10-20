package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanitemssubscribeAPIRequest 百川批量商品订阅 API请求
// taobao.baichuan.items.subscribe
//
// 百川批量添加订阅的商品
type TaobaobaichuanitemssubscribeAPIRequest struct {
	model.Params
	// 订阅的商品id列表
	_itemIds []string
}

// NewTaobaobaichuanitemssubscribeRequest 初始化TaobaobaichuanitemssubscribeAPIRequest对象
func NewTaobaobaichuanitemssubscribeRequest() *TaobaobaichuanitemssubscribeAPIRequest {
	return &TaobaobaichuanitemssubscribeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanitemssubscribeAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.items.subscribe"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanitemssubscribeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanitemssubscribeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 订阅的商品id列表
func (r *TaobaobaichuanitemssubscribeAPIRequest) SetItemIds(_itemIds []string) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaobaichuanitemssubscribeAPIRequest) GetItemIds() []string {
	return r._itemIds
}
