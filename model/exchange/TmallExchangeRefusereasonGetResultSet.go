package exchange

// TmallexchangerefusereasongetResultSet 结构体
type TmallexchangerefusereasongetResultSet struct {
	// 拒绝原因列表
	Results []Reason `json:"results,omitempty" xml:"results>reason,omitempty"`
	// 异常信息
	Exception string `json:"exception,omitempty" xml:"exception,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 拒绝原因总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
