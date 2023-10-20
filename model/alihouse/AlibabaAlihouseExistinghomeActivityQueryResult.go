package alihouse

// AlibabaalihouseexistinghomeactivityqueryResult 结构体
type AlibabaalihouseexistinghomeactivityqueryResult struct {
	// 结果信息
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 处理完成后的消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误/成功码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
