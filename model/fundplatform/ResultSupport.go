package fundplatform

// ResultSupport 
type ResultSupport struct {
    // errorCode
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // errorMessage
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // 出参对象
    Module   *AccountChargeResponse `json:"module,omitempty" xml:"module,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
