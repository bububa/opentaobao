package wdk

// T 
/* model for simplify = false
type T struct {

    // 商家code
    
    MerchantCode   string `json:"merchant_code,omitempty"`
    

    // 商品skucode
    
    SkuCode   string `json:"sku_code,omitempty"`
    

    // 商品itemId
    
    ItemId   int64 `json:"item_id,omitempty"`
    

    // 门店id
    
    ShopId   string `json:"shop_id,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

}
*/

// T 
type T struct {

    // 商家code
    MerchantCode   string `json:"merchant_code,omitempty"`

    // 商品skucode
    SkuCode   string `json:"sku_code,omitempty"`

    // 商品itemId
    ItemId   int64 `json:"item_id,omitempty"`

    // 门店id
    ShopId   string `json:"shop_id,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

}
