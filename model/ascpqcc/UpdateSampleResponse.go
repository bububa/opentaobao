package ascpqcc

// UpdateSampleResponse 
type UpdateSampleResponse struct {

    // 业务系统错误信息
    
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    

    // 业务系统错误编号
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 业务系统是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
