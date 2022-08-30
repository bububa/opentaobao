package security

// DiAfiResult 结构体
type DiAfiResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 用户参数
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
}
