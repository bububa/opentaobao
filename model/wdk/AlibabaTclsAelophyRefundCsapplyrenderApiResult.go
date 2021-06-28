package wdk

// AlibabaTclsAelophyRefundCsapplyrenderApiResult 
type AlibabaTclsAelophyRefundCsapplyrenderApiResult struct {

    // 回调返回编码
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // 回调是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 回调错误信息
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

    // 数据
    
    Model   *RefundCsApplyRenderResponseDto `json:"model,omitempty" xml:"model,omitempty"`
    

}
