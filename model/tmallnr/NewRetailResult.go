package tmallnr

// NewRetailResult 结构体
type NewRetailResult struct {
	// 请求结果
	ResultData *TagRespDto `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 成功或者失败
	SuccessFlag bool `json:"success_flag,omitempty" xml:"success_flag,omitempty"`
}
