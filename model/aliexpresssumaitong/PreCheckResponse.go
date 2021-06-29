package aliexpresssumaitong

// PreCheckResponse 
type PreCheckResponse struct {

    // 错误码
    
    ErrorCode   *ErrorCode `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 错误信息
    
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
    

    // 请求traceId
    
    RequestTrace   string `json:"request_trace,omitempty" xml:"request_trace,omitempty"`
    

    // 请求是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 返回的token
    
    Token   string `json:"token,omitempty" xml:"token,omitempty"`
    

}
