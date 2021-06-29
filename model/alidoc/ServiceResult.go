package alidoc

// ServiceResult 
type ServiceResult struct {

    // success
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // errMessage
    
    ErrMessage   string `json:"err_message,omitempty" xml:"err_message,omitempty"`
    

    // errCode
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // 返回数据对象
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    

}
