package opentrade

// MarkUserInfo 
type MarkUserInfo struct {

    // 用户openId
    
    UserOpenId   string `json:"user_open_id,omitempty" xml:"user_open_id,omitempty"`
    

    // 商品ID
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 商品SKU ID
    
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    

    // 专属下单商品数量
    
    Quality   int64 `json:"quality,omitempty" xml:"quality,omitempty"`
    

    // 用户状态
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

}
