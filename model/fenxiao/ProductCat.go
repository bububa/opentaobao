package fenxiao

// ProductCat 
/* model for simplify = false
type ProductCat struct {

    // 产品线名称
    
    Name   string `json:"name,omitempty"`
    

    // 产品数量
    
    ProductNum   int64 `json:"product_num,omitempty"`
    

    // 产品线ID
    
    Id   int64 `json:"id,omitempty"`
    

    // 最低零售价百分比
    
    RetailLowPercent   string `json:"retail_low_percent,omitempty"`
    

    // 最高零食价百分比
    
    RetailHighPercent   string `json:"retail_high_percent,omitempty"`
    

    // 代销采购价百分比
    
    CostPercentAgent   string `json:"cost_percent_agent,omitempty"`
    

    // 经销采购价百分比
    
    CostPercentDealer   string `json:"cost_percent_dealer,omitempty"`
    

}
*/

// ProductCat 
type ProductCat struct {

    // 产品线名称
    Name   string `json:"name,omitempty"`

    // 产品数量
    ProductNum   int64 `json:"product_num,omitempty"`

    // 产品线ID
    Id   int64 `json:"id,omitempty"`

    // 最低零售价百分比
    RetailLowPercent   string `json:"retail_low_percent,omitempty"`

    // 最高零食价百分比
    RetailHighPercent   string `json:"retail_high_percent,omitempty"`

    // 代销采购价百分比
    CostPercentAgent   string `json:"cost_percent_agent,omitempty"`

    // 经销采购价百分比
    CostPercentDealer   string `json:"cost_percent_dealer,omitempty"`

}
