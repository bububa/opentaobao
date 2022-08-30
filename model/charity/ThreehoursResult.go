package charity

// ThreehoursResult 结构体
type ThreehoursResult struct {
	// 结果码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否失败
	Fail string `json:"fail,omitempty" xml:"fail,omitempty"`
	// 信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 是否成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}
