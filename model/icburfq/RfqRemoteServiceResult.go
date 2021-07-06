package icburfq

// RfqRemoteServiceResult 结构体
type RfqRemoteServiceResult struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误类型
	ErrType string `json:"err_type,omitempty" xml:"err_type,omitempty"`
	// 错误类型
	ErrorType string `json:"error_type,omitempty" xml:"error_type,omitempty"`
	// 返回结果信息
	Result *RfqQuotationRemoteDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
