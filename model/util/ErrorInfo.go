package util

// ErrorInfo 
type ErrorInfo struct {

    // errorCode
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // subErrorCode
    
    SubErrorCode   string `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty"`
    

    // errorMessage
    
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    

}
