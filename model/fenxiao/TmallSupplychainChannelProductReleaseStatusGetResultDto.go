package fenxiao

// TmallSupplychainChannelProductReleaseStatusGetResultDto 
/* model for simplify = false
type TmallSupplychainChannelProductReleaseStatusGetResultDto struct {

    // 执行结果
    
    Success   bool `json:"success,omitempty"`
    

    // 查询结果
    
    Module  *struct {
        TopProductStatusResult  *TopProductStatusResult `json:"top_product_status_result,omitempty"`
    } `json:"module,omitempty"`
    

    // 错误码
    
    ErrorMessage   string `json:"error_message,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

}
*/

// TmallSupplychainChannelProductReleaseStatusGetResultDto 
type TmallSupplychainChannelProductReleaseStatusGetResultDto struct {

    // 执行结果
    Success   bool `json:"success,omitempty"`

    // 查询结果
    Module   *TopProductStatusResult `json:"module,omitempty"`

    // 错误码
    ErrorMessage   string `json:"error_message,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

}
