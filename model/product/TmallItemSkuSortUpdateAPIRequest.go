package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSkuSortUpdateAPIRequest 商品销售属性排序更新 API请求
// tmall.item.sku.sort.update
//
// 商品销售属性排序更新
type TmallItemSkuSortUpdateAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
	// 销售属性排序集合
	_itemSalePropSort *ItemSalePropSort
}

// NewTmallItemSkuSortUpdateRequest 初始化TmallItemSkuSortUpdateAPIRequest对象
func NewTmallItemSkuSortUpdateRequest() *TmallItemSkuSortUpdateAPIRequest {
	return &TmallItemSkuSortUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSkuSortUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.sku.sort.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSkuSortUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallItemSkuSortUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemSkuSortUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetItemSalePropSort is ItemSalePropSort Setter
// 销售属性排序集合
func (r *TmallItemSkuSortUpdateAPIRequest) SetItemSalePropSort(_itemSalePropSort *ItemSalePropSort) error {
	r._itemSalePropSort = _itemSalePropSort
	r.Set("item_sale_prop_sort", _itemSalePropSort)
	return nil
}

// GetItemSalePropSort ItemSalePropSort Getter
func (r TmallItemSkuSortUpdateAPIRequest) GetItemSalePropSort() *ItemSalePropSort {
	return r._itemSalePropSort
}
