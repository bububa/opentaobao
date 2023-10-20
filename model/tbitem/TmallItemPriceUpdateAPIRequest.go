package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitempriceupdateAPIRequest 天猫商品/SKU价格更新接口 API请求
// tmall.item.price.update
//
// 天猫商品/SKU价格更新接口，支持商品、SKU价格同时更新，支持同一商品下的SKU批量更新。
type TmallitempriceupdateAPIRequest struct {
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

// NewTmallitempriceupdateRequest 初始化TmallitempriceupdateAPIRequest对象
func NewTmallitempriceupdateRequest() *TmallitempriceupdateAPIRequest {
	return &TmallitempriceupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitempriceupdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.price.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitempriceupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitempriceupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuPrices is SkuPrices Setter
// 更新SKU价格时候的SKU价格对象；如果没有SKU或者不更新SKU价格，可以不填;查找SKU目前支持ID，属性串和商家编码三种模式，建议选用一种最合适的，切勿滥用，一次调用中如果混合使用，更新结果不可预期！
func (r *TmallitempriceupdateAPIRequest) SetSkuPrices(_skuPrices []UpdateSkuPrice) error {
	r._skuPrices = _skuPrices
	r.Set("sku_prices", _skuPrices)
	return nil
}

// GetSkuPrices SkuPrices Getter
func (r TmallitempriceupdateAPIRequest) GetSkuPrices() []UpdateSkuPrice {
	return r._skuPrices
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallitempriceupdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallitempriceupdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetItemPrice is ItemPrice Setter
// 被更新商品价格
func (r *TmallitempriceupdateAPIRequest) SetItemPrice(_itemPrice float64) error {
	r._itemPrice = _itemPrice
	r.Set("item_price", _itemPrice)
	return nil
}

// GetItemPrice ItemPrice Getter
func (r TmallitempriceupdateAPIRequest) GetItemPrice() float64 {
	return r._itemPrice
}

// SetOptions is Options Setter
// 商品价格更新时候的可选参数
func (r *TmallitempriceupdateAPIRequest) SetOptions(_options *UpdateItemPriceOption) error {
	r._options = _options
	r.Set("options", _options)
	return nil
}

// GetOptions Options Getter
func (r TmallitempriceupdateAPIRequest) GetOptions() *UpdateItemPriceOption {
	return r._options
}
