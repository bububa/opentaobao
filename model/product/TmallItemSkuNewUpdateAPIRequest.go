package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemskunewupdateAPIRequest 更新sku销售属性标新状态 API请求
// tmall.item.sku.new.update
//
// 更新sku销售属性标新状态
type TmallitemskunewupdateAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
	// 商品销售属性集合
	_itemSalePropNew *ItemSalePropNew
}

// NewTmallitemskunewupdateRequest 初始化TmallitemskunewupdateAPIRequest对象
func NewTmallitemskunewupdateRequest() *TmallitemskunewupdateAPIRequest {
	return &TmallitemskunewupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemskunewupdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.sku.new.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemskunewupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemskunewupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallitemskunewupdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallitemskunewupdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetItemSalePropNew is ItemSalePropNew Setter
// 商品销售属性集合
func (r *TmallitemskunewupdateAPIRequest) SetItemSalePropNew(_itemSalePropNew *ItemSalePropNew) error {
	r._itemSalePropNew = _itemSalePropNew
	r.Set("item_sale_prop_new", _itemSalePropNew)
	return nil
}

// GetItemSalePropNew ItemSalePropNew Getter
func (r TmallitemskunewupdateAPIRequest) GetItemSalePropNew() *ItemSalePropNew {
	return r._itemSalePropNew
}
