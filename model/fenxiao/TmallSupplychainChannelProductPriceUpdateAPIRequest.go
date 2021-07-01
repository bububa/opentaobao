package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallSupplychainChannelProductPriceUpdateAPIRequest
渠道价格更新接口 API请求
tmall.supplychain.channel.product.price.update

更新渠道产品价格 */
type TmallSupplychainChannelProductPriceUpdateAPIRequest struct {
	model.Params
	// 币种，非必填，仅支持当商品记为外币价格时使用
	_currencyType string
	// 产品数字ID
	_productId int64
	// 1.指导价(默认) 2.区域价
	_priceType int64
	// 属于这个sku的商品的价格 取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。
	_skuPrice string
	// 产品价格，必填
	_productPrice string
	// SKU ID
	_skuId int64
	// 渠道编码
	_channelCode int64
}

// NewTmallSupplychainChannelProductPriceUpdateRequest 初始化TmallSupplychainChannelProductPriceUpdateAPIRequest对象
func NewTmallSupplychainChannelProductPriceUpdateRequest() *TmallSupplychainChannelProductPriceUpdateAPIRequest {
	return &TmallSupplychainChannelProductPriceUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductPriceUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.supplychain.channel.product.price.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductPriceUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CurrencyType Setter
// 币种，非必填，仅支持当商品记为外币价格时使用
func (r *TmallSupplychainChannelProductPriceUpdateAPIRequest) SetCurrencyType(_currencyType string) error {
	r._currencyType = _currencyType
	r.Set("currency_type", _currencyType)
	return nil
}

// Get CurrencyType Getter
func (r TmallSupplychainChannelProductPriceUpdateAPIRequest) GetCurrencyType() string {
	return r._currencyType
}

// Set is ProductId Setter
// 产品数字ID
func (r *TmallSupplychainChannelProductPriceUpdateAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// Get ProductId Getter
func (r TmallSupplychainChannelProductPriceUpdateAPIRequest) GetProductId() int64 {
	return r._productId
}

// Set is PriceType Setter
// 1.指导价(默认) 2.区域价
func (r *TmallSupplychainChannelProductPriceUpdateAPIRequest) SetPriceType(_priceType int64) error {
	r._priceType = _priceType
	r.Set("price_type", _priceType)
	return nil
}

// Get PriceType Getter
func (r TmallSupplychainChannelProductPriceUpdateAPIRequest) GetPriceType() int64 {
	return r._priceType
}

// Set is SkuPrice Setter
// 属于这个sku的商品的价格 取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。
func (r *TmallSupplychainChannelProductPriceUpdateAPIRequest) SetSkuPrice(_skuPrice string) error {
	r._skuPrice = _skuPrice
	r.Set("sku_price", _skuPrice)
	return nil
}

// Get SkuPrice Getter
func (r TmallSupplychainChannelProductPriceUpdateAPIRequest) GetSkuPrice() string {
	return r._skuPrice
}

// Set is ProductPrice Setter
// 产品价格，必填
func (r *TmallSupplychainChannelProductPriceUpdateAPIRequest) SetProductPrice(_productPrice string) error {
	r._productPrice = _productPrice
	r.Set("product_price", _productPrice)
	return nil
}

// Get ProductPrice Getter
func (r TmallSupplychainChannelProductPriceUpdateAPIRequest) GetProductPrice() string {
	return r._productPrice
}

// Set is SkuId Setter
// SKU ID
func (r *TmallSupplychainChannelProductPriceUpdateAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// Get SkuId Getter
func (r TmallSupplychainChannelProductPriceUpdateAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// Set is ChannelCode Setter
// 渠道编码
func (r *TmallSupplychainChannelProductPriceUpdateAPIRequest) SetChannelCode(_channelCode int64) error {
	r._channelCode = _channelCode
	r.Set("channel_code", _channelCode)
	return nil
}

// Get ChannelCode Getter
func (r TmallSupplychainChannelProductPriceUpdateAPIRequest) GetChannelCode() int64 {
	return r._channelCode
}
