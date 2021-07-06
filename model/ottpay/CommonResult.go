package ottpay

// CommonResult 结构体
type CommonResult struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data *TvOrderResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 返回码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 返回结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
