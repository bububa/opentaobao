package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbitemgetAPIRequest 根据商品ID获取商品信息 API请求
// taobao.wlb.item.get
//
// 根据商品ID获取商品信息,除了获取商品信息外还可获取商品属性信息和库存信息。
type TaobaowlbitemgetAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewTaobaowlbitemgetRequest 初始化TaobaowlbitemgetAPIRequest对象
func NewTaobaowlbitemgetRequest() *TaobaowlbitemgetAPIRequest {
	return &TaobaowlbitemgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbitemgetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.item.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbitemgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbitemgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TaobaowlbitemgetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaowlbitemgetAPIRequest) GetItemId() int64 {
	return r._itemId
}
