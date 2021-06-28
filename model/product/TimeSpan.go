package product

// TimeSpan 
/* model for simplify = false
type TimeSpan struct {

    // 变更开始时间
    
    Start   string `json:"start,omitempty"`
    

    // 变更结束时间
    
    End   string `json:"end,omitempty"`
    

    // 变更时间段内的具体有变更的日期掩码
    
    ChangeDateMask   string `json:"change_date_mask,omitempty"`
    

}
*/

// TimeSpan 
type TimeSpan struct {

    // 变更开始时间
    Start   string `json:"start,omitempty"`

    // 变更结束时间
    End   string `json:"end,omitempty"`

    // 变更时间段内的具体有变更的日期掩码
    ChangeDateMask   string `json:"change_date_mask,omitempty"`

}
