package icburfq

// RfqRemoteServiceResult 
type RfqRemoteServiceResult struct {
    // 错误信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 返回结果信息
    Result   *RfqQuotationRemoteDTO `json:"result,omitempty" xml:"result,omitempty"`
    // 错误类型
    ErrType   string `json:"err_type,omitempty" xml:"err_type,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误类型
    ErrorType   string `json:"error_type,omitempty" xml:"error_type,omitempty"`
}
