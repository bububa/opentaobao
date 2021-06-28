package servicecenter

// CarItemInfoDto 
/* model for simplify = false
type CarItemInfoDto struct {

    // 品牌
    
    Brand   string `json:"brand,omitempty"`
    

    // 商品id
    
    ItemId   int64 `json:"item_id,omitempty"`
    

    // 车系
    
    Line   string `json:"line,omitempty"`
    

    // 型号
    
    Model   string `json:"model,omitempty"`
    

    // 不会返回
    
    SellerId   int64 `json:"seller_id,omitempty"`
    

    // 不会返回
    
    SellerNick   string `json:"seller_nick,omitempty"`
    

    // skuId不会返回
    
    SkuId   int64 `json:"sku_id,omitempty"`
    

    // 年款
    
    Year   string `json:"year,omitempty"`
    

}
*/

// CarItemInfoDto 
type CarItemInfoDto struct {

    // 品牌
    Brand   string `json:"brand,omitempty"`

    // 商品id
    ItemId   int64 `json:"item_id,omitempty"`

    // 车系
    Line   string `json:"line,omitempty"`

    // 型号
    Model   string `json:"model,omitempty"`

    // 不会返回
    SellerId   int64 `json:"seller_id,omitempty"`

    // 不会返回
    SellerNick   string `json:"seller_nick,omitempty"`

    // skuId不会返回
    SkuId   int64 `json:"sku_id,omitempty"`

    // 年款
    Year   string `json:"year,omitempty"`

}
