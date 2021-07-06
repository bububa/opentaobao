package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallItemPriceUpdateAPIRequest 天猫商品/SKU价格更新接口 API请求
// tmall.item.price.update
//
// 天猫商品/SKU价格更新接口，支持商品、SKU价格同时更新，支持同一商品下的SKU批量更新。
type TmallItemPriceUpdateAPIRequest struct {
	model.Params
	// 更新SKU价格时候的SKU价格对象；如果没有SKU或者不更新SKU价格，可以不填;查找SKU目前支持ID，属性串和商家编码三种模式，建议选用一种最合适的，切勿滥用，一次调用中如果混合使用，更新结果不可预期！
	_skuPrices []UpdateSkuPrice
	// 商品ID
	_itemId int64
	// 被更新商品价格
	_itemPrice float64
	// 商品价格更新时候的可选参数
	_options *UpdateItemPriceOption
}

// NewTmallItemPriceUpdateRequest 初始化TmallItemPriceUpdateAPIRequest对象
func NewTmallItemPriceUpdateRequest() *TmallItemPriceUpdateAPIRequest {
	return &TmallItemPriceUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemPriceUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.price.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemPriceUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSkuPrices is SkuPrices Setter
// 更新SKU价格时候的SKU价格对象；如果没有SKU或者不更新SKU价格，可以不填;查找SKU目前支持ID，属性串和商家编码三种模式，建议选用一种最合适的，切勿滥用，一次调用中如果混合使用，更新结果不可预期！
func (r *TmallItemPriceUpdateAPIRequest) SetSkuPrices(_skuPrices []UpdateSkuPrice) error {
	r._skuPrices = _skuPrices
	r.Set("sku_prices", _skuPrices)
	return nil
}

// GetSkuPrices SkuPrices Getter
func (r TmallItemPriceUpdateAPIRequest) GetSkuPrices() []UpdateSkuPrice {
	return r._skuPrices
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallItemPriceUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemPriceUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetItemPrice is ItemPrice Setter
// 被更新商品价格
func (r *TmallItemPriceUpdateAPIRequest) SetItemPrice(_itemPrice float64) error {
	r._itemPrice = _itemPrice
	r.Set("item_price", _itemPrice)
	return nil
}

// GetItemPrice ItemPrice Getter
func (r TmallItemPriceUpdateAPIRequest) GetItemPrice() float64 {
	return r._itemPrice
}

// SetOptions is Options Setter
// 商品价格更新时候的可选参数
func (r *TmallItemPriceUpdateAPIRequest) SetOptions(_options *UpdateItemPriceOption) error {
	r._options = _options
	r.Set("options", _options)
	return nil
}

// GetOptions Options Getter
func (r TmallItemPriceUpdateAPIRequest) GetOptions() *UpdateItemPriceOption {
	return r._options
}
