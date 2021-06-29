package c2m

// TaobaoSebpCompanyGetorderinfoResultDo 
type TaobaoSebpCompanyGetorderinfoResultDo struct {

    // 调用是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 结果信息
    
    Module   *PageInfo `json:"module,omitempty" xml:"module,omitempty"`
    

    // 错误信息
    
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

}
