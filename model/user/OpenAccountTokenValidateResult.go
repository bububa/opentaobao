package user

// OpenAccountTokenValidateResult 
type OpenAccountTokenValidateResult struct {

    // token中的数据
    
    Data   *TokenInfo `json:"data,omitempty" xml:"data,omitempty"`
    

    // 错误信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 是否成功
    
    Successful   bool `json:"successful,omitempty" xml:"successful,omitempty"`
    

    // 错误码
    
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    

}
