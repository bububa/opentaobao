package wdk

// ExpirePeriodInfo 
/* model for simplify = false
type ExpirePeriodInfo struct {

    // 优惠具体折扣，1到1000
    
    PromotionValue   int64 `json:"promotion_value,omitempty"`
    

    // 0到24，小时数值
    
    EndHour   int64 `json:"end_hour,omitempty"`
    

    // 0到24，小时数值
    
    StartHour   int64 `json:"start_hour,omitempty"`
    

    // 外部商家id
    
    OutId   int64 `json:"out_id,omitempty"`
    

}
*/

// ExpirePeriodInfo 
type ExpirePeriodInfo struct {

    // 优惠具体折扣，1到1000
    PromotionValue   int64 `json:"promotion_value,omitempty"`

    // 0到24，小时数值
    EndHour   int64 `json:"end_hour,omitempty"`

    // 0到24，小时数值
    StartHour   int64 `json:"start_hour,omitempty"`

    // 外部商家id
    OutId   int64 `json:"out_id,omitempty"`

}
