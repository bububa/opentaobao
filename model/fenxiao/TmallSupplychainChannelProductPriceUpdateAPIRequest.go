package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallsupplychainchannelproductpriceupdateAPIRequest 渠道价格更新接口 API请求
// tmall.supplychain.channel.product.price.update
//
// 更新渠道产品价格
type TmallsupplychainchannelproductpriceupdateAPIRequest struct {
	model.Params
	// 币种，非必填，仅支持当商品记为外币价格时使用
	_currencyType string
	// 属于这个sku的商品的价格 取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。
	_skuPrice string
	// 产品价格，必填
	_productPrice string
	// 产品数字ID
	_productId int64
	// 1.指导价(默认) 2.区域价
	_priceType int64
	// SKU ID
	_skuId int64
	// 渠道编码
	_channelCode int64
}

// NewTmallsupplychainchannelproductpriceupdateRequest 初始化TmallsupplychainchannelproductpriceupdateAPIRequest对象
func NewTmallsupplychainchannelproductpriceupdateRequest() *TmallsupplychainchannelproductpriceupdateAPIRequest {
	return &TmallsupplychainchannelproductpriceupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallsupplychainchannelproductpriceupdateAPIRequest) GetApiMethodName() string {
	return "tmall.supplychain.channel.product.price.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallsupplychainchannelproductpriceupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallsupplychainchannelproductpriceupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCurrencyType is CurrencyType Setter
// 币种，非必填，仅支持当商品记为外币价格时使用
func (r *TmallsupplychainchannelproductpriceupdateAPIRequest) SetCurrencyType(_currencyType string) error {
	r._currencyType = _currencyType
	r.Set("currency_type", _currencyType)
	return nil
}

// GetCurrencyType CurrencyType Getter
func (r TmallsupplychainchannelproductpriceupdateAPIRequest) GetCurrencyType() string {
	return r._currencyType
}

// SetSkuPrice is SkuPrice Setter
// 属于这个sku的商品的价格 取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。
func (r *TmallsupplychainchannelproductpriceupdateAPIRequest) SetSkuPrice(_skuPrice string) error {
	r._skuPrice = _skuPrice
	r.Set("sku_price", _skuPrice)
	return nil
}

// GetSkuPrice SkuPrice Getter
func (r TmallsupplychainchannelproductpriceupdateAPIRequest) GetSkuPrice() string {
	return r._skuPrice
}

// SetProductPrice is ProductPrice Setter
// 产品价格，必填
func (r *TmallsupplychainchannelproductpriceupdateAPIRequest) SetProductPrice(_productPrice string) error {
	r._productPrice = _productPrice
	r.Set("product_price", _productPrice)
	return nil
}

// GetProductPrice ProductPrice Getter
func (r TmallsupplychainchannelproductpriceupdateAPIRequest) GetProductPrice() string {
	return r._productPrice
}

// SetProductId is ProductId Setter
// 产品数字ID
func (r *TmallsupplychainchannelproductpriceupdateAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallsupplychainchannelproductpriceupdateAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetPriceType is PriceType Setter
// 1.指导价(默认) 2.区域价
func (r *TmallsupplychainchannelproductpriceupdateAPIRequest) SetPriceType(_priceType int64) error {
	r._priceType = _priceType
	r.Set("price_type", _priceType)
	return nil
}

// GetPriceType PriceType Getter
func (r TmallsupplychainchannelproductpriceupdateAPIRequest) GetPriceType() int64 {
	return r._priceType
}

// SetSkuId is SkuId Setter
// SKU ID
func (r *TmallsupplychainchannelproductpriceupdateAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TmallsupplychainchannelproductpriceupdateAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// SetChannelCode is ChannelCode Setter
// 渠道编码
func (r *TmallsupplychainchannelproductpriceupdateAPIRequest) SetChannelCode(_channelCode int64) error {
	r._channelCode = _channelCode
	r.Set("channel_code", _channelCode)
	return nil
}

// GetChannelCode ChannelCode Getter
func (r TmallsupplychainchannelproductpriceupdateAPIRequest) GetChannelCode() int64 {
	return r._channelCode
}
