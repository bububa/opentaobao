package aliexpress

// Response 
type Response struct {

    // 错误码
    
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    

    // 成功返回结果，json字符串
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    

    // 是否请求成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 错误信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

}
