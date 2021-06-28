package wdk

// AlibabaWdkReverseRefundResult 
/* model for simplify = false
type AlibabaWdkReverseRefundResult struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // 退款单ID
    
    Model   int64 `json:"model,omitempty"`
    

}
*/

// AlibabaWdkReverseRefundResult 
type AlibabaWdkReverseRefundResult struct {

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 退款单ID
    Model   int64 `json:"model,omitempty"`

}
