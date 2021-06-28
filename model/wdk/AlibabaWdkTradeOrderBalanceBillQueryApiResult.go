package wdk

// AlibabaWdkTradeOrderBalanceBillQueryApiResult 
/* model for simplify = false
type AlibabaWdkTradeOrderBalanceBillQueryApiResult struct {

    // model
    
    Model  *struct {
        OrderBalanceBillResponseDo  *OrderBalanceBillResponseDo `json:"order_balance_bill_response_do,omitempty"`
    } `json:"model,omitempty"`
    

    // 错误码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 成功失败
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlibabaWdkTradeOrderBalanceBillQueryApiResult 
type AlibabaWdkTradeOrderBalanceBillQueryApiResult struct {

    // model
    Model   *OrderBalanceBillResponseDo `json:"model,omitempty"`

    // 错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 成功失败
    Success   bool `json:"success,omitempty"`

}
