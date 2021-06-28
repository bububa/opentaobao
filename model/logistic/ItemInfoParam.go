package logistic

// ItemInfoParam 
/* model for simplify = false
type ItemInfoParam struct {

    // 餐品实际价格
    
    ItemActualPrice   int64 `json:"item_actual_price,omitempty"`
    

    // 餐品id
    
    ItemId   string `json:"item_id,omitempty"`
    

    // 餐品数量
    
    ItemQuantity   int64 `json:"item_quantity,omitempty"`
    

    // 餐品名
    
    ItemName   string `json:"item_name,omitempty"`
    

    // 餐品重量
    
    ItemWeight   int64 `json:"item_weight,omitempty"`
    

    // 餐品类目
    
    ItemCategory   string `json:"item_category,omitempty"`
    

}
*/

// ItemInfoParam 
type ItemInfoParam struct {

    // 餐品实际价格
    ItemActualPrice   int64 `json:"item_actual_price,omitempty"`

    // 餐品id
    ItemId   string `json:"item_id,omitempty"`

    // 餐品数量
    ItemQuantity   int64 `json:"item_quantity,omitempty"`

    // 餐品名
    ItemName   string `json:"item_name,omitempty"`

    // 餐品重量
    ItemWeight   int64 `json:"item_weight,omitempty"`

    // 餐品类目
    ItemCategory   string `json:"item_category,omitempty"`

}
