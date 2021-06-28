package wdk

// ReverseResult 
/* model for simplify = false
type ReverseResult struct {

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // 接口返回model
    
    Model  *struct {
        ApplyReverseResponse  *ApplyReverseResponse `json:"apply_reverse_response,omitempty"`
    } `json:"model,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// ReverseResult 
type ReverseResult struct {

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 接口返回model
    Model   *ApplyReverseResponse `json:"model,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}
