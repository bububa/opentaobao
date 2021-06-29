package aesolution

// SingleSkuPriceByCountryDto 
type SingleSkuPriceByCountryDto struct {

    // sku_code, must existed in  multiple_sku_update_list
    
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
    

    // Value of price configuration. If the price of a specific country is set, it must be greater than or equal to 70% of the original sku price in multiple_sku_update_list
    
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    

    // Value of discount_price for each country
    
    DiscountPrice   string `json:"discount_price,omitempty" xml:"discount_price,omitempty"`
    

    // Deprecated. Please do not use.
    
    Value   string `json:"value,omitempty" xml:"value,omitempty"`
    

}
