package alitripbp

// AdResult 结构体
type AdResult struct {
	// 1
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 1
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 查询结果
	Model *ChannelExamineResultDto `json:"model,omitempty" xml:"model,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
