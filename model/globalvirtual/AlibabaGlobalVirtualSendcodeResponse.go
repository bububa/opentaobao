package globalvirtual

// AlibabaGlobalVirtualSendcodeResponse 
type AlibabaGlobalVirtualSendcodeResponse struct {
    // request result
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // send code result
    Module   bool `json:"module,omitempty" xml:"module,omitempty"`
    // error code
    ErrorCode   *ErrorCode `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // request is repeated
    Repeated   bool `json:"repeated,omitempty" xml:"repeated,omitempty"`
    // request need retry
    Retry   bool `json:"retry,omitempty" xml:"retry,omitempty"`
}
