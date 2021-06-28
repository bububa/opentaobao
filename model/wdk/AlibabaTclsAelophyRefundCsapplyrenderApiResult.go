package wdk

// AlibabaTclsAelophyRefundCsapplyrenderApiResult 
/* model for simplify = false
type AlibabaTclsAelophyRefundCsapplyrenderApiResult struct {

    // 回调返回编码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 回调是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 回调错误信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 数据
    
    Model  *struct {
        RefundCsApplyRenderResponseDto  *RefundCsApplyRenderResponseDto `json:"refund_cs_apply_render_response_dto,omitempty"`
    } `json:"model,omitempty"`
    

}
*/

// AlibabaTclsAelophyRefundCsapplyrenderApiResult 
type AlibabaTclsAelophyRefundCsapplyrenderApiResult struct {

    // 回调返回编码
    ErrCode   string `json:"err_code,omitempty"`

    // 回调是否成功
    Success   bool `json:"success,omitempty"`

    // 回调错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 数据
    Model   *RefundCsApplyRenderResponseDto `json:"model,omitempty"`

}
