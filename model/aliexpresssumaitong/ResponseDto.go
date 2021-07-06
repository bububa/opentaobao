package aliexpresssumaitong

// ResponseDto 结构体
type ResponseDto struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 数据
	Data *HjTaxCalculateResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 成功标记
	Succeeded bool `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
}
