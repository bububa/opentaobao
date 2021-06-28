package wdk

// AlibabaWdkFinanceOrderBackflowApiResult 
/* model for simplify = false
type AlibabaWdkFinanceOrderBackflowApiResult struct {

    // 调用接口返回的具体结果
    
    Models  *struct {
        FinanceOrderDetailResponse  *FinanceOrderDetailResponse `json:"finance_order_detail_response,omitempty"`
    } `json:"models,omitempty"`
    

    // 调用接口返回错误码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 调用接口返回错误信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 调用接口返回成功失败
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlibabaWdkFinanceOrderBackflowApiResult 
type AlibabaWdkFinanceOrderBackflowApiResult struct {

    // 调用接口返回的具体结果
    Models   *FinanceOrderDetailResponse `json:"models,omitempty"`

    // 调用接口返回错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 调用接口返回错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 调用接口返回成功失败
    Success   bool `json:"success,omitempty"`

}
