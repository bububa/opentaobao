package drug

// OrderPromotionDto 
/* model for simplify = false
type OrderPromotionDto struct {

    // 优惠类型
    
    Type   int64 `json:"type,omitempty"`
    

    // 优惠名称
    
    Name   string `json:"name,omitempty"`
    

    // 优惠金额
    
    Amount   int64 `json:"amount,omitempty"`
    

}
*/

// OrderPromotionDto 
type OrderPromotionDto struct {

    // 优惠类型
    Type   int64 `json:"type,omitempty"`

    // 优惠名称
    Name   string `json:"name,omitempty"`

    // 优惠金额
    Amount   int64 `json:"amount,omitempty"`

}
