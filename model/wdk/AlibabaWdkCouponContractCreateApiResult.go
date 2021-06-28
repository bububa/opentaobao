package wdk

// AlibabaWdkCouponContractCreateApiResult 
/* model for simplify = false
type AlibabaWdkCouponContractCreateApiResult struct {

    // 合同ID
    
    Model   int64 `json:"model,omitempty"`
    

    // 错误码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 错误码说明
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 是否调用成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlibabaWdkCouponContractCreateApiResult 
type AlibabaWdkCouponContractCreateApiResult struct {

    // 合同ID
    Model   int64 `json:"model,omitempty"`

    // 错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 错误码说明
    ErrMsg   string `json:"err_msg,omitempty"`

    // 是否调用成功
    Success   bool `json:"success,omitempty"`

}
