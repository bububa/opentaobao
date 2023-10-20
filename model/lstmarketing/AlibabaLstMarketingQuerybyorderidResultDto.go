package lstmarketing

// AlibabalstmarketingquerybyorderidResultDto 结构体
type AlibabalstmarketingquerybyorderidResultDto struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 订单实体
	Content *LstTopOrderDto `json:"content,omitempty" xml:"content,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
