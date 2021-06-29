package aesolution

// SynchronizeSkuRequestDto 
type SynchronizeSkuRequestDto struct {

    // sku code
    
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
    

    // inventory
    
    Inventory   int64 `json:"inventory,omitempty" xml:"inventory,omitempty"`
    

    // price of an sku
    
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    

    // discount_price of an sku. If not set, the discount_price will be erased.
    
    DiscountPrice   string `json:"discount_price,omitempty" xml:"discount_price,omitempty"`
    

}
