package exchange

// RefundMessageResult 结构体
type RefundMessageResult struct {
	// 当前页面的留言条数
	PageResults int64 `json:"page_results,omitempty" xml:"page_results,omitempty"`
	// 留言记录
	Results []RefundMessage `json:"results,omitempty" xml:"results>refund_message,omitempty"`
	// 所有留言记录数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
	// 异常信息
	Exception string `json:"exception,omitempty" xml:"exception,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
