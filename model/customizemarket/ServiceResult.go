package customizemarket

// ServiceResult 
type ServiceResult struct {

    // data
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    

    // errorCode
    
    ErrorCode   int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // errorMsg
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

    // suc
    
    Suc   bool `json:"suc,omitempty" xml:"suc,omitempty"`
    

}
