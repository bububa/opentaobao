package normalvisa

// TaobaoalitriptravelnormalvisastoreuserResultSet 结构体
type TaobaoalitriptravelnormalvisastoreuserResultSet struct {
	// 结果：数字数组
	Results []int64 `json:"results,omitempty" xml:"results>int64,omitempty"`
	// 异常
	Exception string `json:"exception,omitempty" xml:"exception,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 结果数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否包含下一个
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}
