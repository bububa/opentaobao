package tmallservice

// AlibabasscsupplyplatformserviceworkerregisterResult 结构体
type AlibabasscsupplyplatformserviceworkerregisterResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 结果model
	Value *WorkerCreateDto `json:"value,omitempty" xml:"value,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
