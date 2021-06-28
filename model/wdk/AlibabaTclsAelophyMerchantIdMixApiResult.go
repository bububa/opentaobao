package wdk

// AlibabaTclsAelophyMerchantIdMixApiResult 
/* model for simplify = false
type AlibabaTclsAelophyMerchantIdMixApiResult struct {

    // 获取mixId成功
    
    Success   bool `json:"success,omitempty"`
    

    // 错误码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 返回混淆id
    
    Model   string `json:"model,omitempty"`
    

}
*/

// AlibabaTclsAelophyMerchantIdMixApiResult 
type AlibabaTclsAelophyMerchantIdMixApiResult struct {

    // 获取mixId成功
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 返回混淆id
    Model   string `json:"model,omitempty"`

}
