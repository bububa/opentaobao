package scs

// ResultDto 结构体
type ResultDto struct {
	// 可用余额
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 结果信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结果code
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
