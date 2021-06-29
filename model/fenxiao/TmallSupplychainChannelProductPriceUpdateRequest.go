package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道价格更新接口 API请求
tmall.supplychain.channel.product.price.update

更新渠道产品价格
*/
type TmallSupplychainChannelProductPriceUpdateRequest struct {
    model.Params
    // 币种，非必填，仅支持当商品记为外币价格时使用
    currencyType   string
    // 产品数字ID
    productId   int64
    // 1.指导价(默认) 2.区域价
    priceType   int64
    // 属于这个sku的商品的价格 取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。
    skuPrice   string
    // 产品价格，必填
    productPrice   string
    // SKU ID
    skuId   int64
    // 渠道编码
    channelCode   int64
}

// 初始化TmallSupplychainChannelProductPriceUpdateRequest对象
func NewTmallSupplychainChannelProductPriceUpdateRequest() *TmallSupplychainChannelProductPriceUpdateRequest{
    return &TmallSupplychainChannelProductPriceUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductPriceUpdateRequest) GetApiMethodName() string {
    return "tmall.supplychain.channel.product.price.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductPriceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CurrencyType Setter
// 币种，非必填，仅支持当商品记为外币价格时使用
func (r *TmallSupplychainChannelProductPriceUpdateRequest) SetCurrencyType(currencyType string) error {
    r.currencyType = currencyType
    r.Set("currency_type", currencyType)
    return nil
}

// CurrencyType Getter
func (r TmallSupplychainChannelProductPriceUpdateRequest) GetCurrencyType() string {
    return r.currencyType
}
// ProductId Setter
// 产品数字ID
func (r *TmallSupplychainChannelProductPriceUpdateRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TmallSupplychainChannelProductPriceUpdateRequest) GetProductId() int64 {
    return r.productId
}
// PriceType Setter
// 1.指导价(默认) 2.区域价
func (r *TmallSupplychainChannelProductPriceUpdateRequest) SetPriceType(priceType int64) error {
    r.priceType = priceType
    r.Set("price_type", priceType)
    return nil
}

// PriceType Getter
func (r TmallSupplychainChannelProductPriceUpdateRequest) GetPriceType() int64 {
    return r.priceType
}
// SkuPrice Setter
// 属于这个sku的商品的价格 取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。
func (r *TmallSupplychainChannelProductPriceUpdateRequest) SetSkuPrice(skuPrice string) error {
    r.skuPrice = skuPrice
    r.Set("sku_price", skuPrice)
    return nil
}

// SkuPrice Getter
func (r TmallSupplychainChannelProductPriceUpdateRequest) GetSkuPrice() string {
    return r.skuPrice
}
// ProductPrice Setter
// 产品价格，必填
func (r *TmallSupplychainChannelProductPriceUpdateRequest) SetProductPrice(productPrice string) error {
    r.productPrice = productPrice
    r.Set("product_price", productPrice)
    return nil
}

// ProductPrice Getter
func (r TmallSupplychainChannelProductPriceUpdateRequest) GetProductPrice() string {
    return r.productPrice
}
// SkuId Setter
// SKU ID
func (r *TmallSupplychainChannelProductPriceUpdateRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

// SkuId Getter
func (r TmallSupplychainChannelProductPriceUpdateRequest) GetSkuId() int64 {
    return r.skuId
}
// ChannelCode Setter
// 渠道编码
func (r *TmallSupplychainChannelProductPriceUpdateRequest) SetChannelCode(channelCode int64) error {
    r.channelCode = channelCode
    r.Set("channel_code", channelCode)
    return nil
}

// ChannelCode Getter
func (r TmallSupplychainChannelProductPriceUpdateRequest) GetChannelCode() int64 {
    return r.channelCode
}
