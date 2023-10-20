package alihealth2

// AlibabaalihealthtmscutconfirmResult 结构体
type AlibabaalihealthtmscutconfirmResult struct {
	// 结果
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
