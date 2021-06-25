package fenxiao

// TopProductPriceResult 
type TopProductPriceResult struct {

    // 产品ID
    ProductId   int64 `json:"product_id,omitempty"`

    // SKU ID
    SkuId   int64 `json:"sku_id,omitempty"`

    // 渠道编码
    ChannelCode   int64 `json:"channel_code,omitempty"`

    // 价格类型，区域价、指导价
    PriceType   int64 `json:"price_type,omitempty"`

    // 价格
    Price   string `json:"price,omitempty"`

    // 币种
    CurrencyType   string `json:"currency_type,omitempty"`

}
