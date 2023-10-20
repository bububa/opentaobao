package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmallitemsaleareagetAPIRequest 查询商品可售区域 API请求
// taobao.openmall.item.salearea.get
//
// 获取openmall商品的可售区域
type TaobaoopenmallitemsaleareagetAPIRequest struct {
	model.Params
	// 商品SKU
	_skuIds string
	// 商品ID
	_itemId int64
}

// NewTaobaoopenmallitemsaleareagetRequest 初始化TaobaoopenmallitemsaleareagetAPIRequest对象
func NewTaobaoopenmallitemsaleareagetRequest() *TaobaoopenmallitemsaleareagetAPIRequest {
	return &TaobaoopenmallitemsaleareagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenmallitemsaleareagetAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.item.salearea.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenmallitemsaleareagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenmallitemsaleareagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuIds is SkuIds Setter
// 商品SKU
func (r *TaobaoopenmallitemsaleareagetAPIRequest) SetSkuIds(_skuIds string) error {
	r._skuIds = _skuIds
	r.Set("sku_ids", _skuIds)
	return nil
}

// GetSkuIds SkuIds Getter
func (r TaobaoopenmallitemsaleareagetAPIRequest) GetSkuIds() string {
	return r._skuIds
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TaobaoopenmallitemsaleareagetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoopenmallitemsaleareagetAPIRequest) GetItemId() int64 {
	return r._itemId
}
