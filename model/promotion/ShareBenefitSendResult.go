package promotion

// ShareBenefitSendResult 
/* model for simplify = false
type ShareBenefitSendResult struct {

    // 错误信息
    
    ErrorMessage   string `json:"error_message,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 发放结果是否正常
    
    Success   bool `json:"success,omitempty"`
    

    // 发放结果
    
    ResultMap   string `json:"result_map,omitempty"`
    

}
*/

// ShareBenefitSendResult 
type ShareBenefitSendResult struct {

    // 错误信息
    ErrorMessage   string `json:"error_message,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 发放结果是否正常
    Success   bool `json:"success,omitempty"`

    // 发放结果
    ResultMap   string `json:"result_map,omitempty"`

}
