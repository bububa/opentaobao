package wdk

// ItemStairSku 
type ItemStairSku struct {

    // 商品skucode
    
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
    

    // 单位为分
    
    ExchangePrice   int64 `json:"exchange_price,omitempty" xml:"exchange_price,omitempty"`
    

    // 换购商品总数限制
    
    ExchangeTotalLimit   int64 `json:"exchange_total_limit,omitempty" xml:"exchange_total_limit,omitempty"`
    

}
