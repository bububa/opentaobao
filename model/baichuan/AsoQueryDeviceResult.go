package baichuan

// AsoQueryDeviceResult 
type AsoQueryDeviceResult struct {
    // result
    Results   []AsoDeviceCheckResult `json:"results,omitempty" xml:"results>aso_device_check_result,omitempty"`
    // errorDetail
    ErrorDetail   string `json:"error_detail,omitempty" xml:"error_detail,omitempty"`
    // errorCode
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
