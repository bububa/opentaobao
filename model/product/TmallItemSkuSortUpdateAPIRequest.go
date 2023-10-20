package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemskusortupdateAPIRequest 商品销售属性排序更新 API请求
// tmall.item.sku.sort.update
//
// 商品销售属性排序更新
type TmallitemskusortupdateAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
	// 销售属性排序集合
	_itemSalePropSort *ItemSalePropSort
}

// NewTmallitemskusortupdateRequest 初始化TmallitemskusortupdateAPIRequest对象
func NewTmallitemskusortupdateRequest() *TmallitemskusortupdateAPIRequest {
	return &TmallitemskusortupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemskusortupdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.sku.sort.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemskusortupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemskusortupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallitemskusortupdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallitemskusortupdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetItemSalePropSort is ItemSalePropSort Setter
// 销售属性排序集合
func (r *TmallitemskusortupdateAPIRequest) SetItemSalePropSort(_itemSalePropSort *ItemSalePropSort) error {
	r._itemSalePropSort = _itemSalePropSort
	r.Set("item_sale_prop_sort", _itemSalePropSort)
	return nil
}

// GetItemSalePropSort ItemSalePropSort Getter
func (r TmallitemskusortupdateAPIRequest) GetItemSalePropSort() *ItemSalePropSort {
	return r._itemSalePropSort
}
