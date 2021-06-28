package flight

// RefundOrderFillConfirmFeeRs 
type RefundOrderFillConfirmFeeRs struct {

    // 错误编码
    
    ApiErrorCode   int64 `json:"api_error_code,omitempty" xml:"api_error_code,omitempty"`
    

    // 错误秒速
    
    ApiErrorMsg   string `json:"api_error_msg,omitempty" xml:"api_error_msg,omitempty"`
    

    // errTrace
    
    ErrTrace   string `json:"err_trace,omitempty" xml:"err_trace,omitempty"`
    

    // 是否失败
    
    Failure   bool `json:"failure,omitempty" xml:"failure,omitempty"`
    

    // 主机名称
    
    HostName   string `json:"host_name,omitempty" xml:"host_name,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
