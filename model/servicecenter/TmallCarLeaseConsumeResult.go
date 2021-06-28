package servicecenter

// TmallCarLeaseConsumeResult 
type TmallCarLeaseConsumeResult struct {

    // 耗时
    
    CostTime   int64 `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
    

    // 错误吗
    
    ErrorCode   int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 错误提示
    
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    

    // 当前时间
    
    GmtCurrentTime   int64 `json:"gmt_current_time,omitempty" xml:"gmt_current_time,omitempty"`
    

    // 无需关心了
    
    Object   bool `json:"object,omitempty" xml:"object,omitempty"`
    

    // 成功与否
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
