package baichuan

// AsoQueryDeviceResult 
/* model for simplify = false
type AsoQueryDeviceResult struct {

    // result
    
    Results  struct {
        AsoDeviceCheckResult  []AsoDeviceCheckResult `json:"aso_device_check_result,omitempty"`
    } `json:"results,omitempty"`
    

    // errorDetail
    
    ErrorDetail   string `json:"error_detail,omitempty"`
    

    // errorCode
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AsoQueryDeviceResult 
type AsoQueryDeviceResult struct {

    // result
    Results   []AsoDeviceCheckResult `json:"results,omitempty"`

    // errorDetail
    ErrorDetail   string `json:"error_detail,omitempty"`

    // errorCode
    ErrorCode   string `json:"error_code,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
