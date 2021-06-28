package refund

// RefundMappingResult 
/* model for simplify = false
type RefundMappingResult struct {

    // 退款ID
    
    RefundId   string `json:"refund_id,omitempty"`
    

    // 结果信息
    
    Message   string `json:"message,omitempty"`
    

    // 是否成功
    
    Succ   bool `json:"succ,omitempty"`
    

}
*/

// RefundMappingResult 
type RefundMappingResult struct {

    // 退款ID
    RefundId   string `json:"refund_id,omitempty"`

    // 结果信息
    Message   string `json:"message,omitempty"`

    // 是否成功
    Succ   bool `json:"succ,omitempty"`

}
