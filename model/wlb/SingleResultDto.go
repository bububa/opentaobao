package wlb

// SingleResultDto 
type SingleResultDto struct {

    // errorCode
    ErrorCode   string `json:"error_code,omitempty"`

    // errorMessage
    ErrorMessage   string `json:"error_message,omitempty"`

    // success
    Flag   bool `json:"flag,omitempty"`

}