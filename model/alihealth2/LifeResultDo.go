package alihealth2

// LifeResultDo 
type LifeResultDo struct {

    // RT
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

    // 描述信息
    
    Info   string `json:"info,omitempty" xml:"info,omitempty"`
    

    // 1代表成功！
    
    RetCode   string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
    

    // errorCode
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 是不是成功！
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
