package lsttrade

// AlibabalsttradesellerorderdetailqueryResultDto 结构体
type AlibabalsttradesellerorderdetailqueryResultDto struct {
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回值
	Content *Content `json:"content,omitempty" xml:"content,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
