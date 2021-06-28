package wdk

// AlibabaWdkChannelOrderStatusUpdateApiResult 
/* model for simplify = false
type AlibabaWdkChannelOrderStatusUpdateApiResult struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 返回内容
    
    Model  *struct {
        OrderOperateResult  *OrderOperateResult `json:"order_operate_result,omitempty"`
    } `json:"model,omitempty"`
    

    // 错误编码
    
    ErrCode   string `json:"err_code,omitempty"`
    

}
*/

// AlibabaWdkChannelOrderStatusUpdateApiResult 
type AlibabaWdkChannelOrderStatusUpdateApiResult struct {

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 返回内容
    Model   *OrderOperateResult `json:"model,omitempty"`

    // 错误编码
    ErrCode   string `json:"err_code,omitempty"`

}
