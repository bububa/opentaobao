package txcs

// AccessBaseResult 
type AccessBaseResult struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 错误信息
    
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    

    // 错误编码
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 具体内容
    
    Model   *WebPageData `json:"model,omitempty" xml:"model,omitempty"`
    

}
