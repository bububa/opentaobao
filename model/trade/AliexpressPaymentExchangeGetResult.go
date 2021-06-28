package trade

// AliexpressPaymentExchangeGetResult 
/* model for simplify = false
type AliexpressPaymentExchangeGetResult struct {

    // 是否不成功
    
    NotSuccess   bool `json:"not_success,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 返汇率相关数据
    
    Module  *struct {
        AliexpressPaymentExchangeGetModule  *AliexpressPaymentExchangeGetModule `json:"aliexpress_payment_exchange_get_module,omitempty"`
    } `json:"module,omitempty"`
    

    // 错误信息
    
    ErrorCode  *struct {
        ErrorCode  *ErrorCode `json:"error_code,omitempty"`
    } `json:"error_code,omitempty"`
    

    // 是否重复重复
    
    Repeated   bool `json:"repeated,omitempty"`
    

    // 是否重试
    
    Retry   bool `json:"retry,omitempty"`
    

}
*/

// AliexpressPaymentExchangeGetResult 
type AliexpressPaymentExchangeGetResult struct {

    // 是否不成功
    NotSuccess   bool `json:"not_success,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 返汇率相关数据
    Module   *AliexpressPaymentExchangeGetModule `json:"module,omitempty"`

    // 错误信息
    ErrorCode   *ErrorCode `json:"error_code,omitempty"`

    // 是否重复重复
    Repeated   bool `json:"repeated,omitempty"`

    // 是否重试
    Retry   bool `json:"retry,omitempty"`

}
