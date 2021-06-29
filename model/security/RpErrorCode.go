package security

// RpErrorCode 
type RpErrorCode struct {
    // errorCode
    ErrorCode   int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // errorName
    ErrorName   string `json:"error_name,omitempty" xml:"error_name,omitempty"`
}
