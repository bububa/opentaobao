package tvupadmin

// TopResult 
type TopResult struct {
    // 实际内容
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
    // retCode
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
    // errorCode
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // object
    Object   bool `json:"object,omitempty" xml:"object,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}
