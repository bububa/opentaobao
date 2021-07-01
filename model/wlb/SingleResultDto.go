package wlb

// SingleResultDto 
type SingleResultDto struct {
    // errorCode
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // errorMessage
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // success
    Flag   bool `json:"flag,omitempty" xml:"flag,omitempty"`
}
