package tmallcampus

// TResult 
type TResult struct {

    // 调用失败
    
    Fail   bool `json:"fail,omitempty" xml:"fail,omitempty"`
    

    // 数据
    
    Data   *CollegeAuditStatusResponse `json:"data,omitempty" xml:"data,omitempty"`
    

    // 调用是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 错误信息
    
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    

    // 错误代码
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

}
