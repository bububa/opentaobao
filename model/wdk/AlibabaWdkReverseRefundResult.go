package wdk

// AlibabaWdkReverseRefundResult 
type AlibabaWdkReverseRefundResult struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

    // 退款单ID
    
    Model   int64 `json:"model,omitempty" xml:"model,omitempty"`
    

}
