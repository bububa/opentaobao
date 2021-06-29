package ascpchannel

// TopDistributorPriceResult 
type TopDistributorPriceResult struct {

    // 产品ID
    
    ProductId   string `json:"product_id,omitempty" xml:"product_id,omitempty"`
    

    // 产品价格详情
    
    ProductPrice   *TopChannelPriceDetail `json:"product_price,omitempty" xml:"product_price,omitempty"`
    

    // sku价格详情
    
    SkuPrices   []TopChannelSkuPrice `json:"sku_prices,omitempty" xml:"sku_prices,omitempty"`
    

}
