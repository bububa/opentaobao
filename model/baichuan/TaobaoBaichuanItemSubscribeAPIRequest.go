package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanitemsubscribeAPIRequest 单个商品订阅 API请求
// taobao.baichuan.item.subscribe
//
// 百川单个商品订阅
type TaobaobaichuanitemsubscribeAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTaobaobaichuanitemsubscribeRequest 初始化TaobaobaichuanitemsubscribeAPIRequest对象
func NewTaobaobaichuanitemsubscribeRequest() *TaobaobaichuanitemsubscribeAPIRequest {
	return &TaobaobaichuanitemsubscribeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanitemsubscribeAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.item.subscribe"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanitemsubscribeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanitemsubscribeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaobaichuanitemsubscribeAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaobaichuanitemsubscribeAPIRequest) GetItemId() int64 {
	return r._itemId
}
