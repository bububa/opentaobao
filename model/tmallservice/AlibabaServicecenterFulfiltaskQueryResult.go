package tmallservice

// AlibabaServicecenterFulfiltaskQueryResult 结构体
type AlibabaServicecenterFulfiltaskQueryResult struct {
	// 核销单
	Values []Value `json:"values,omitempty" xml:"values>value,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
