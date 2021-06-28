package ticket

// DateInventory 
/* model for simplify = false
type DateInventory struct {

    // 日期
    
    Date   string `json:"date,omitempty"`
    

    // 价格，精确到分
    
    Price   int64 `json:"price,omitempty"`
    

    // 库存
    
    Stock   int64 `json:"stock,omitempty"`
    

    // 日期级别自定义商家编码，为该sku下每一天都设置一个自定义商家编码。如果outSkuDateId为空，则该天的商家自定义编码将以outSkuId为准
    
    OutSkuDateId   string `json:"out_sku_date_id,omitempty"`
    

}
*/

// DateInventory 
type DateInventory struct {

    // 日期
    Date   string `json:"date,omitempty"`

    // 价格，精确到分
    Price   int64 `json:"price,omitempty"`

    // 库存
    Stock   int64 `json:"stock,omitempty"`

    // 日期级别自定义商家编码，为该sku下每一天都设置一个自定义商家编码。如果outSkuDateId为空，则该天的商家自定义编码将以outSkuId为准
    OutSkuDateId   string `json:"out_sku_date_id,omitempty"`

}
