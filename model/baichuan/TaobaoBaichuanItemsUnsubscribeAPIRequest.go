package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanitemsunsubscribeAPIRequest 批量删除商品订阅 API请求
// taobao.baichuan.items.unsubscribe
//
// 批量删除商品订阅
type TaobaobaichuanitemsunsubscribeAPIRequest struct {
	model.Params
	// 删除的商品id
	_itemIds []string
}

// NewTaobaobaichuanitemsunsubscribeRequest 初始化TaobaobaichuanitemsunsubscribeAPIRequest对象
func NewTaobaobaichuanitemsunsubscribeRequest() *TaobaobaichuanitemsunsubscribeAPIRequest {
	return &TaobaobaichuanitemsunsubscribeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanitemsunsubscribeAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.items.unsubscribe"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanitemsunsubscribeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanitemsunsubscribeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 删除的商品id
func (r *TaobaobaichuanitemsunsubscribeAPIRequest) SetItemIds(_itemIds []string) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaobaichuanitemsunsubscribeAPIRequest) GetItemIds() []string {
	return r._itemIds
}
