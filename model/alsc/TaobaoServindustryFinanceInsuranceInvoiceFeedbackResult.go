package alsc

// TaobaoServindustryFinanceInsuranceInvoiceFeedbackResult 结构体
type TaobaoServindustryFinanceInsuranceInvoiceFeedbackResult struct {
	// 鹰眼id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 开票结果
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
