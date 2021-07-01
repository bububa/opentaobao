package logistic

// BaseResultDto 结构体
type BaseResultDto struct {
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回信息
	Module *ReachableServiceWaybillForTopResponseDto `json:"module,omitempty" xml:"module,omitempty"`
	// 请求错误信息
	OneErrorInfo *ErrorInfo `json:"one_error_info,omitempty" xml:"one_error_info,omitempty"`
}
