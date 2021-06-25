package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道价格查询接口 APIRequest
tmall.supplychain.channel.product.price.get

渠道价格查询接口
*/
type TmallSupplychainChannelProductPriceGetRequest struct {
    model.Params

    // 产品ID
    productId   int64 

    // 区域价、指导价
    priceType   int64 

    // SKU ID
    skuId   int64 

    // 渠道ID（台湾市场/供销平台/大农业…..）
    channelCode   int64 

}

func NewTmallSupplychainChannelProductPriceGetRequest() *TmallSupplychainChannelProductPriceGetRequest{
    return &TmallSupplychainChannelProductPriceGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallSupplychainChannelProductPriceGetRequest) GetApiMethodName() string {
    return "tmall.supplychain.channel.product.price.get"
}

func (r TmallSupplychainChannelProductPriceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallSupplychainChannelProductPriceGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TmallSupplychainChannelProductPriceGetRequest) GetProductId() int64 {
    return r.productId
}

func (r *TmallSupplychainChannelProductPriceGetRequest) SetPriceType(priceType int64) error {
    r.priceType = priceType
    r.Set("price_type", priceType)
    return nil
}

func (r TmallSupplychainChannelProductPriceGetRequest) GetPriceType() int64 {
    return r.priceType
}

func (r *TmallSupplychainChannelProductPriceGetRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

func (r TmallSupplychainChannelProductPriceGetRequest) GetSkuId() int64 {
    return r.skuId
}

func (r *TmallSupplychainChannelProductPriceGetRequest) SetChannelCode(channelCode int64) error {
    r.channelCode = channelCode
    r.Set("channel_code", channelCode)
    return nil
}

func (r TmallSupplychainChannelProductPriceGetRequest) GetChannelCode() int64 {
    return r.channelCode
}

