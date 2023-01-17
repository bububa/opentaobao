package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSkuNewUpdateAPIRequest 更新sku销售属性标新状态 API请求
// tmall.item.sku.new.update
//
// 更新sku销售属性标新状态
type TmallItemSkuNewUpdateAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
	// 商品销售属性集合
	_itemSalePropNew *ItemSalePropNew
}

// NewTmallItemSkuNewUpdateRequest 初始化TmallItemSkuNewUpdateAPIRequest对象
func NewTmallItemSkuNewUpdateRequest() *TmallItemSkuNewUpdateAPIRequest {
	return &TmallItemSkuNewUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSkuNewUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.sku.new.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSkuNewUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemSkuNewUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallItemSkuNewUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemSkuNewUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetItemSalePropNew is ItemSalePropNew Setter
// 商品销售属性集合
func (r *TmallItemSkuNewUpdateAPIRequest) SetItemSalePropNew(_itemSalePropNew *ItemSalePropNew) error {
	r._itemSalePropNew = _itemSalePropNew
	r.Set("item_sale_prop_new", _itemSalePropNew)
	return nil
}

// GetItemSalePropNew ItemSalePropNew Getter
func (r TmallItemSkuNewUpdateAPIRequest) GetItemSalePropNew() *ItemSalePropNew {
	return r._itemSalePropNew
}
