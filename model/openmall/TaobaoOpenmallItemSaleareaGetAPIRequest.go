package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallItemSaleareaGetAPIRequest 查询商品可售区域 API请求
// taobao.openmall.item.salearea.get
//
// 获取openmall商品的可售区域
type TaobaoOpenmallItemSaleareaGetAPIRequest struct {
	model.Params
	// 商品SKU
	_skuIds string
	// 商品ID
	_itemId int64
}

// NewTaobaoOpenmallItemSaleareaGetRequest 初始化TaobaoOpenmallItemSaleareaGetAPIRequest对象
func NewTaobaoOpenmallItemSaleareaGetRequest() *TaobaoOpenmallItemSaleareaGetAPIRequest {
	return &TaobaoOpenmallItemSaleareaGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallItemSaleareaGetAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.item.salearea.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallItemSaleareaGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SkuIds Setter
// 商品SKU
func (r *TaobaoOpenmallItemSaleareaGetAPIRequest) SetSkuIds(_skuIds string) error {
	r._skuIds = _skuIds
	r.Set("sku_ids", _skuIds)
	return nil
}

// Get SkuIds Getter
func (r TaobaoOpenmallItemSaleareaGetAPIRequest) GetSkuIds() string {
	return r._skuIds
}

// Set is ItemId Setter
// 商品ID
func (r *TaobaoOpenmallItemSaleareaGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoOpenmallItemSaleareaGetAPIRequest) GetItemId() int64 {
	return r._itemId
}
