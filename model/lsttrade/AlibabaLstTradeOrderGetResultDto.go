package lsttrade

// AlibabaLstTradeOrderGetResultDto 结构体
type AlibabaLstTradeOrderGetResultDto struct {
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 返回模型
	Content *LstTopOrderDto `json:"content,omitempty" xml:"content,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
