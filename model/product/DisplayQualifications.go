package product

// DisplayQualifications 结构体
type DisplayQualifications struct {
	// 返回消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 填充列表
	Model string `json:"model,omitempty" xml:"model,omitempty"`
}
