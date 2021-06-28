package fenxiao

// TmallSupplychainChannelProductDownshelfResultDto 
/* model for simplify = false
type TmallSupplychainChannelProductDownshelfResultDto struct {

    // 执行结果
    
    Success   bool `json:"success,omitempty"`
    

    // 下架结果
    
    Module   bool `json:"module,omitempty"`
    

    // 错误码
    
    ErrorMessage   string `json:"error_message,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

}
*/

// TmallSupplychainChannelProductDownshelfResultDto 
type TmallSupplychainChannelProductDownshelfResultDto struct {

    // 执行结果
    Success   bool `json:"success,omitempty"`

    // 下架结果
    Module   bool `json:"module,omitempty"`

    // 错误码
    ErrorMessage   string `json:"error_message,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

}
