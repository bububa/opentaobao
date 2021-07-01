package tmallgenie

// AiCloudResult 结构体
type AiCloudResult struct {
	// 语音地址
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 返回item列表
	Models []string `json:"models,omitempty" xml:"models>string,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
