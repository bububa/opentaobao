package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道价格查询接口 API请求
tmall.supplychain.channel.product.price.get

渠道价格查询接口
*/
type TmallSupplychainChannelProductPriceGetRequest struct {
    model.Params
    // 产品ID
    _productId   int64
    // 区域价、指导价
    _priceType   int64
    // SKU ID
    _skuId   int64
    // 渠道ID（台湾市场/供销平台/大农业…..）
    _channelCode   int64
}

// 初始化TmallSupplychainChannelProductPriceGetRequest对象
func NewTmallSupplychainChannelProductPriceGetRequest() *TmallSupplychainChannelProductPriceGetRequest{
    return &TmallSupplychainChannelProductPriceGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductPriceGetRequest) GetApiMethodName() string {
    return "tmall.supplychain.channel.product.price.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductPriceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品ID
func (r *TmallSupplychainChannelProductPriceGetRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TmallSupplychainChannelProductPriceGetRequest) GetProductId() int64 {
    return r._productId
}
// PriceType Setter
// 区域价、指导价
func (r *TmallSupplychainChannelProductPriceGetRequest) SetPriceType(_priceType int64) error {
    r._priceType = _priceType
    r.Set("price_type", _priceType)
    return nil
}

// PriceType Getter
func (r TmallSupplychainChannelProductPriceGetRequest) GetPriceType() int64 {
    return r._priceType
}
// SkuId Setter
// SKU ID
func (r *TmallSupplychainChannelProductPriceGetRequest) SetSkuId(_skuId int64) error {
    r._skuId = _skuId
    r.Set("sku_id", _skuId)
    return nil
}

// SkuId Getter
func (r TmallSupplychainChannelProductPriceGetRequest) GetSkuId() int64 {
    return r._skuId
}
// ChannelCode Setter
// 渠道ID（台湾市场/供销平台/大农业…..）
func (r *TmallSupplychainChannelProductPriceGetRequest) SetChannelCode(_channelCode int64) error {
    r._channelCode = _channelCode
    r.Set("channel_code", _channelCode)
    return nil
}

// ChannelCode Getter
func (r TmallSupplychainChannelProductPriceGetRequest) GetChannelCode() int64 {
    return r._channelCode
}
