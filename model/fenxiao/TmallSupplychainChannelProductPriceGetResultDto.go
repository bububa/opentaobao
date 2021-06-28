package fenxiao

// TmallSupplychainChannelProductPriceGetResultDto 
/* model for simplify = false
type TmallSupplychainChannelProductPriceGetResultDto struct {

    // 执行结果
    
    Success   bool `json:"success,omitempty"`
    

    // 询价结果
    
    Module  *struct {
        TopProductPriceResult  *TopProductPriceResult `json:"top_product_price_result,omitempty"`
    } `json:"module,omitempty"`
    

    // 错误码
    
    ErrorMessage   string `json:"error_message,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

}
*/

// TmallSupplychainChannelProductPriceGetResultDto 
type TmallSupplychainChannelProductPriceGetResultDto struct {

    // 执行结果
    Success   bool `json:"success,omitempty"`

    // 询价结果
    Module   *TopProductPriceResult `json:"module,omitempty"`

    // 错误码
    ErrorMessage   string `json:"error_message,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

}
