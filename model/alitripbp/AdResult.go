package alitripbp

// AdResult 结构体
type AdResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
