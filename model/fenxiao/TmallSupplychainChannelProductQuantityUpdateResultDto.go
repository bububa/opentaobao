package fenxiao

// TmallSupplychainChannelProductQuantityUpdateResultDto 
/* model for simplify = false
type TmallSupplychainChannelProductQuantityUpdateResultDto struct {

    // 执行结果
    
    Success   bool `json:"success,omitempty"`
    

    // 更新内容
    
    Module  *struct {
        TopProductQuantityResult  *TopProductQuantityResult `json:"top_product_quantity_result,omitempty"`
    } `json:"module,omitempty"`
    

    // 错误码
    
    ErrorMessage   string `json:"error_message,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

}
*/

// TmallSupplychainChannelProductQuantityUpdateResultDto 
type TmallSupplychainChannelProductQuantityUpdateResultDto struct {

    // 执行结果
    Success   bool `json:"success,omitempty"`

    // 更新内容
    Module   *TopProductQuantityResult `json:"module,omitempty"`

    // 错误码
    ErrorMessage   string `json:"error_message,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

}
