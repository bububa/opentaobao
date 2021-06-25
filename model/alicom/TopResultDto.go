package alicom

// TopResultDto 
type TopResultDto struct {

    // 错误信息
    ErrorMessage   string `json:"error_message,omitempty"`

    // data
    Data   *ProductActivityInfoResponseDto `json:"data,omitempty"`

    // errorCode
    ErrorCode   string `json:"error_code,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
