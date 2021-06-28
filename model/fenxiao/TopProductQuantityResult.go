package fenxiao

// TopProductQuantityResult 
/* model for simplify = false
type TopProductQuantityResult struct {

    // 产品数字ID
    
    ProductId   int64 `json:"product_id,omitempty"`
    

    // SKU ID
    
    SkuId   int64 `json:"sku_id,omitempty"`
    

    // 更新后库存数量
    
    Result   string `json:"result,omitempty"`
    

    // 更新时间
    
    ModifiedTime   string `json:"modified_time,omitempty"`
    

}
*/

// TopProductQuantityResult 
type TopProductQuantityResult struct {

    // 产品数字ID
    ProductId   int64 `json:"product_id,omitempty"`

    // SKU ID
    SkuId   int64 `json:"sku_id,omitempty"`

    // 更新后库存数量
    Result   string `json:"result,omitempty"`

    // 更新时间
    ModifiedTime   string `json:"modified_time,omitempty"`

}
