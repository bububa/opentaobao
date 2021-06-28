package simba

// ExtraAttributes 
/* model for simplify = false
type ExtraAttributes struct {

    // 商品在淘宝的发布时间
    
    PublishTime   string `json:"publish_time,omitempty"`
    

    // 库存数量
    
    Quantity   int64 `json:"quantity,omitempty"`
    

    // 商品的累积销量
    
    SalesCount   int64 `json:"sales_count,omitempty"`
    

}
*/

// ExtraAttributes 
type ExtraAttributes struct {

    // 商品在淘宝的发布时间
    PublishTime   string `json:"publish_time,omitempty"`

    // 库存数量
    Quantity   int64 `json:"quantity,omitempty"`

    // 商品的累积销量
    SalesCount   int64 `json:"sales_count,omitempty"`

}
