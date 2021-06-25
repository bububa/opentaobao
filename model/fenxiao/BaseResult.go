package fenxiao

// BaseResult 
type BaseResult struct {

    // errorCode
    ErrorCode   string `json:"error_code,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

    // data
    Data   string `json:"data,omitempty"`

}