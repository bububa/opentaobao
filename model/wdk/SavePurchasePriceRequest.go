package wdk

// SavePurchasePriceRequest 
/* model for simplify = false
type SavePurchasePriceRequest struct {

    // 2-经销，3-代销，6-寄售，默认为【6-寄售】
    
    MarketingType   int64 `json:"marketing_type,omitempty"`
    

    // 门店ID
    
    OuCode   string `json:"ou_code,omitempty"`
    

    // 商品编码
    
    SkuCode   string `json:"sku_code,omitempty"`
    

    // 含税采购价格，单位【分】
    
    PurchasePriceWithTax   int64 `json:"purchase_price_with_tax,omitempty"`
    

    // 1-基准价格，3-区间价格
    
    PriceType   int64 `json:"price_type,omitempty"`
    

    // 区间价格生效时间
    
    EffectiveStartTime   int64 `json:"effective_start_time,omitempty"`
    

    // 区间价格失效时间
    
    EffectiveEndTime   int64 `json:"effective_end_time,omitempty"`
    

    // 幂等ID
    
    OutId   string `json:"out_id,omitempty"`
    

}
*/

// SavePurchasePriceRequest 
type SavePurchasePriceRequest struct {

    // 2-经销，3-代销，6-寄售，默认为【6-寄售】
    MarketingType   int64 `json:"marketing_type,omitempty"`

    // 门店ID
    OuCode   string `json:"ou_code,omitempty"`

    // 商品编码
    SkuCode   string `json:"sku_code,omitempty"`

    // 含税采购价格，单位【分】
    PurchasePriceWithTax   int64 `json:"purchase_price_with_tax,omitempty"`

    // 1-基准价格，3-区间价格
    PriceType   int64 `json:"price_type,omitempty"`

    // 区间价格生效时间
    EffectiveStartTime   int64 `json:"effective_start_time,omitempty"`

    // 区间价格失效时间
    EffectiveEndTime   int64 `json:"effective_end_time,omitempty"`

    // 幂等ID
    OutId   string `json:"out_id,omitempty"`

}
