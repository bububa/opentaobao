package shenjing

// PageResult 
type PageResult struct {

    // 返回内容(分页对象)
    
    Content   *Page `json:"content,omitempty" xml:"content,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

    // 成功标示
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
