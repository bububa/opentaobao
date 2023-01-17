package flight

// ResultDto 结构体
type ResultDto struct {
	// 错误code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 查询协同单详情返回内容
	Data *CaseResultDetailDto `json:"data,omitempty" xml:"data,omitempty"`
	// true
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
